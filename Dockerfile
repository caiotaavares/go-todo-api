FROM golang:1.18 AS builder

RUN mkdir -p /app
WORKDIR /app

COPY . .

RUN go build -o server

FROM gcr.io/distroless/static

COPY --from=builder /app/server /server

EXPOSE 8080

ENTRYPOINT [ "/server" ]