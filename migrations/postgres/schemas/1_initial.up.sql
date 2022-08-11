create table users (
    id bigserial primary key,
    username varchar(255) unique not null,
    password varchar(255) not null
);

create table lists (
    id bigserial primary key,
    user_id bigint not null references users(id) on delete cascade,
    title varchar(255) not null,
    description varchar(255) not null
);

create table notes (
   id bigserial primary key,
   list_id bigint not null references lists(id) on delete cascade,
   title varchar(255) not null,
   content varchar(8000) not null
);