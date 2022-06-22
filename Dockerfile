FROM golang:1.18-alpine

WORKDIR /app

COPY . .

ENV APP_PORT=${APP_PORT}

RUN go get ./...
RUN go install ./...

RUN go build -o /binary

EXPOSE 2000

CMD ["/binary"]