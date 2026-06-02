FROM golang:1.22-alpine AS builder
WORKDIR /app
# Cache dependency layer separately
COPY go.mod go.sum ./
RUN GOPROXY=https://goproxy.io,https://goproxy.cn,direct go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main .
RUN mkdir -p uploads
EXPOSE 8080
CMD ["./main"]
