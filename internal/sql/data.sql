create table public.author
(
    id   uuid primary key default gen_random_uuid(),
    name varchar(100) not null
);

create table public.book
(
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name      varchar(100) not null,
    author_id uuid         not null,
    constraint author_fx foreign key (author_id) references public.author
);

INSERT INTO public.author (name)
values ('Danil');

INSERT INTO public.author (name)
values ('Anytolies');

insert into public.book (name, author_id)
VALUES ('Fantastic balls', '02149f25-3f4a-4e0e-994f-263f8ed2ab5c');