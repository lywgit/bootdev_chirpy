# Learn HTTP Servers in Go

This repository contains the code developed as I worked through the "[Learn HTTP Servers in Go](https://www.boot.dev/courses/learn-http-servers-golang)" course on [Boot.dev](https://www.boot.dev). The lessons guide you through the building of simple a HTTP web server *Chirpy* from scratch.

# Usage

## Prerequisites:

1. Go
2. A postgres server (Or get access to one)
3. A Database migration tool `goose` (Or alternitively create the tables and columns directly with SQL)
4. A HTTP client to send request to the server. Such as Curl or Postman

You can refer to the course content for step-by-step guides (materials are freely viewable).

## Start the chirp API server 

1. Clone the repo 
2. Prepare your own `.env` file (see env.example)
3. `cd` into the project root.
4. Run database migration: 
    1. go to the schema dir: `cd sql/schema` 
    2. migration up: `goose postgres postgres://lywang:@localhost:5432/chirpy up` (your connection string will be different)
5. Start API server 
    1. `cd` into the project root.
    2. Run `go run .`
    3. Server will be available at `http://localhost:8080` and APIs at `http://localhost:8080/api/`

## APIs

(TODO)

# My Takeways

Some of the lessons are pretty challenging for me as many concepts are new.
Key concepts learned:

- Routing and handler
- Http requests and information (header and json body)
- Http respond status code
- Authentication process: password, JWT and refresh tokens
- UUID in database
- goose: a database migration tool
- sqlc: a tool that converts SQL queries into Go functions


# Some Qeustions

- No Goroutine is explicitly used. Perhaps this is automatically handled under http.Server?
