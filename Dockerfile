FROM golang:1.19-alpine as builder
WORKDIR /app
ENV GO111MODULE=on

ARG version
ENV IMAGE_STORAGE_APP_VERSION=$version

COPY . .
RUN go mod download
RUN go build -o image_storage_server cmd/server/main.go

FROM alpine
RUN apk update && \
    apk add bash && \
    adduser -D -H -h /app server && \
    mkdir /app && \
    chown -R server:server /app
WORKDIR /app
USER server

COPY --chown=server --from=builder /app/image_storage_server /app

EXPOSE 5000
CMD ["/app/image_storage_server"]