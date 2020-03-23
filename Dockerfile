FROM golang:alpine as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o bin/app

FROM scratch as runtime
COPY --from=builder /app/bin ./
ENTRYPOINT ["/app"]