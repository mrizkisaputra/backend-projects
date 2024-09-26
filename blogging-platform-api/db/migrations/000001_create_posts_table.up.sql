-- migrate -path=db/migrations -database="mysql://root:root@tcp(localhost:3306)/blog" up

CREATE TABLE posts
(
    id         bigint UNSIGNED not null,
    title      varchar(150)    not null,
    content    text            not null,
    category   varchar(100)    not null,
    tags       text,
    created_at datetime default null,
    updated_at datetime default null
);

ALTER TABLE posts
    ADD CONSTRAINT posts_pkey_id PRIMARY KEY (id)