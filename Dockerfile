# Use the official Go image as a base
FROM golang:1.22.0

# Install dependencies
RUN apt-get update && \
    apt-get install -y curl jq git bash && \
    rm -rf /var/lib/apt/lists/*

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
