# E Learning API
An API Service E Learning. Build with Go, Postgres, and Docker. This project using Repository pattern as its persistance data layer.

## Getting Started
First, you need to clone this repository. After that, you run below to install necessary dependencies:

```sh
$ go mod tidy
$ go install  
```

After that, run Postgres Container using below command:

```sh
docker compose up -d
```

And then lastly, you can run Gateway Server:

```sh 
go run main.go http-gw-srv --port 8080
```
