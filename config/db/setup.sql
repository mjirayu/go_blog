drop table posts;

create table posts (
  id         serial primary key,
  title      varchar(40),
  body       text,
  created_at timestamp not null
);
