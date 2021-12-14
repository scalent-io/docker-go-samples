# from golang 1.16 version with alpine linux which is light weight,  naming it as a "build"
FROM golang:1.16-alpine  AS build

# creating directory in the container withh name /app, this will be user as base diretory 
WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# copy all go files inside the /app directory 
COPY *.go ./

# by default CGO is desabled in alpine image
ENV CGO_ENABLED=0

# creating golang build and naming the build as "go-mysql-app" wit -o flag 
RUN go build -o /go-mysql-app


#
## Multistage build to reduce the image size 

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-mysql-app /go-mysql-app


USER nonroot:nonroot

# running golang binary inside the container
ENTRYPOINT ["/go-mysql-app"]
