FROM golang:1.22.0

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait.sh

# build go app
RUN go mod download
RUN go build -o todo-app ./cmd/REST-API/main.go

CMD ["./restapi"]