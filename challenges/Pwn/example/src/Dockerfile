# Use a base Ubuntu image
FROM ubuntu:18.04

# Update packages and install necessary dependencies
RUN apt-get update && \
    apt-get install -y \
    wget \
    curl \
    git \
    nano \
    vim \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# Create the working directory
WORKDIR /home/hackincariri/app

# Download the Go file with proper permissions
RUN wget -O go.tar.gz https://golang.org/dl/go1.21.1.linux-amd64.tar.gz

# Extract the Go file
RUN tar -C /home/hackincariri -xzf go.tar.gz && \
    rm go.tar.gz

# Set environment variables for Go
ENV PATH $PATH:/home/hackincariri/go/bin
ENV GOPATH /home/hackincariri/go

# Copy source code to the container
COPY . .

# Compile the Go code
RUN go build -o app -buildvcs=false

# Remove the Dockerfile after compilation
RUN rm -f Dockerfile
RUN rm -f go.mod
RUN rm -f README.md
RUN rm -f main.go

# Add a read-only file as root
USER root
RUN touch /root/flag_pwn.txt && chmod 400 /root/flag_pwn.txt
RUN echo "HIK_PWN_6630db853468e9c768a584981349e924" > /root/flag_pwn.txt

# Switch back to non-root user
USER hackincariri

# Command to start the SSH server and Go server when the container is executed
CMD ./app