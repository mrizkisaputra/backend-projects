CREATE TABLE tasks
(
    id varchar(100) not null,
    id_user varchar(100) not null,
    title varchar(100) not null,
    description varchar(200) not null,
    status char(12) default "in-progress",
    created_at bigint default 0,
    updated_at bigint default 0
);

ALTER TABLE tasks
ADD CONSTRAINT tasks_pk_id PRIMARY KEY (id);

ALTER TABLE tasks
ADD CONSTRAINT tasks_fk_id_user FOREIGN KEY (id_user) REFERENCES users(id);