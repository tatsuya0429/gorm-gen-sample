FROM golang:alpine AS builder

ARG USERNAME=user
ARG GROUPNAME=user
ARG UID=1000
ARG GID=1000
ARG PASSWORD=user

RUN apk update && apk add --no-cache git
RUN addgroup -g ${GID} ${GROUPNAME} && \
    adduser -s /bin/sh -u ${UID} -G ${GROUPNAME} ${USERNAME} -D

# Set the Current Working Directory inside the container
WORKDIR /app

RUN chown -R ${USERNAME}:${GROUPNAME} /app

USER ${USERNAME}

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

FROM builder as dev

RUN go install github.com/cosmtrek/air@latest

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

CMD [ "air", "-c", ".air.toml" ]