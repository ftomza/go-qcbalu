FROM golang:1.15 AS build_base

WORKDIR /tmp/service

COPY . .

RUN go get -d -v ./...

RUN go test -v ./wallet/...

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o service ./wallet/service/app.go

FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /tmp/service/service /app/service

CMD ["/app/service"]