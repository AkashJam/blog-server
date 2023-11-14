# server

Blog server used to create api's from access to persistent data

## Setup
```go mod init [module_name]``` to create new modules
```go mod tidy``` to add required modules

## Launch
```go run .``` in server directory to run app.go

## Deploy
```docker compose build``` in BEdraft directory to build docker image
```docker compose up``` to create a container of the image