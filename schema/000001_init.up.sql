CREATE TABLE users
(
    id serial primary key,
    name varchar(30),
    surname varchar(30),
    patronymic varchar(30),
    age int,
    country varchar(30),
    gender varchar(30)
);