CREATE USER todo_app;
CREATE DATABASE todo_list;
GRANT ALL PRIVILEGES ON DATABASE todo_list TO todo_app;

CREATE TABLE Todo(
    id uuid not null primary key,
    name text not null,
    created_at timestamp not null,
    finished_at timestamp,
    finished bool,
    active bool not null
);

CREATE TABLE TodoItem(
    id uuid not null primary key,
    todo_id uuid not null,
    description text not null,
    created_at timestamp not null,
    active bool not null,
    FOREIGN KEY(todo_id)
    REFERENCES Todo(id)
);


