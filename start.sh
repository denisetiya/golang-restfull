#!/bin/bash

# Script untuk menjalankan aplikasi REST API dengan PostgreSQL

set -e

echo "🚀 Starting REST API Application Setup"

# Function to check if MongoDB is running
check_mongodb() {
    if nc -z localhost 27017 >/dev/null 2>&1; then
        echo "✅ MongoDB is running"
        return 0
    else
        echo "❌ MongoDB is not running"
        return 1
    fi
}

# Function to start MongoDB with Docker
start_mongodb_docker() {
    echo "🐳 Starting MongoDB with Docker..."
    if ! command -v docker &> /dev/null; then
        echo "❌ Docker is not installed. Please install Docker first."
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        echo "❌ Docker Compose is not installed. Please install Docker Compose first."
        exit 1
    fi
    
    docker-compose up -d mongodb
    echo "⏳ Waiting for MongoDB to be ready..."
    sleep 5
    
    # Wait for MongoDB to be ready
    for i in {1..30}; do
        if check_mongodb; then
            break
        fi
        echo "⏳ Waiting for MongoDB... ($i/30)"
        sleep 2
    done
    
    if ! check_mongodb; then
        echo "❌ MongoDB failed to start properly"
        exit 1
    fi
}

# Function to install dependencies
install_dependencies() {
    echo "📦 Installing Go dependencies..."
    go mod tidy
    echo "✅ Dependencies installed"
}

# Function to build the application
build_application() {
    echo "🔨 Building the application..."
    mkdir -p bin
    go build -o bin/rest-api main.go
    echo "✅ Application built successfully"
}

# Function to run the application
run_application() {
    echo "🚀 Starting the REST API server..."
    ./bin/rest-api
}

# Main execution
main() {
    install_dependencies
    
    if ! check_mongodb; then
        echo "📡 MongoDB is not running. Starting with Docker..."
        start_mongodb_docker
    fi
    
    build_application
    run_application
}

# Handle script interruption
trap 'echo "⚠️  Shutting down..."; docker-compose down; exit 0' INT TERM

# Execute main function
main
