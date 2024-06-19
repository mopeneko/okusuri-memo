FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY back ./back/

RUN CGO_ENABLED=0 go build -o app ./back/cmd/back/...

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=0 /app/app /app

CMD ["/app"]
