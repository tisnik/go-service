create table users (
    ID            integer primary key asc,
    name          text not null,
    surname       text not null
);

insert into users (id, name, surname) values (0, 'Linus', 'Torvalds');
insert into users (id, name, surname) values (1, 'Rob', 'Pike');
