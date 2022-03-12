
FROM golang:1.14 as builder

WORKDIR /app

COPY . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o server *.go

FROM scratch
WORKDIR /opt
COPY --from=builder /app/server .
EXPOSE 8080
ENTRYPOINT ["/opt/server"]