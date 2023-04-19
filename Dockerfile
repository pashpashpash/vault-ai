# Stage 1: Build Go backend
FROM golang:1.18.9-alpine3.17 AS backend-builder

# Install necessary build dependencies
RUN apk add --no-cache build-base gcc g++ musl-dev

WORKDIR /go/src/app

# Set Go environment variables taken from scripts/source-me.sh
ENV GO111MODULE=on \
    GOBIN=/go/src/app/bin \
    GOPATH=/root/go \
    PATH=$PATH:/go/src/app/bin:/go/src/app/tools/protoc-3.6.1/bin \
    DOCKER_BUILDKIT=1

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy Go source files
COPY errorlist errorlist/
COPY form form/
COPY serverutil serverutil/
COPY validator validator/
COPY vault-web-server vault-web-server/

# Build Go application
RUN go build -o main ./vault-web-server

# Stage 2: Build React frontend
FROM node:19-alpine3.17 AS frontend-builder

WORKDIR /app

# Copy JavaScript package files and install dependencies
COPY package*.json ./
RUN npm install --ignore-scripts


# Copy React source files and other assets
COPY components components/
COPY config config/
COPY static static/
COPY webpack.config.js .

# Build React frontend
RUN npx webpack --config webpack.config.js

# Stage 3: Assemble the final image
FROM alpine:3.17

# Install only required runtime dependencies
RUN apk add --no-cache ca-certificates poppler-utils

# Copy Go binary from the backend-builder stage
COPY --from=backend-builder /go/src/app/main /app/main

# Copy static files from the frontend-builder stage
COPY --from=frontend-builder /app/static /static
COPY config/websites.json /config/websites.json

# Copy scripts directory
COPY scripts scripts/

# Expose the application port
EXPOSE 8100

# Set the file permission for go-compile.sh and execute Go binary
RUN chmod +x scripts/go-compile.sh
CMD ["/app/main"]

