#!/bin/bash

# Install dependencies
go install -v github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install -v github.com/ogen-go/ogen/cmd/ogen@latest
go install -v github.com/air-verse/air@latest

go mod tidy 
npm install --prefix apps/mobile 
npm install --prefix apps/web