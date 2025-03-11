#!/bin/bash

echo "Setting up Backend Image Processing Service..."

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

# Install dependencies
echo "Installing dependencies..."
go mod tidy

# Build the project
echo "Building project..."
go build -o backend-image-service

# Run the server
echo "Starting the server..."
./backend-image-service
