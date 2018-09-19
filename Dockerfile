FROM golang:1.11.0-alpine3.8

WORKDIR /app
COPY . .

RUN go build -o hello .

ENV PORT=80

CMD ["/app/hello"]