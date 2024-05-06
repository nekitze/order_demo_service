FROM golang:1.22.2-bullseye as build

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  app-user

WORKDIR /go/src/orders_service

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/server

FROM gcr.io/distroless/static-debian11 as app

COPY --from=build /go/src/orders_service/ .

#USER app-user:app-user

CMD ["./app"]