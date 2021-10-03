# Overview

Replication log is a distributed systems course homework [assignment](https://docs.google.com/document/d/13akys1yQKNGqV9dGzSEDCGbHPDiKmqsZFOxKhxz841U/edit).

# Running the App

```bash
    # use docker-compose
    docker-compose up

    # or tilt (for UI and live updates)
    tilt up
```

# Using the App

```bash
    # master
    curl localhost:8080/ping
    # secondary-1
    curl localhost:8081/ping
    # secondary-2
    curl localhost:8082/ping
```
