#!/bin/bash

# Build the application
echo "Building the application..."
go build -o bin/rest-api main.go

# Run the application
echo "Starting the REST API server..."
./bin/rest-api
