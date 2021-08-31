FROM golang:latest

COPY . /usr/src/app

WORKDIR /usr/src/app

CMD ["/usr/local/go/bin/go", "run", "."]
