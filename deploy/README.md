# Docker instructions

## Running project

### Dev mode
```sh
docker-compose -f deploy/docker-compose.yml  -f deploy/dev/docker-compose.dev.yml up
```
### Prod mode
```sh
docker-compose -f deploy/docker-compose.yml  -f deploy/prod/docker-compose.dev.yml up
```