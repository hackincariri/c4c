# Web Challenge Submission

## What pattern should I follow?

- The application can be made in any language, as long as it is containerized with Docker.
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
    image: "image_name"
    deploy:
      replicas: 1
      restart_policy:
        condition: any
    ports:
      - "3000:3000" # your local port should be externalized
    networks:
      - hackincariri 
    healthcheck:
      test: "curl -f http://localhost:3000/health || exit 1" # your application should return 200 on this path 
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

FROM node:20

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm", "start"]
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

## Project folder structure

```sh

├── README.md
└── src
    ├── Dockerfile
    ├── index.js
    └── package.json

```

We recommend following the folder structure described above.

Modify the README.md file with your challenge information.

## How to submit a web challenge

To submit your challenge, follow the steps below:

1. Create a private repository with your challenge.

2. Add the email organizacao@hackincariri.com.br as a contributor to the project.

3. Access the form at this [link](https://forms.gle/bnVjrsWELCpWpf1g8).

4. Put the link of the shared project.

5. Select the challenge type Web.

6. Put your nickname and your best email.

7. Wait for our contact. ;)