# Development environment with hot-reloading

# Backend build stage
ARG GO_VERSION=1.23
FROM golang:${GO_VERSION}-alpine AS local-backend

RUN apk add --no-cache make git curl
RUN go install github.com/cespare/reflex@latest

# Set working directory
WORKDIR /go/src/smocker

# Copy only the necessary files for dependency installation
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY Makefile main.go ./
COPY server/ ./server/

# Frontend build stage
FROM node:16 AS local-frontend
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY client/ ./client/
RUN yarn build

# Final stage
FROM alpine
WORKDIR /opt

# Install required tools for development
RUN apk add --no-cache go git make curl
RUN go install github.com/cespare/reflex@latest

# Set environment
ENV LEVEL=debug

# Copy backend code and dependencies
COPY --from=local-backend /go/src/smocker /opt/smocker
COPY --from=local-backend /go/bin/reflex /go/bin/reflex

# Copy the built frontend
COPY --from=local-frontend /app/build/client /opt/smocker/build/client

# Set working directory
WORKDIR /opt/smocker

# Persistence directory
RUN mkdir -p /opt/smocker/sessions

EXPOSE 8080 8081

# Command to start the application with hot-reloading
CMD ["sh", "-c", "/go/bin/reflex --start-service --decoration=none --regex='\\.go$' --inverse-regex='^vendor|node_modules|\\.cache/' -- go run main.go --log-level=${LEVEL} --static-files ./build/client --persistence-directory ./sessions"]
