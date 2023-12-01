# Pre-built binary
FROM philipslabs/iam-service-login:latest as binary
# TODO: add signature checking at some point

FROM alpine:3.18.5
COPY --from=binary /iam-service-login /iam-service-login
ENTRYPOINT ["/iam-service-login"]
