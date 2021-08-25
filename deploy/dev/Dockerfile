FROM golang:1.16.0-alpine3.13

ARG WORK_DIR

WORKDIR $WORK_DIR

COPY . .

RUN apk add --no-cache ca-certificates \
    bash \
    sudo \
    git 

RUN go get -u github.com/cosmtrek/air
RUN go get -u github.com/mgechev/revive

ENTRYPOINT [ "/go/src/github.com/MydroX/geotwin/deploy/dev/entrypoint.sh" ]