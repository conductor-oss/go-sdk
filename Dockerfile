FROM golang:1.23 AS build
RUN mkdir /package
COPY /sdk /package/sdk
COPY /go.mod /package/go.mod
COPY /go.sum /package/go.sum
WORKDIR /package
RUN go build -v ./...

FROM build as test
COPY /test/unit_tests /package/test/unit_tests
# Run SDK unit tests
RUN go test -v -race ./sdk/...
# Run additional unit tests
RUN go test -v -race ./test/unit_tests/...

FROM build as inttest
COPY /test /package/test
ARG CONDUCTOR_AUTH_KEY
ARG CONDUCTOR_AUTH_SECRET
ARG CONDUCTOR_SERVER_URL
ENV CONDUCTOR_AUTH_KEY=${CONDUCTOR_AUTH_KEY}
ENV CONDUCTOR_AUTH_SECRET=${CONDUCTOR_AUTH_SECRET}
ENV CONDUCTOR_SERVER_URL=${CONDUCTOR_SERVER_URL}
RUN go test -v ./test/integration_tests/...

FROM build AS harness-build
COPY /harness /package/harness
WORKDIR /package
RUN CGO_ENABLED=0 go build -o /app/harness ./harness

FROM alpine:3 AS harness
RUN adduser -D -u 65532 nonroot
USER nonroot
COPY --from=harness-build /app/harness /app/harness
WORKDIR /app
ENTRYPOINT ["/app/harness"]