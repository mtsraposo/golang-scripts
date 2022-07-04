# syntax=docker/dockerfile:1

FROM golang:1.18 AS build

WORKDIR /app

COPY go.sum go.mod ./
RUN go install -v ./...

COPY . ./

RUN go build -o /scripts

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /scripts /scripts

ENTRYPOINT [ "/scripts" ]
