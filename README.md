# Golang Web Server
This repository uses Go as a web server for backend APIs.

## Running the Go Server
To run the Go server, use the following command:

```cmd: go run filename.go```

### Initializing a Go Module
To initialize a Go module, use the following command:

```cmd: go mod init package_name```

### Managing Dependencies
To manage and install dependencies, use the following command:

```cmd: go mod tidy```

This command retrieves the dependencies of the module and installs any required packages specified in your import statements, such as "fmt".

### For live reload when you make some change on local host install fresh 

```go install github.com/gravityblast/fresh@latest```

```cmd: fresh```