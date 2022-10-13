create table users
(
    id          int auto_increment
        primary key,
    name        varchar(255)                       null,
    telegram_id int                                null,
    first_name  varchar(255)                       null,
    last_name   varchar(255)                       null,
    chat_id     int                                null,
    created_at  datetime default CURRENT_TIMESTAMP null,
    updated_at  datetime                           null,
    deleted_at  datetime                           null
);

create table adverts
(
    id         int auto_increment
        primary key,
    user_id    int                                not null,
    car_mark   varchar(255)                       not null,
    car_model  varchar(255)                       null,
    price      float    default 0                 null,
    year       int                                null,
    created_at datetime default CURRENT_TIMESTAMP null,
    update_at  datetime                           null,
    deleted_at datetime                           null,
    constraint adverts_users_id_fk
        foreign key (user_id) references users (id)
            on update cascade on delete cascade
);
