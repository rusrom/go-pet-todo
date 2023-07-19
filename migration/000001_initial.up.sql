create table users
(
    id       serial       not null unique,
    name     varchar(255) not null,
    username varchar(255) not null unique,
    password varchar(255) not null
);

create table lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255)
);

create table users_lists
(
    id      serial not null unique,
    user_id int    not null,
    list_id int    not null,
    constraint fk_user foreign key (user_id) references users (id) on delete cascade,
    constraint fk_list foreign key (list_id) references lists (id) on delete cascade

);

create table items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false

);

create table lists_items
(
    id      serial not null unique,
    item_id int    not null,
    list_id int    not null,
    constraint fk_item foreign key (item_id) references items (id) on delete cascade,
    constraint fk_list foreign key (list_id) references lists (id) on delete cascade
);
