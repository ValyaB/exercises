# GRPC, JSON strem, API exercises

- Given a file with ports data (ports.json), write 2 services. 
- First service(ClientAPI) should parse the JSON file and have REST interface.
- The file is of unknown size, it can contain several millions of records.
The service has limited resources available (e.g. 200MB ram).
- While reading the file, it should call a second service(PortDomainService) that either creates a new record in a database, or updates the existing one.
- The end result should be a database containing the ports, representing the latest version found in the JSON. Database can be Map.
- First service(ClientAPI) should provide an endpoint to retrieve the data from the second service(PortDomainService)
- Each service should be built using Dockerfile
- Provide all tests that you think are needed for your assignment. This will allow the reviewer to evaluate your critical thinking as well as your knowledge about testing.
- Please use gRPC as a transport between services.
- Readme should explain how to run code and tests

Choose the approach that you think is best (ie most flexible).

## Bonus points

- Database in docker container
- Domain Driven Design
- Docker-compose file

## Note

As mentioned earlier, the services have limited resources available, and the JSON file can be several hundred megabytes (if not gigabytes) in size.
This means that you will not able to read the entire file at once.

We are looking for the ClientAPI (the service reading the JSON) to be written in a way that is easy to reuse, give or take a few customisations.
The services themselves should handle certain signals correctly (e.g. a TERM or KILL signal should result in a graceful shutdown).

~~~~
    cd JSON_GRPC_microservice
    go run client/main.go
    go run server/main.go

http://127.0.0.1:8000/ports/upload

~~~~