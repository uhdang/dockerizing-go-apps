#Using official golang image

## Build image
$ docker build -t [IMAGENAME].

## Run image as a container
$ docker run -d -p 3000:3000 --name [CONTAINERNAME] [IMAGENAME] 

## what does this do?
$ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


