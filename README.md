# Go Todo REST API Example
A dockerised version of a simple RESTful API example for a todo application written in Go. The purpose of this project is to server as a sample deployable app to demo deployments to container orchestration systems such as ECS or Kubernetes. It has a database dependency, so will be more interesting to play with than a simple `hello-world`.

Based on [mingrammer/go-todo-rest-api-example](https://github.com/mingrammer/go-todo-rest-api-example).

The main changes:

* swapped database engine for Postgres
* dockerised all the things
* added retries so that `docker-compose up` works out of the box.

## Run locally
```bash
# Download this project
git clone endofcake/go-todo-rest-api-example
```

Run `docker-compose`:
```bash
docker-compose up -d
```
This will create a Postgres database and run the schema migration, and then start the web server on port 3000.

Verify that everything is working:

```bash
# no projects yet
curl http://localhost:3000/projects

# add a sample project
curl -d '{"title":"sample project"}' -H "Content-Type: application/json" -X POST http://localhost:3000/projects

# check that it's saved
curl http://localhost:3000/projects
```

Check the database:
```bash
export PGPASSWORD=qwerty

# connect to the db
psql -h localhost -d todoapp -U docker

# poke around in psql
\d+ # and other psql commands
```

## Build
```bash
# Build and Run
cd go-todo-rest-api-example
go build
./go-todo-rest-api-example

# API Endpoint : http://127.0.0.1:3000
```

## Deploy
Deploy as a standard container app using your favourite deployment method.

The app will require certain parameters to be set as environment variables in order to connect to the database (needs to be provisioned separately):
```
PGHOST
PGDATABASE
PGUSER
PGPASSWORD
PGSSLMODE
```

`PGPASSWORD` is a secret and should be treated accordingly. To simplify secret management in AWS, the Docker image includes the [pstore](https://github.com/glassechidna/pstore) utility, so you can keep the password in AWS Parameter Store and decrypt it at runtime by setting `PSTORE_PGPASSWORD` environment variable to point to the secret name. This will require appropriate IAM permissions.

## Structure
```
├── app
│   ├── app.go
│   ├── handler          // Our API core handlers
│   │   ├── common.go    // Common response functions
│   │   ├── projects.go  // APIs for Project model
│   │   └── tasks.go     // APIs for Task model
│   └── model
│       └── model.go     // Models for our application
├── config
│   └── config.go        // Configuration
└── main.go
```

## API

#### /projects
* `GET` : Get all projects
* `POST` : Create a new project

#### /projects/:title
* `GET` : Get a project
* `PUT` : Update a project
* `DELETE` : Delete a project

#### /projects/:title/archive
* `PUT` : Archive a project
* `DELETE` : Restore a project

#### /projects/:title/tasks
* `GET` : Get all tasks of a project
* `POST` : Create a new task in a project

#### /projects/:title/tasks/:id
* `GET` : Get a task of a project
* `PUT` : Update a task of a project
* `DELETE` : Delete a task of a project

#### /projects/:title/tasks/:id/complete
* `PUT` : Complete a task of a project
* `DELETE` : Undo a task of a project
