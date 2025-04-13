CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    email varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE books
(
    id serial not null unique,
    name varchar(255) not null,
    face varchar(255) not null unique,
    last_page int not null UNIQUE,
    path VARCHAR(255) not null unique
);

CREATE TABLE users_books
(
    id serial not null unique,
    from_id int not null UNIQUE,
    to_id int not null UNIQUE 
);

