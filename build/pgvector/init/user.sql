create table "user"
(
    id         bigserial
        primary key,
    name       varchar(50)              default ''::character varying not null,
    email      varchar(191)             default ''::character varying not null,
    password   varchar(255)             default ''::character varying not null,
    created_at timestamp with time zone,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone
);

comment on column "user".id is '主键id';

comment on column "user".name is '用户名称';

comment on column "user".email is '邮箱';

comment on column "user".password is '密码';

comment on column "user".created_at is '创建时间';

comment on column "user".updated_at is '更新时间';

comment on column "user".deleted_at is '软删除字段';

alter table "user"
    owner to root;
