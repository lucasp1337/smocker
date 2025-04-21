# CI environment for running tests - without frontend build

ARG GO_VERSION=1.23
FROM golang:${GO_VERSION}-alpine

# Install dependencies for testing
RUN apk add --no-cache make git curl lsof gcc musl-dev
RUN go install github.com/ovh/venom/cmd/venom@v1.0.0-rc.6
RUN go install github.com/wadey/gocovmerge@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8

# Set working directory
WORKDIR /go/src/smocker

# Copy source code and dependencies
COPY . .

# Create a directory for test coverage reports
RUN mkdir -p coverage

# Create directory for persistence
RUN mkdir -p /go/src/smocker/sessions

# Expose ports for API and UI during tests
EXPOSE 8080 8081

# Set environment for tests
ENV SMOCKER_PERSISTENCE_DIRECTORY=/go/src/smocker/sessions

# Default command for running tests
CMD ["sh", "-c", "make test && make test-integration"]
