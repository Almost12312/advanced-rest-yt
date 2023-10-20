create table public.author
(
    id   uuid primary key default gen_random_uuid(),
    name varchar(100) not null
);

create table public.book
(
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(100) not null
);

create table public.authors_books
(
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id   uuid not null,
    author_id uuid not null,

    CONSTRAINT book_fk foreign key (book_id) references public.book (id),
    CONSTRAINT author_fk foreign key (author_id) references public.author (id),
    CONSTRAINT book_author_unique UNIQUE (book_id, author_id)
);

INSERT INTO public.author (id, name)
values ('cb7f6f2b-8663-467f-9c80-d3ea364be7ef', 'Danil');

INSERT INTO public.author (id, name)
values ('02149f25-3f4a-4e0e-994f-263f8ed2ab5c', 'Barak');

insert into public.book (id, name)
VALUES ('16a7fb15-6289-46c9-8738-863ea6292d6f', 'Fantastic balls');

insert into public.authors_books(book_id, author_id)
values ('16a7fb15-6289-46c9-8738-863ea6292d6f', '02149f25-3f4a-4e0e-994f-263f8ed2ab5c');
insert into public.authors_books(book_id, author_id)
values ('16a7fb15-6289-46c9-8738-863ea6292d6f', 'cb7f6f2b-8663-467f-9c80-d3ea364be7ef');

select b.id,
       b.name,
       array(select (ba.author_id) from public.authors_books ba where ba.book_id = b.id) as authors
from public.book b;

select a.name,
       author_id,
       book_id
from public.authors_books
         JOIN public.author a on a.id = authors_books.author_id
where book_id = '16a7fb15-6289-46c9-8738-863ea6292d6f';

select
    *,
    (select count(*) from authors_books where authors_books.book_id = b.id) as ath_count
from public.book b;