FROM node:14.17.5-alpine3.14

WORKDIR /app

COPY ./web .

RUN apk add --no-cache ca-certificates \
    bash \
    sudo \
    git \
    yarn=1.22.10-r0

ENTRYPOINT [ "./entrypoint.sh" ]