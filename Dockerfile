FROM golang:1.21.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app .

FROM alpine:3.19.1 

RUN apk --no-cache add ca-certificates logrotate

WORKDIR /app

RUN addgroup -S app && adduser -S app -G app && chown -R app:app /app

USER app

COPY --from=builder /app/app .

CMD ["./app"] 
