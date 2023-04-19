# Base image
FROM golang:1.18.9-alpine3.17

# Install dependencies
RUN apk add --no-cache curl bash build-base gcc g++ git musl-dev poppler-utils nodejs-current npm

# Set environment variables
ENV GO111MODULE=on \
    GOBIN="/app/bin" \
    GOPATH="/root/go" \
    PATH="/app/bin:/app/tools/protoc-3.6.1/bin:$PATH" \
    DOCKER_BUILDKIT=1

# Set work directory
WORKDIR /app

# Copy source code
COPY . .

# Build frontend
RUN npm install && npm run dev

# Build backend
RUN go mod download || (echo "Failed to download Go modules" && exit 1) 
RUN go build -o opvault ./vault-web-server

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates

# Expose the port
EXPOSE 8100

# Run the application
CMD ["/app/opvault"]
