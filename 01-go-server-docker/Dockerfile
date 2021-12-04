# from golang 1.16 version with alpine linux which is light weight 
FROM golang:1.16-alpine

# creating directory in the container withh name /app, this will be user as base diretory 
WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# copy all go files inside the /app directory 
COPY *.go ./

# creating golang build and naming the build as "go-docker-app" wit -o flag 
RUN go build -o /go-docker-app

# expose port 8080 to outside application 
EXPOSE 8080

# running golang binary inside the container
CMD ["/go-docker-app"]