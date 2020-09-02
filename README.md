# go-web-docker
Example Golang web app with Docker support

## Build Docker image

```bash
docker image build -t go-web-docker:v1.0.0 -f build/Dockerfile .
```

## Run Docker container

```bash
docker container run -d --name go-web-docker-1 -p 8080:8080 go-web-docker:v1.0.0
```

## Verify container is uprunning

```bash
curl localhost:8080
```

You should see

```bash
Hello world from Go Web Server!
```