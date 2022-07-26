create table if not exists todo (
  id int not null primary key auto_increment,
  title varchar(100) not null ,
  author varchar(30) not null ,
  created_at datetime
);