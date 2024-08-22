FROM golang:1.22-alpine AS build-env

ENV APP_NAME_CRON=service-go-rest-simple
ENV TZ=Asia/Jakarta

RUN apk update && apk upgrade && \
    apk add --no-cache --virtual .build-deps --no-progress -q \
    curl \
    busybox-extras \
    make \
    git \
    tzdata

ADD . /app
WORKDIR /app

RUN ./.generate-migration-config

# run migration
RUN go install github.com/gobuffalo/pop/soda@latest

RUN soda migrate up

RUN go build -o main main.go
RUN go build -o $APP_NAME_CRON src/interface/cli/cli.go

RUN chmod 755 ./main
RUN chmod 755 ./$APP_NAME_CRON

EXPOSE ${HTTP_PORT}

CMD "./main"
