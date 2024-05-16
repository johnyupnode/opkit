# Use the official Alpine image as a base
FROM alpine:latest

# Set environment variables for Go
ENV GO_VERSION=1.22.0
ENV PATH=$PATH:/usr/local/go/bin

# Install dependencies
RUN apk update && \
    apk add --no-cache curl wget jq git build-base bash && \
    rm -rf /var/cache/apk/*

# Install Go
RUN wget https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz && \
    rm go$GO_VERSION.linux-amd64.tar.gz

# Install Ignite CLI
RUN curl https://get.ignite.com/cli! | bash

# Create a working directory
WORKDIR /usr/src/app

# Copy all files from the current directory to the working directory in the container
COPY . /usr/src/app

# Make the entrypoint script executable
RUN chmod +x /usr/src/app/entrypoint.sh

# Set the entrypoint
ENTRYPOINT ["/usr/src/app/entrypoint.sh"]
