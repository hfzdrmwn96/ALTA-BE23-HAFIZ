drop table users;
drop table admins;
drop table book_request;
drop table genres;
drop table books;
drop table loans;

CREATE TABLE users (
  id serial primary key not null, 
  name VARCHAR(100) ,
  email VARCHAR(100),
  password VARCHAR(255),
  phone VARCHAR(15),
  birth_date date
);

drop tab

CREATE TABLE admins (
  id SERIAL primary key not null, 
  name VARCHAR(100) ,
  email VARCHAR(100),
  password VARCHAR(255),
  phone VARCHAR(15)
);

CREATE TABLE book_request (
  id serial primary key not null, 
  user_id int references users(id) ,
  approved_admin_id int references admins(id),
  title varchar(100),
  author varchar(100),
  publisher varchar(100),
  publish_year int,
  status_request varchar(100)
);

CREATE TABLE genres (
  id serial primary key not null,
  genre_name varchar(20)
);

CREATE TABLE books (
  id serial primary key not null,
  genre_id int references genres(id),
  title varchar(100),
  author varchar(100),
  publisher varchar(100),
  publish_year int
);

CREATE TABLE loans (
  id serial primary key not null,
  user_id int references users(id),
  book_id int references books(id),
  deadline_date date,
  date_loan date,
  date_return date,
  status_loan varchar(50)
);

alter table users add column address varchar(255);

insert into admins (name, email, password, phone)
values
('Hafiz', 'hafizdarmawan1996@gmail.com','12345678','082178561963'),
('Budi', 'budimanrusdi1996@gmail.com','11223344','082178561964'),
('Rana', 'ranaandelap@gmail.com','45645645','082178561965');

insert into genres (genre_name)
values
('Romance'),
('Comedy'),
('Action'),
('Thriller'),
('Horror');

insert into users (name, email, password, address, phone, birth_date)
values
('Yusman', 'yusmantedi@gmail.com', '787878788', 'Jalan soekarno-hatta no 2 Jakarta', '082178561997', '1996-04-22' ),
('Rere', 'rere@gmail.com', '12345678', 'Jalan Ir Soekamto no 3 Jakarta', '082178561996', '1995-03-12' ),
('Alif', 'alif@gmail.com', '12345678', 'Jalan Ir Soekamto no 4 Jakarta', '082178564796', '1995-07-4' ),
('Indra', 'indra@gmail.com', '12345678', 'Jalan DI Panjaitan no 19 Jakarta', '081257561996', '1995-01-27' ),
('Hendra', 'hendra@gmail.com', '12345678', 'Jalan Sukamare no 69 Jakarta', '085778561996', '2001-11-06' );

insert into books (genre_id , title , author, publisher , publish_year)
values
(1,'Outlander' , 'Diana Gabaldon', 'Delacorte Books' , 1991),
(1,'Twilight, Book 1 : Twilight' , 'Stephenie Meyer', 'Little, Brown and Company' , 2005),
(1,'Fifty Shades of Grey' , 'E.L. James', 'Vintage' , 2011),
(1,'Easy' , 'Tammara Webber', 'Penguin Berkley' , 2012),
(1,'A Walk to Remember' , 'Nicholas Sparks', 'Grand Central Publishing' , 1991),
(2,'Bossypants' , 'Tina Fey', 'Reagan Arthur Books' , 2011),
(2,'Yes Please' , 'Amy Poehler', 'Dey St.' , 2014),
(2,'Mort' , 'Terry Pratchett', 'Harpertorch' , 1987),
(2,'The Princess Bride' , 'William Goldman', 'Ballantine Books' , 1973),
(2,'Equal Rites' , 'Terry Pratchett', 'Harper Perennial' , 1987),
(3,'Jurassic Park' , 'Michael Crichton', 'Ballantine Books' , 1990),
(3,'The Hobbit' , 'J.R.R. Tolkien', 'Houghton Mifflin' , 1937),
(3,'Divergent' , 'Veronica Roth', 'Katherine Tegen Books' , 2011),
(3,'A Game of Thrones' , 'George R.R. Martin', 'Bantam' , 1996),
(3,'Patriot Games' , 'Tom Clancy', 'Berkley' , 1987),
(4,'Birnam Wood' , 'Eleanor Catton', 'McClelland & Stewart' , 2023),
(4,'The Arrangement' , 'Mel Taylor', 'Severn River Publishing' , 2024),
(4,'Lights, Camera, Bones' , 'Carolyn Haines', 'Minotaur Books' , 2024),
(4,'The Return of Ellie Black' , 'Emiko Jean', 'Simon & Schuster' , 2024),
(4,'Bad Men' , 'Julie Mae Cohen', 'Harry N. Abrams' , 2024),
(5,'Never Lie' , 'Freida McFadden', 'Hollywood Upstairs Press' , 2022),
(5,'One by One' , 'Freida McFadden', 'Hollywood Upstairs Press' , 2020),
(5,'Two Twisted Crowns' , 'Rachel Gillig', 'Orbit' , 2023),
(5,'Ask for Andrea' , 'Noelle W. Ihli', 'Dynamite Books' , 2022),
(5,'The Locked Door' , 'Freida McFadden', 'Hollywood Upstairs Press' , 2021);

select * from books b where title = 'Divergent';
select * from books b where genre_id = 4;

insert into loans (user_id , book_id , deadline_date , date_loan , status_loan)
values
('1' , '1', '2024-06-14' ,'2024-06-07', 'dipinjam'),
('1' , '2', '2024-06-9' ,'2024-06-1', 'dipinjam'),
('1' , '8', '2024-06-11' ,'2024-06-5', 'dipinjam'),
('2' , '3', '2024-06-8' ,'2024-06-2', 'dipinjam'),
('3' , '1', '2024-06-14' ,'2024-06-7', 'dipinjam');

update loans
set 
date_return = '2024-06-07',
status_loan = 'dikembalikan'
where loans.user_id = 1 and book_id = 1;

select * 
from users u 
left join loans l on user_id = u.id;

select * 
from loans l 
join books b on b.id = l.book_id 
where l.status_loan ='dikembalikan';


