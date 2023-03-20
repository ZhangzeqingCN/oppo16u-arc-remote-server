# Entity

## users
```sqlite
create table users
(
    id       integer
        primary key
        constraint fk_notifications_receiver
            references main.notifications,
    name     text,
    password text,
    email    text
        unique,
    token    text
);
```

## tasks
```sqlite
create table main.tasks
(
    id             integer
        primary key,
    name           text,
    user_id        integer -- 任务的所有者
        constraint fk_tasks_user
            references main.users
            on update cascade on delete set null,
    pre_task_of_id integer -- 作为谁的前置任务
        constraint fk_tasks_pre_tasks
            references main.tasks,
    sub_task_of_id integer -- 作为谁的子任务
        constraint fk_tasks_sub_tasks
            references main.tasks
);
```

## devices
```sqlite
create table devices
(
    id   integer
        primary key,
    name text
);
```

## device_task_relations
```sqlite
create table device_task_relations
(
    id        integer
        primary key,
    device_id integer
        constraint fk_device_task_relations_device
            references main.devices,
    task_id   integer
        constraint fk_device_task_relations_task
            references main.tasks
);
```

## user_relations
```sqlite 
create table user_relations
(
    id       integer
        primary key,
    user1_id integer
        constraint fk_user_relations_user1
            references main.users,
    user2_id integer
        constraint fk_user_relations_user2
            references main.users,
    tag      text
);
```
