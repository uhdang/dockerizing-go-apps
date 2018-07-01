#Using official golang image

## Build image
$ docker build -t [IMAGENAME].

## Run image as a container
$ docker run -d -p 3000:3000 --name [CONTAINERNAME] [IMAGENAME] 
