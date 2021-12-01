FROM golang:latest
WORKDIR /app
RUN go mod init github.com/mattssll/go_paradise/go_mysql_docker
RUN go mod download
COPY . .
RUN go build -o gomysql .
CMD ["./gomysql"]