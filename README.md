# Overview

Replication log is a distributed systems course homework [assignment](https://docs.google.com/document/d/13akys1yQKNGqV9dGzSEDCGbHPDiKmqsZFOxKhxz841U/edit).

# Running the App

```bash
    # tilt (for UI and live updates)
    tilt up --hud=false

    # or build images
    docker build -f Dockerfile.master -t master .
    docker build -f Dockerfile.secondary-1 -t secondary-1 .
    docker build -f Dockerfile.secondary-2 -t secondary-2 .

    # and use docker-compose
    docker-compose up
```

# Using the App

```bash
    # ping master
    curl localhost:8080/ping
    # ping secondary-1
    curl localhost:8081/ping
    # ping secondary-2
    curl localhost:8082/ping

    # post a message (master only)
    curl -X POST -d '123' localhost:8080/messages

    # list messages of master
    curl localhost:8080/messages
    # list messages of secondary-1
    curl localhost:8081/messages
    # list messages of secondary-2
    curl localhost:8082/messages
```

# Running Tests

```bash
    # using the docker-compose (make sure that the service is up)
    docker-compose run test
```

# Sources

- [Golang installation](https://golang.org/doc/install)
- [Using docker-compose with go](https://docs.docker.com/language/golang/build-images/)
- [Tilt with docker-compose](https://docs.tilt.dev/docker_compose.html)
- [Gin server](https://golang.org/doc/tutorial/web-service-gin)
