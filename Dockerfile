FROM docker.io/golang:1.21-alpine as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY ./ /usr/src/app/
RUN go build -ldflags -w -o ./bin/learn-auth-golang .

FROM docker.io/alpine:latest as run

RUN apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /apps
COPY --from=builder /usr/src/app/bin /apps/
COPY .env /apps/

ENTRYPOINT ["/apps/learn-auth-golang"]
