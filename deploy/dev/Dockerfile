FROM golang:1.17.0-alpine3.13

ARG WORK_DIR

WORKDIR $WORK_DIR

COPY . .

RUN apk add --no-cache ca-certificates \
    bash \
    sudo \
    git 

RUN go install github.com/cosmtrek/air@v1.27.3
RUN go install github.com/mgechev/revive@v1.1.1

ENTRYPOINT [ "./entrypoint.sh" ]