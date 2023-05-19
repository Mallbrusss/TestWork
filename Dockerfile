FROM golang:1.20.4
WORKDIR /web
COPY go.mod ./
COPY . .
RUN go build -o app
EXPOSE 8080
CMD ["./app"]
