CREATE TABLE public.user(
                            id int primary key generated always as identity unique,
                            username varchar(100) not null unique ,
                            password varchar(100) not null ,
                            email varchar(100) not null unique
);


create table public.category(
                                id int primary key generated always as identity,
                                name varchar(100) unique
);

CREATE TABLE public.item(
                            id int primary key generated always as identity,
                            name varchar(100) not null unique ,
                            price int not null
);

CREATE TABLE public.item_category(
                            item_id int not null,
                            category_id int not null,
                            FOREIGN KEY (item_id) references public.item(id),
                            FOREIGN KEY (category_id) references public.category(id)
);