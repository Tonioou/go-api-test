CREATE USER todo_app;
CREATE DATABASE todo_list;

GRANT ALL PRIVILEGES ON DATABASE todo_list TO todo_app;

ALTER USER todo_app WITH PASSWORD 'todo_app';
