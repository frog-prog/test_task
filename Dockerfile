FROM golang:1.20-alpine AS builder
WORKDIR /usr/local/src
RUN apk --no-cache add bash gcc gettext musl-dev
COPY ["go.mod","go.sum","./"]
RUN go mod download
COPY . ./
RUN go build  main.go
FROM alpine
COPY --from=builder /usr/local/src /
COPY .env /.env
EXPOSE 3000
CMD ["./main"]