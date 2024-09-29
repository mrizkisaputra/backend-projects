-- migrate -path=db/migrations -database="mysql://todosapi:4PG5tmTC9CMdWqmMIXZtARNctiBBQr@tcp(localhost:3306)/todosapi" up

CREATE TABLE users
(
    id varchar(100) not null,
    name varchar(100) not null,
    email varchar(100) not null,
    password varchar(100) not null,
    token varchar(100) default null,
    created_at bigint default 0
);

ALTER TABLE users
ADD CONSTRAINT users_pk_id PRIMARY KEY (id)
