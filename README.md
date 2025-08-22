# Schrödinger's Database

A key-value store built in Go that **randomly breaks** — sometimes returning the wrong data or failing mysteriously. 

## Features
- Store, retrieve, and delete key-value pairs.
- List all stored key-value pairs.
- Randomized behavior to mimic "unstable" database responses.
- Simple CLI powered by [Cobra](https://github.com/spf13/cobra).
- PostgreSQL backend connection with `.env` configuration.

## Installation
```bash
git clone https://github.com/rodopiip/schrodinger-db.git

# go to repo folder for initial docker image build
cd <repo-folder>

# build app container
docker build -t schrodinger-db .

# run postgres container
docker run --rm --name schrodinger-postgres -e POSTGRES_USER=maria -e POSTGRES_PASSWORD=5432 -e POSTGRES_DB=schrodingerdatabase -p 5432:5432 postgres:16

# run app container that displays help panel
docker run --rm --name schrodinger-app -e HOST=host.docker.internal -e PORT=5432 -e USER=maria -e PASSWORD=5432 -e DB_NAME=schrodingerdatabase schrodinger-db --help
```

## Shcrodinger CLI Commands 
#### OUTSIDE docker container 
```bash
# store a value
go run . put mykey myvalue

# retrieve a value
go run . get mykey

# delete a value
go run . del mykey

# dump all values
go run . dump

# show help
go run . --help

```
### CLI Showcase
![img.png](img.png)

## Unit Tests
```bash
# Run unit tests

go test .
```