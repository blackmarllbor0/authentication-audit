# Benefit for launching the application.

## 1. In the config package, create a ```config.yaml``` file and put the following content in it:
```
app:
  server:
    port: int
  db:
    dsn: string
```

## 2. To run the database in a docker container, you need to create an .env file with current data:
```
POSTGRES_USER=string
POSTGRES_PASSWORD=string
POSTGRES_DB=string
POSTGRES_PORT=int
```

## 3. You must have Docker & Docker-compose on your local machine.
To start the database in a container separate from the application, run the `make up_db` command.

### All `make` functions are described in the `Makefile`.

### The project structure was generated with [create-project-struct](https://github.com/blackmarllbor0/create-project-struct).