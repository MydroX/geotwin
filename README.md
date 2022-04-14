# Geotwin

Technical test for Geotwin
 
Le format de chaque coordonnées est un tableau de 4 éléments correspondant à :
[
            longitude,
            latitude,
            altitude,
            timestamp
]

L'objectif est de créer une application Web exposant les endpoint REST suivants :

http://<url>/getPath (retourne la donnée telle quelle au format json)
http://<url>/getDistance (retourne la distance totale du parcours)
http://<url>/getDuration  (retourne la durée totale du déplacement, en nanosecondes)

La distance et durée doivent être calculées lors de l'appel au endpoint.

Vous pouvez développer cette application dans le langage de votre choix. Le résultat du test doit être fourni sur un dépôt Github contenant un README expliquant comment construire et lancer le serveur et l'adresse exacte de chaque endpoint.

Bonus 1 : l'application est buildée avec Docker et fournie sur Dockerhub en plus du code sur Github

Bonus 2 : vous fournissez également une application Javascript capable d'afficher sur une carte le déplacement via un appel à votre endpoint getPath (comme sur le serveur web : github + README + éventuellement dockerhub)

Bonus 3 : les deux applications peuvent être lancées ensemble via docker compose

## Install project

```sh
mkdir -p ~/go/src/github.com/MydroX
```

```sh
cd ~/go/src/github.com/MydroX
```

```sh
git clone https://github.com/MydroX/geotwin.git
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
