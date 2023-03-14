FROM golang:1.17

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /ej-docker
EXPOSE 8081

CMD ["/ej-docker"]