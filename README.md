# Benefit for launching the application.

## 1. At the root of the project, create and ```.env``` file and place th following variables there:
```
POSTGRES_PASSWORD=DEFAULT_VALUE # string
POSTGRES_USER=DEFAULT_VALUE # string
POSTGRES_DB=DEFAULT_VALUE # string
POSTGRES_PORT=DEFAULT_VALUE # number
POSTGRES_HOST=DEFAULT_VALUE # string
POSTGRES_SSL_MODE=DEFAULT_VALUE # bool

SERVER_PORT=DEFAULT_VALUE # number
```

## 2. You must have Docker & Docker-compose on your local machine.
To start the database in a container separate from the application, run the `make up_db` command.

### All `make` functions are described in the `Makefile`.

### The project structure was generated with [create-project-struct](https://github.com/blackmarllbor0/create-project-struct).