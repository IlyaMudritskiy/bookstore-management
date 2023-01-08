# Bookstore API

## Description

CRUD API using Golang and MySQL

Before running the project:
- `pkg/settings`
  - Change name from `settings_example.go` to `settings.go`
  - Uncomment code
  - Replace credentials

- `docker-compose.yml`
  - Replace credentials according to `settings.go`

## Run

- `docker-compose up --build`
- `go build cmd/main/main.go`
- `./main`