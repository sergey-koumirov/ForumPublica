create table fp_users(
  id int not null auto_increment primary key,
  role varchar(2) not null default "U"
);

create table esi_characters(
  id int not null primary key,
  name varchar(255) not null
);
