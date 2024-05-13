# Stage 1: Build stage
FROM golang:1.22-alpine AS build

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -o rapid-go .

# Stage 2: Final stage
FROM alpine:edge as server

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/rapid-go .

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

# Set the entrypoint command
ENTRYPOINT ["/app/rapid-go"]