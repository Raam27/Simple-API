FROM golang:1.16-alpine

WORKDIR /echo-gorm
COPY . ./
RUN go mod download

EXPOSE 1323

CMD [ "go","run","server.go" ]