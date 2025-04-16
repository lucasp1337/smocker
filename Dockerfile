#Backend build stage
ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine AS build-backend
RUN apk add --no-cache make
ARG VERSION=snapshot
ARG COMMIT
WORKDIR /go/src/smocker
COPY go.mod go.sum ./
RUN go mod download
COPY Makefile main.go ./
COPY server/ ./server/
RUN make VERSION=$VERSION COMMIT=$COMMIT RELEASE=1 build

# Frontend build stage
FROM node:16 AS build-frontend
WORKDIR /app
COPY package.json yarn.lock ./
COPY client/ ./client/
RUN yarn install --frozen-lockfile
RUN yarn build

# Final stage
FROM alpine
LABEL org.opencontainers.image.source="https://github.com/smocker-dev/smocker"
WORKDIR /opt
EXPOSE 8080 8081
# Copy the built frontend from the frontend build stage
COPY --from=build-frontend /app/build/client client/
# Copy the built backend from the backend build stage
COPY --from=build-backend /go/src/smocker/build/* /opt/
CMD ["/opt/smocker"]
