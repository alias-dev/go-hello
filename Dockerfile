FROM golang:1.11.0-alpine3.8
WORKDIR /app
COPY . .
RUN go build -o hello .

FROM alpine:latest  
WORKDIR /root/
COPY --from=0 /app/greeting.tmpl.html .
COPY --from=0 /app/hello .
ENV PORT=80
CMD ["./hello"]