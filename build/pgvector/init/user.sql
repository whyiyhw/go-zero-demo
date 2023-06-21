create table public."user"
(
    id         bigserial    not null primary key,
    name       varchar(50)  not null    default '':: character varying,
    email      varchar(191) not null    default '':: character varying,
    password   varchar(255) not null    default '':: character varying,
    created_at timestamp with time zone,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone
);

comment
on column "user".id is '主键id';

comment
on column "user".name is '用户名称';

comment
on column "user".email is '邮箱';

comment
on column "user".password is '密码';

comment
on column "user".created_at is '创建时间';

comment
on column "user".updated_at is '更新时间';

alter table public."user"
    owner to root;

create unique index user_email_idx_unique on public."user" (email);

comment
on index user_email_idx_unique is '用户邮箱唯一索引';