FROM golang:1.22.1-bookworm as builder
WORKDIR /go/src/app

COPY . .

RUN go build -v -o api ./cmd/api/main.go

FROM golang:1.22.1-bookworm as dev

WORKDIR /bin

COPY --from=builder /go/src/app/api .
CMD ["./api"]

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /go/src/app/

CMD ["air"]
