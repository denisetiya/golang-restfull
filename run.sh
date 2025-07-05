#!/bin/bash

# Install dependencies
echo "Installing dependencies..."
go mod tidy

# Check if PostgreSQL is running
echo "Checking PostgreSQL connection..."
if ! command -v psql &> /dev/null; then
    echo "PostgreSQL is not installed. Please install PostgreSQL first."
    exit 1
fi

# Create database if it doesn't exist
echo "Creating database..."
createdb rest_api_db 2>/dev/null || echo "Database already exists or couldn't be created"

# Run the application
echo "Starting the application..."
go run main.go
