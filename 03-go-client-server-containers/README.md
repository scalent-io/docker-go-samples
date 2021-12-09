# Multi-Stage Docker Image with Go web-server


## Create Go-Server Docker image 
Creating docker go-server using docker file Dockerfile.server, notice there is a dot at the end of command

```bash
docker build -t go-server:goServerTag -f Dockerfile.server .
```


## Create Go-client Docker image 
Creating docker go-client using docker file Dockerfile.client

```bash
docker build -t go-client:goClientTag -f Dockerfile.client .
```


## List Docker images 
```bash
docker images
```

## Create a network for communicating between containers  
```bash
docker network create mynet
```

## List Docker Networks 
```bash
docker images
```


## Running the docker images 
naming the server as "goserver.host" so that we can call using the name, also added to the same network 
```bash
docker run --name goserver.host --network mynet -it go-server:goServerTag
```

Run go client

```bash
docker run -p 8090:8090 --network mynet  -it go-client:goClientTag
```

## Visit the link 
Open browser and visit [http://localhost:8080/hi](http://localhost:8080/hi)