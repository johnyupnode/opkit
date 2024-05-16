# Use the official Ubuntu image as a base
FROM ubuntu:22.04

# Set environment variables for Go
ENV GO_VERSION=1.21.10
ENV PATH=$PATH:/usr/local/go/bin

# Install dependencies
RUN apt-get update && \
    apt-get install -y curl wget jq git && \
    rm -rf /var/lib/apt/lists/*

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
