# My Golang Gin Sample API #
This project is for learning Golang (with Gin-gonic)

[![ci](https://github.com/bankierubybank/golang-gin/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/bankierubybank/golang-gin/actions/workflows/main.yml)

##### Environment Variable
| Variable name | Description | Default | Mandatory |
| ------ | ------ | ------ | ------ |
| PORT | Application Port | 8080 | YES |

### Setup (Docker)
```
docker build -t golang-gin:latest .
docker run --name my-golang-gin -d -p 80:8080 golang-gin:latest
```

### Landing Page
```
http://<host-ip>:8080/
```

### API Routes
```
http://<host-ip>/albums/
```