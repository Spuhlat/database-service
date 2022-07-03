# database-service
Service to interface with all database management systems (DBMS)

## Prerequisites
- Go version 1.18
- MongoDB Community Version 5.0

## Getting Started
MongoDB server must be up and running for service to correctly operate.

To start service, from the root of the project, run:

```go run .```

to start the server.

Note that when developing and changes are made, the server will need to be restarted for changes to take effect.

## Environment Variables
`MONGO_HOST`: Hostname/IP address of mongodb server. Defaulted to `localhost`.

`MONGO_PORT`: Port of mongodb server. Defaulted to `27017`.