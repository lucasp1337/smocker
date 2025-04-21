# Backend build stage base
ARG GO_VERSION=1.23
FROM golang:${GO_VERSION}-alpine AS build-backend-base
RUN apk add --no-cache make git
WORKDIR /go/src/smocker
COPY go.mod go.sum ./
RUN go mod download
COPY Makefile main.go ./

# Frontend build stage base
FROM node:16 AS build-frontend-base
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile --ignore-scripts
