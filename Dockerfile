## Build
FROM golang:1.19.4-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

RUN go build -o /golang-gin

## Deploy
FROM gcr.io/distroless/base-debian10:debug

WORKDIR /

COPY --from=build /golang-gin /golang-gin

ENV PORT=8080
EXPOSE ${PORT}

USER nonroot:nonroot

ENTRYPOINT ["/golang-gin"]