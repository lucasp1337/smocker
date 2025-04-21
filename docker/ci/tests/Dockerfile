# CI environment for running tests - without frontend build - TODO: remake

ARG GO_VERSION=1.23
FROM golang:${GO_VERSION}-alpine

RUN apk add --no-cache make git curl lsof gcc musl-dev
RUN go install github.com/ovh/venom/cmd/venom@v1.0.0-rc.6
RUN go install github.com/wadey/gocovmerge@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8

WORKDIR /go/src/smocker

COPY . .

RUN mkdir -p coverage

RUN mkdir -p /go/src/smocker/sessions

EXPOSE 8080 8081

ENV SMOCKER_PERSISTENCE_DIRECTORY=/go/src/smocker/sessions

CMD ["sh", "-c", "make test && make test-integration"]
