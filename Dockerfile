# builder image
FROM golang:1.17.1 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV GOOS linux
ENV GOARCH ${GOARCH:-amd64}
ENV CGO_ENABLED=0

RUN go build -v -o ngps-api cmd/ngps-api/main.go

# generate clean, final image for end users

FROM alpine:3.14

ARG VERSION
ENV APP_VERSION=$VERSION
ENV GIN_MODE=release

COPY --from=builder /build/ngps-api /usr/local/bin/ngps-api

EXPOSE 8080

ENTRYPOINT [ "ngps-api" ]
