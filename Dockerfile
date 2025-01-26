FROM golang:alpine as build 
WORKDIR /
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go get .
RUN go build -o /usr/local/bin/app

FROM alpine:latest
COPY --from=build /usr/local/bin/app /usr/local/bin/app

ARG VERSION
ENV VERSION=${VERSION}
CMD ["app"]
EXPOSE 8090