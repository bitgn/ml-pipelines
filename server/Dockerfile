# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Rinat Abdullin <rinat@abdullin.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -a -installsuffix cgo -o main .

# pack the resources
RUN find ./web -name "*.html" -o -name "*.css" | tar -cf pack -T -


######## Start a new stage from scratch #######
FROM ubuntu:bionic

WORKDIR /root/

RUN mkdir db

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# import and extract the resources
COPY --from=builder /app/pack .
RUN tar -xvf pack && rm pack


# Expose port 8080 to the outside world
EXPOSE 80
EXPOSE 9111

# Command to run the executable
CMD ["./main","-web=:80","-grpc=:9111"]