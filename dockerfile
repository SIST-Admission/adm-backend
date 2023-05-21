# pull golang bullseye image
FROM golang:1.20.4-bullseye

# COPY application to WORKDIR
COPY . /app
WORKDIR /app

# build application
RUN go mod download && go mod verify
RUN go mod tidy
RUN go build -o main app/service.go

EXPOSE 8080

# run application
CMD ["./main"]