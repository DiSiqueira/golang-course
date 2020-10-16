FROM golang:1.15

RUN curl -sL https://deb.nodesource.com/setup_14.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g nodemon

COPY . /menu-proxy-service
WORKDIR /menu-proxy-service

RUN go build ./cmd/menu-proxy-service/main.go

CMD nodemon --exec "make run-server"
