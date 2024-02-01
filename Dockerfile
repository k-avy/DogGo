# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /DogGo

COPY go.mod go.mod
COPY cmd/ cmd/
COPY pkg/ pkg/
COPY ui/ ui/

RUN CGO_ENABLED=0 GOOS=linux go build -a -o web ./cmd/web

# COPY /DogGo/web /usr/local/bin/web

EXPOSE 8888

ENTRYPOINT [ "./web" ]