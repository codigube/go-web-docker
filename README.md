# go-web-docker

Example Golang web app with Docker support

## Build Docker image

```bash
docker image build -t go-web-docker:v1.0.0 -f build/Dockerfile .
```

## Run Docker container

```bash
docker container run -d --name go-web-docker-1 --env GO_ENV=production --env PORT=8081 -p 8081:8081 go-web-docker:v1.0.0
```

## Verify container is up-running

```bash
curl localhost:8080
```

You should see

```bash
Hello world from Go Web Server!
```
