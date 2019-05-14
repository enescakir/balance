FROM golang:1.12 as builder

LABEL maintainer="Enes Çakır <enes@cakir.web.tr>"

WORKDIR /go/src/github.com/enescakir/balance

COPY . .

# Download dependencies
RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/server ./server


# FINAL IMAGE
FROM alpine:latest

# Dockerize is used for waiting MySQL container to start and run
ENV DOCKERIZE_VERSION v0.6.1
RUN apk add --no-cache openssl ca-certificates \
    && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

RUN mkdir /app

WORKDIR /app

COPY server/templates ./templates

COPY --from=builder /go/bin/server .

EXPOSE 8080

CMD ["./server"]
