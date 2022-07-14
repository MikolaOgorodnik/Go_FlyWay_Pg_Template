# Build stage
FROM golang:1.18-buster AS go-env

ADD . /src
WORKDIR /src

RUN go build -o /app

# Run stage
FROM golang:1.18-buster

EXPOSE ${APP_PORT}

WORKDIR /
COPY --from=go-env /app /

CMD ["/app"]