# hello-go

## Docker
```bash
# Build the image
docker build -t go-hello .

# Run it
docker run -it --rm -p 8080:80 --env GREETING='Hello, Docker!' --name go-hello-app go-hello
````