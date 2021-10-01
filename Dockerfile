# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.17-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Pedre Viljoen (pedre@andcru.io)"

# Set the Current Working Directory inside the container
WORKDIR /go-service

# Copy go mod and sum files
COPY go.mod go.sum ./


# Setup Auth for private modules
COPY ssh_copy/ /root/.ssh
RUN git config --global url."git@github.com:".insteadOf https://github.com/

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY src/ .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 3000

# Run the executable
CMD ["./main"]