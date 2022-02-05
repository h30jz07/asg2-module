# syntax=docker/dockerfile:1

FROM golang:alpine3.13 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o module-backend server.go

FROM alpine:3.13 
WORKDIR /app
COPY --from=builder /app/module-backend .
EXPOSE 9061

CMD [ "/app/module-backend" ]

#STARTUP COMMAND 
#docker build --tag module-backend .
#docker run -d -p 9061:9061 --name module-ms module-backend
