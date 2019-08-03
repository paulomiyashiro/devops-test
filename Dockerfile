FROM golang:latest

LABEL maintainer="Paulo Miyashiro <paulo.miyashiro@gmail.com>"

WORKDIR /app

COPY src/. .

COPY server.crt /etc/ssl/certs/server.crt

COPY server.key /etc/ssl/private/server.key

RUN go build -o main .

EXPOSE 8080

EXPOSE 8443

CMD ["./main"]