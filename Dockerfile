FROM golang:1.22

WORKDIR /app

COPY / ./

ENV TODO_PORT="7540"
ENV TODO_DBFILE="/db/scheduler.db"
ENV TODO_PASSWORD="password"
ENV SECRET_KEY="my_secret_key"

EXPOSE 7540

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /my_app ./cmd
CMD ["/my_app"]