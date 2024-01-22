###
# Stage: builder
# Install dependencies and build agate binary.
# Output:
#  /app/agate
###
FROM golang:1.21.6-alpine3.19 AS builder

# Define working directory.
WORKDIR /app

# Install build dependencies.
RUN apk add build-base

# Install Golang dependencies.
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

# Copy source code.
COPY . .

# Build binary.
RUN go build -o /app/agate ./main.go

###
# Stage: app
# Expose agate binary entrypoint in a light image.
###
FROM alpine:3.19 AS app

# Define workding directory
WORKDIR /app

# Copy binary from builder stage to alpine image.
COPY --from=builder /app/agate /app/agate

# Copy database migration files to alpine image.
COPY --from=builder /app/db /app/db

# Set entrypoint to  binary.
ENTRYPOINT ["/app/agate"]