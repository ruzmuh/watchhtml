# Stage 1: Build the app from source
FROM golang:1.20 as builder

WORKDIR /app

# Copy the source code into the container
COPY . .

ENV GOOS=linux
ENV GARCH=amd64 
ENV CGO_ENABLED=0 
# Build the app
RUN go build -o watchhtlm

# Stage 2: Use a lightweight base image for the final image
FROM alpine:latest

# Copy the built binary from the builder stage
COPY --from=builder /app/watchhtlm /usr/local/bin/watchhtlm

# Run the app
ENTRYPOINT ["watchhtlm"]
