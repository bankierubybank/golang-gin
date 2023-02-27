# My Golang Gin Sample API #
This project is for learning Golang (with Gin-gonic)
- How to write sample API with Golang
- How to Dockerize Golang application
- How to CI using GitHub Actions
- How to deploy Golang container on Docker and K8s/OpenShift

[![CI](https://github.com/bankierubybank/golang-gin/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/bankierubybank/golang-gin/actions/workflows/main.yml)

##### Environment Variable
| Variable name | Description | Default | Mandatory |
| ------ | ------ | ------ | ------ |
| PORT | Application Port | 8080 | YES |

### Setup (Docker) - Pull my image
```
docker push bankierubybank/golang-gin:latest
docker run --name my-golang-gin -d -p 80:8080 bankierubybank/golang-gin:latest
```

### Setup (Docker) - Build From Source
```
docker build -t golang-gin:latest .
docker run --name my-golang-gin -d -p 80:8080 golang-gin:latest
```

### Swagger
```
http://<HOST-IP>:<PORT>/swagger/index.html
```

##### API Routes
| Method | Route | Purpose |
| ------ | ------ | ------ |
| GET | /api/v1/albums | List albums |
| GET | /api/v1/albums/{id} | Get an album by ID |
| POST | /api/v1/albums | Create an album from provided JSON |
| GET | /api/v1/debug | Get debug information |