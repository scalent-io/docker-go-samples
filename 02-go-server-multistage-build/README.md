# Multi-Stage Docker Image with Go web-server


## Create Multi-Stage Docker image 
Create a Multi-Stage docker image using Dockerfile, the syntax will be like "docker build . -t image-name:tag-name"

```bash
docker build . -t godocker-app-02:goDockerMultistage
```

## List Docker images 
```bash
docker images
```
  

## Running the docker image 
Run the docker image as a docker container, syntax "docker run -p HOST-PORT:CONTAINER-PORT IMAGE-ID" 
```bash
docker run -p 8080:80 9b7626897732
```

## Visit the link 
Open browser and visit [http://localhost:8080/hi](http://localhost:8080/hi)