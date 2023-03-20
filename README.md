# R2 BACKEND APPLICATION

## Installation

Install the dependencies and start the server.

Run in docker:
```sh
(inside project directory)
docker build . -t r2-back 
docker run --name r2-back -p 8080:8080 -d r2-back    
```
Run locally
```sh
go mod download
go run cmd/main.go
```

## App Structure_
```sh
 - pkg
      /app
         fibonacci.service.go
      /config
      /domain
             fibonacci.go
             error.go
      /http
          /handlers
          /server
      /logger 
```
The main idea here is to create a reusable package. Following the principle "Domain Driven development" I define, all objects and the set of interactions we can perform in the objects inside the domain layer. This will not have any dependency. I use interfaces in order to separate definition from the implementation. This help us to do testing.

The, we have subpackages group by dependency. For example app subpackage is for business logic. HTTP is for all related to http functionality such as handlers, routes, etc. I also take into account, creating a wrapper when interacting with dependencies in order to decouple code.

 
