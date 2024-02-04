# syntax=docker/dockerfile:1

# It imports go image 
# used because we need a go image to build a go application.

FROM golang:1.19 AS build-stage

# It sets the work directory of image as DogGo
# used same as the name of repository

WORKDIR /DogGo

# It copies all the require file  and packages 
# used to copy the content of the files to work directories from the project directory.

COPY go.mod go.mod
COPY cmd/ cmd/
COPY pkg/ pkg/
COPY ui/ ui/

# Commmand to run build the executables using linux as os
# to build the executable from the ./cmd/web
# used cgo_enabled=0 to make static execution rather than using c libraries

RUN CGO_ENABLED=0 GOOS=linux go build -a -o web ./cmd/web

# COPY /DogGo/web /usr/local/bin/web

# it imports the distroless image for multistage build
# used distroless to reduce the size and increase the security

FROM gcr.io/distroless/base-debian11 AS build-release-stage

# Made work directories for multistage build 

WORKDIR /

# copied the executables and files from the build stage to multistage build directories

COPY --from=build-stage /DogGo/web /web
COPY --from=build-stage /DogGo/ui /ui

# Expose the executable to 8888 port

EXPOSE 8888

# no root user to build docker image such that no one gets the root of image
USER nonroot:nonroot

# command to run when we build the docker image

ENTRYPOINT [ "./web" ]
