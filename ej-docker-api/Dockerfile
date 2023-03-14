FROM golang:latest

# set current work dir
WORKDIR /app

# copy all the content
COPY . .

# download dependencies
RUN go mod download

# build
RUN go build -o /ej-docker-api

# environments
ENV PORT=8081
ENV ENVIRONMENT=production

# port
EXPOSE 8081

# run
CMD ["/ej-docker-api"]