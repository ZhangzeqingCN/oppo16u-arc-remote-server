create table notifications
(
    id          integer
        primary key,
    content     text,
    receiver_id integer
        constraint fk_notifications_receiver
            references users
);