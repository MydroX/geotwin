FROM golang:1.17.0

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN go build ./main.go

CMD [ "./main"]