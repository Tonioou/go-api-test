# TODO List

That's a todo crud, where i'm exercising some things and trying others
like using open telemetry for tracing, writing my one openapi3 file, and
exercising golang architecture.

## How to run

You must have docker and docker compose installed on your machine

To start the database and application run, if running in a linux with make installed run:

`make compose-up` and to shut down server `make compose-down`

If you don't have make on your system then run from the root folder:

`docker compose up -f build/docker-compose.yaml` to start the database and application

and `docker compose down -f build/docker-compose.yaml` to stop.