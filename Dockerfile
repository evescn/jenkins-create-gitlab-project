FROM golang:1.18.10-alpine3.17

WORKDIR /data
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go && \
    mkdir /app && \
    mv ./main /app

WORKDIR /app
CMD ["tail", "-f", "/etc/fstab"]
