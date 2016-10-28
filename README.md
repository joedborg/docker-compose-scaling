# Docker Compose DNS Load Balancing Example
This is to demonstrate the Docker Compose `scale` command and how it can be
used with Nginx for instant scaling.

## Building
### Requirements
* Go
* Docker Compose
* GNU Make

To build, first compile the Go web application server.  This will simply
respond with the container's ID.

```
$ make
$ docker-compose up -d
```
