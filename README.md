# APT Gateway

![](https://img.shields.io/badge/language-golang_v1.20-lightblue)
![](https://img.shields.io/badge/tests-passed-green)
![](https://img.shields.io/badge/version-0.1.1-red)

Gateway is the ```apt``` restful API for handling client http requests.
This module handles the user interface logic in order to communicate with
```core``` and ```ftp```. It is the system main gateway app.

## Image

Gateway app docker image address:

```shell
docker pull amirhossein21/apt-gateway:v0.1.1
```

### configs

Make sure to create ```config.yml``` file with the following variable init:

```yaml
http:
  port: 8080
  core: 'http://localhost:9090/api'
  dev_mode: true
jwt:
  private_key: 'super'
  expire_time: 180 # minute
mysql:
  host: 'localhost'
  port: 3306
  user: root
  pass: ''
  database: 'apt'
  migrate: false
ftp:
  host: 'http://localhost:9091'
  secret: 'secret'
  access: 'access'
```

## Setup

Setup ui application in docker container with following command:

```shell
docker run -d \
  -v type=bind,source=$(pwd)/config.yml,dest=/app/config.yml \
  -p 80:80 \
  amirhossein21/apt-gateway:v0.1.1
```
