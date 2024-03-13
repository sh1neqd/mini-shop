CREATE TABLE public.user(
                            id int primary key generated always as identity,
                            username varchar(100) not null ,
                            password varchar(100) not null ,
                            email varchar(100) not null
);


create table public.category(
                                id int primary key generated always as identity,
                                name varchar(100)
);

CREATE TABLE public.item(
                            id int primary key generated always as identity,
                            name varchar(100) not null,
                            category_id int REFERENCES category(id)
);