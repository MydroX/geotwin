# Geotwin

Technical test for Geotwin

## Install project

```sh
mkdir -p ~/go/src/github.com/MydroX && git clone https://github.com/MydroX/geotwin.git
```

```sh
cd ~/go/src/github.com/MydroX/geotwin
```

Install dev
```sh
echo "GIN_MODE=debug" >> .env
```
```sh
  docker-compose -f deploy/docker-compose.yml -f deploy/dev/docker-compose.dev.yml build
```

Install prod
```sh
echo "GIN_MODE=release" >> .env
```
```sh
  docker-compose -f deploy/docker-compose.yml -f deploy/dev/docker-compose.prod.yml build
```

## Run project

Run dev
```sh
  docker-compose -f deploy/docker-compose.yml -f deploy/dev/docker-compose.dev.yml up
```

Run prod
```sh
  docker-compose -f deploy/docker-compose.yml -f deploy/dev/docker-compose.prod.yml up
```

## Endpoints

### API

Base URL:  http://localhost:8080

| endpoint | description |
|----------|-------------|
| /getPath | Returns "path.geojson" data in JSON format |
| /getDistance | Returns the total distance between all the points (in kilometers) |
| /getDuration | Returns the total duration between all the points (in nanoseconds) |

### Web

URL: http://localhost:8000

## Repositories docker

API: https://hub.docker.com/repository/docker/mydrox/geotwin/general

Web: https://hub.docker.com/repository/docker/mydrox/geotwin-web/general