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
    CONSTRAINT author_fk foreign key (author_id) references public.author (id)
);

INSERT INTO public.author (id, name)
values ('cb7f6f2b-8663-467f-9c80-d3ea364be7ef', 'Danil');

INSERT INTO public.author (id, name)
values ('02149f25-3f4a-4e0e-994f-263f8ed2ab5c', 'Barak');

insert into public.book (name) VALUES ('Fantastic balls');
