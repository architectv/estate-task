FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN sed -i -e 's/\r$//' *.sh
RUN chmod +x *.sh

# build go app
RUN go mod download -x
RUN go build -o app ./cmd/main.go

CMD ["./app"]