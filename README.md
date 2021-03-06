# Overview

Replication log is a distributed systems course homework [assignment](https://docs.google.com/document/d/13akys1yQKNGqV9dGzSEDCGbHPDiKmqsZFOxKhxz841U/edit).

# Running the App

```bash
    # tilt (for UI and live updates)
    tilt up

    # or build images
    docker build -f Dockerfile -t node .
    docker build -f Dockerfile.test -t test .

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
    curl -X POST -d '{"message":"abc"}' localhost:8080/messages

    # list messages of master
    curl localhost:8080/messages
    # list messages of secondary-1
    curl localhost:8081/messages
    # list messages of secondary-2
    curl localhost:8082/messages

    # namespaces
    curl localhost:8080/namespaces
    curl -X POST -d '{"message":"local","namespace":"ns"}' localhost:8080/messages
    curl 'localhost:8080/messages?namespace=ns'
    curl 'localhost:8080/messages?namespace=ns&namespace=default'
    # flush all the messages inside the namespace
    curl -X DELETE localhost:8080/namespaces/ns

    # post a message with a 5 second delay (secondaries reply with the delay)
    curl -X POST -d '{"message":"abc","delay":1000}' localhost:8080/messages
```

# Running Tests

```bash
    # using the docker-compose (make sure that the service is up)
    docker-compose run test

    # or locally
    CONFIG_PATH=$(pwd)/test-local.json go test ./integration
```

# Sources

- [Golang installation](https://golang.org/doc/install)
- [Using docker-compose with go](https://docs.docker.com/language/golang/build-images/)
- [Tilt with docker-compose](https://docs.tilt.dev/docker_compose.html)
- [Gin server](https://golang.org/doc/tutorial/web-service-gin)
