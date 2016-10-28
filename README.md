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
```

## Running
```
$ docker-compose up -d
```
Then browse to [http://localhost](http://localhost) and see the name of the single application
container.

## Scaling
To scale the application server, run
```
$ docker-compose scale application=2
```
(or any number of instances).  Then refresh the browser to see the container
name change as Docker's DNS resolver does round-robin.
