FROM golang:1.16.5 as builder
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN git rev-parse --short HEAD
RUN GIT_COMMIT=$(git rev-parse --short HEAD) && \
    CGO_ENABLED=0 go build -o /iam-service-login -ldflags "-X main.GitCommit=${GIT_COMMIT}"

# Container image that runs your code
FROM alpine:3.14.0

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY --from=builder /iam-service-login /iam-service-login
COPY entrypoint.sh /entrypoint.sh

# Code file to execute when the docker container starts up (`entrypoint.sh`)
ENTRYPOINT ["/entrypoint.sh"]
