FROM golang:alpine as build-env
COPY . /src
WORKDIR /src/demo/module
RUN go build -o module

FROM alpine:3.19.1
WORKDIR /app
COPY --from=build-env /src/demo/module/module /app/
ENTRYPOINT ["./module"]
