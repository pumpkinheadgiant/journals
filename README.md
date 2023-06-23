
## Installing
Having pulled to source down to a machine with a Golang runtime, type `go mod tidy` at the root to install dependencies. 
## Building
Execute the `build.sh` script (currently set to build for Mac M1 processor, but can be switched to linux by changing the uncommenting/commenting the appropriate lines in `build.sh`)
## Running
Type `./journals` to start the server, which will be running on port `8080`.
## Interacting
There is a POSTMAN collection available for importing in the `postman` directory.

### GET /journals
#### localhost:8080/journals
This endpoint will return a list of all existing journals.

### POST /journal
#### localhost:8080/journal
This endpoint allows for the creation of a new journal
