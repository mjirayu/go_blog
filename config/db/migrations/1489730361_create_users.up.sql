create table users (
  id          serial primary key,
  name        varchar(40),
  email       varchar(50),
  password    varchar(100),
  created_at  timestamp not null
);
