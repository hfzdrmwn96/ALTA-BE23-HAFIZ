CREATE TABLE users (
  id int primary key, 
  name VARCHAR(100) ,
  email VARCHAR(100),
  password VARCHAR(255),
  phone VARCHAR(15),
  birth_date date
);

CREATE TABLE admins (
  id int primary key, 
  name VARCHAR(100) ,
  email VARCHAR(100),
  password VARCHAR(255),
  phone VARCHAR(15)
);

CREATE TABLE book_request (
  id int primary key, 
  user_id int references users(id) ,
  approved_admin_id int references admins(id),
  title varchar(100),
  author varchar(100),
  publisher varchar(100),
  publish_year int,
  status_request varchar(100)
);

CREATE TABLE genres (
  id int primary key,
  genre_name varchar(20)
);

CREATE TABLE books (
  id int primary key,
  genre_id int references genres(id),
  title varchar(100),
  author varchar(100),
  publisher varchar(100),
  publish_year int
);

CREATE TABLE loans (
  id int primary key,
  user_id int references users(id),
  book_id int references books(id),
  deadline_date date,
  date_loan date,
  date_return date,
  status_loan varchar(50)
);

alter table users add column address varchar(255);

insert into admins (id, name, email, password, phone)
values
(1,'Hafiz', 'hafizdarmawan1996@gmail.com','12345678','082178561963' ),
(2,'Budi', 'budimanrusdi1996@gmail.com','11223344','082178561964'),
(3,'Rana', 'ranaandelap@gmail.com','45645645','082178561965');

insert into genres (id, genre_name)
values
(1,'Romance'),
(2,'Horror');

insert into users (id, name, email, password, address, phone, birth_date)
values
(1, 'Yusman', 'yusmantedi@gmail.com', '787878788', 'Jalan soekarno-hatta no 2 Jakarta', '08127845613544', '1996-04-22' );

insert into books (id, genre_id , title , author, publisher , publish_year)
values
(1, 1 , 'Outlander' , 'Diana Gabaldon', 'Delacorte Books' , 1991);