drop table if exists public.author cascade;
drop table if exists public.book cascade;
drop table if exists public.authors_books cascade;

create table public.author
(
    id         uuid primary key      default gen_random_uuid(),
    name       varchar(100) not null,
    age        int,
    is_alive   bool,
    created_at timestamp    not null default (now() at time zone 'utc')
);

create index idx_author_created_at_pagination on public.author (created_at, id);
create index idx_author_age_pagination on public.author (age, id);

create table public.book
(
    id         UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    name       varchar(100) not null,
    age        int,
    created_at timestamp    not null default (now() at time zone 'utc')
);

create index idx_book_created_at_pagination on public.book (created_at, id);
create index idx_book_age_pagination on public.book (age, id);

create table public.authors_books
(
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id   uuid not null,
    author_id uuid not null,

    CONSTRAINT book_fk foreign key (book_id) references public.book (id),
    CONSTRAINT author_fk foreign key (author_id) references public.author (id),
    CONSTRAINT book_author_unique UNIQUE (book_id, author_id)
);

insert into public.book (id, name, age)
VALUES ('16a7fb15-6289-46c9-8738-863ea6292d6f', 'Fantastic balls', 100);

INSERT INTO public.author (name, age, is_alive)
values ( 'Danil', 1, true);
INSERT INTO public.author (name, age, is_alive)
values ( 'Barak', 54, true);
INSERT INTO public.author (name, age, is_alive)
values ( 'Bebra', 17, true);
INSERT INTO public.author (name, age, is_alive)
values ( 'Stone', 20, false);
INSERT INTO public.author (name, age, is_alive)
values ( 'Brabus', 24, false);
INSERT INTO public.author (name, age, is_alive)
values ( 'Cloze', 43, false);
INSERT INTO public.author (name, age, is_alive)
values ( 'Traire', 62, false);
INSERT INTO public.author (name, age, is_alive)
values ( 'Blaze', 51, true);
-- INSERT INTO public.author (id, name, age, is_alive)
-- values ('cb7f6f2b-8663-467f-9c80-d3ea364be7ef', 'Mraive', 83, true);
-- INSERT INTO public.author (id, name, age, is_alive)
-- values ('02149f25-3f4a-4e0e-994f-263f8ed2ab5c', 'Treaz', 77, true);

-- insert into public.authors_books(book_id, author_id)
-- values ('16a7fb15-6289-46c9-8738-863ea6292d6f', '02149f25-3f4a-4e0e-994f-263f8ed2ab5c');
-- insert into public.authors_books(book_id, author_id)
-- values ('16a7fb15-6289-46c9-8738-863ea6292d6f', 'cb7f6f2b-8663-467f-9c80-d3ea364be7ef');

insert into public.book (id, name)
VALUES ('e552d411-ea26-49c9-b208-418ca895fbd4', 'Crazy cuts');

INSERT INTO public.author (id, name, age, is_alive)
values ('82220fe0-71b7-4684-a6fd-b8a93e62df43', 'Vagran', 67, false);
INSERT INTO public.author (id, name, age, is_alive)
values ('7334cf1b-2b57-4592-9b65-f1a69b6acd54', 'Trump', 64, true);

insert into public.authors_books(book_id, author_id)
values ('e552d411-ea26-49c9-b208-418ca895fbd4', '82220fe0-71b7-4684-a6fd-b8a93e62df43');
insert into public.authors_books(book_id, author_id)
values ('e552d411-ea26-49c9-b208-418ca895fbd4', '7334cf1b-2b57-4592-9b65-f1a69b6acd54');


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

select *,
       (select count(*) from authors_books where authors_books.book_id = b.id) as ath_count
from public.book b;