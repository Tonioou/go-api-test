\connect todo_list

CREATE TABLE todo(
    id uuid not null primary key,
    name text not null,
    created_at timestamp not null,
    finished_at timestamp,
    finished bool,
    active bool not null
);

CREATE TABLE todo_item(
    id uuid not null primary key,
    todo_id uuid not null,
    description text not null,
    created_at timestamp not null,
    active bool not null,
    FOREIGN KEY(todo_id)
    REFERENCES todo(id)
);


