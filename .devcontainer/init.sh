#!/bin/bash

# Install dependencies
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/air-verse/air@latest

go mod tidy 
npm install --prefix apps/mobile 
npm install --prefix apps/web