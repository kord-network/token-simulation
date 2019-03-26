#!/bin/bash

# Build the Go binary
go build -o bin/simulate cli.go

# Copy to local bin directory
cp bin/simulate /usr/local/bin
