FROM golang:1.16.5 as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /iam-service-login

# Container image that runs your code
FROM alpine:3.14.0

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY --from=builder /build/iam-service-login /iam-service-login
COPY entrypoint.sh /entrypoint.sh

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/entrypoint.sh"]
