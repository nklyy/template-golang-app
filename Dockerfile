FROM golang:1.19-alpine as base_build

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY .. .

RUN CGO_ENABLED=0 go build -o compiled ./cmd/main.go

CMD ["/app/cmd/main"]

FROM scratch
COPY --from=base_build ./app .
CMD ["./compiled"]