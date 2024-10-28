#!/bin/bash


export POSTGRES_USER=user
export POSTGRES_PASSWORD=password
export POSTGRES_PORT=6560
export POSTGRES_HOST=127.0.0.1
export POSTGRES_NAME=file_manager

go run main.go

# Exit