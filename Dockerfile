# Build
FROM golang:1.14-alpine3.11 AS build

WORKDIR /sensor-service

COPY . .

RUN apk add build-base
RUN go build -o ./app .

# Deployment
FROM alpine:3.11
EXPOSE 8080

WORKDIR /app
RUN mkdir /app/db

COPY --from=build /sensor-service/app ./
COPY ./db/aranet.db ./db/aranet.db

ENTRYPOINT [ "./app" ]