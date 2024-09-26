FROM golang:1.23.0-alpine as build-stage

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/service

RUN go build -o server .

FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /app/cmd/service/.env .
RUN ls -a
COPY --from=build-stage /app/cmd/service/server .

EXPOSE 8081

CMD ["./server"]