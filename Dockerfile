# Stage 1: Build Go backend
FROM golang:1.19 AS go-builder

WORKDIR /go/src/app

# Set environment variables for Go
ENV GO111MODULE=on \
    GOBIN=/go/src/app/bin \
    GOPATH=/root/go \
    PATH=$PATH:/go/src/app/bin:/go/src/app/tools/protoc-3.6.1/bin \
    DOCKER_BUILDKIT=1

# Copy Go files and dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the Go source files
COPY errorlist errorlist/
COPY form form/
COPY serverutil serverutil/
COPY validator validator/
COPY vault-web-server vault-web-server/


# Build the Go application
RUN go build -o main ./vault-web-server 

# Stage 2: Build React frontend
FROM node:19 AS js-builder

WORKDIR /app

# Copy JavaScript files and dependencies
COPY package*.json ./
RUN npm install

# Copy the React source files and other assets
COPY components components/
COPY config config/
COPY static static/
COPY webpack.config.js .

# Build the React frontend
RUN npx webpack --config webpack.config.js

# Stage 3: Assemble the final image
FROM debian:buster-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=go-builder /go/src/app/main /app/main
COPY --from=js-builder /app/static /static

COPY config/websites.json /config/websites.json

EXPOSE 8100

CMD ["/app/main"]
