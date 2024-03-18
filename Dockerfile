FROM golang:1.22.0

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client
RUN chmod +x wait.sh
RUN go mod download
RUN go build -o todo-app ./cmd/REST-API/main.go

CMD ["./restapi"]