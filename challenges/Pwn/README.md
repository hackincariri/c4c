# Pwn Challenge Submission

## What pattern should I follow?

- The application can be made in any language or Linux images, as long as it is containerized with Docker.
- The application needs to be in a container with all necessary dependencies.
- The application needs to have a health check route.

## Project Structure

The project structure should follow the following format:

Dependencies

 - Docker
 - Docker swarm
 - Dockerfile

To run your service in Docker swarm, you will need to follow these steps:

```sh
docker swarm init
```

After initializing Docker swarm on your machine, let's go to the steps to make the service available:

Your Docker should use networks created from this command:

```sh

docker network create --driver overlay hackincariri

```
There should be a docker-compose.yml file within the project that follows the following structure:

```yaml

version: "3.8"

services:
  app:
    image: "nome_da_imagem"
    deploy:
      replicas: 1
      restart_policy:
        condition: any
    ports:
      - "8080:8080" # your local port should be externalized
    networks:
      - hackincariri 
    healthcheck:
      test: "curl -f http://localhost:8080/health || exit 1" # your application should return 200 on this path 
      interval: 10s
      timeout: 5s
      retries: 3
      start_period: 20s
networks:
  hackincariri:
    external: true
```

There should be a Dockerfile like the one below:

```Dockerfile
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
```

To run your container you will need to use the following commands:

Generate the Dockerfile image:

```sh
docker build . -t "image_name"
```

Deploy the service in Docker swarm from the docker-compose.yml:

```sh
docker stack deploy -c docker-compose.yml "service_name"
```

### Project folder structure

```sh

├── README.md
└── src
    ├── Dockerfile
    ├── index.js
    └── package.json

```

We recommend following the folder structure described above.

Modify the README.md file with your challenge information.

## How to submit a pwn challenge

To submit your challenge, follow the steps below:

1. Create a private repository with your challenge.

2. Add the email organizacao@hackincariri.com.br as a contributor to the project.

3. Access the form at this [link](https://forms.gle/bnVjrsWELCpWpf1g8).

4. Put the link of the shared project.

5. Select the challenge type Pwn.

6. Put your nickname and your best email.

7. Wait for our contact. ;)


