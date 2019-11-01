# the alpine tag identifies the latest golang image on alpine
# using alpine mainly so that apk is guarenteed to exist in builder container
# CA certs must be downloaded by builder container and copied into final container
FROM golang:alpine AS builder
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY . .
# cross-compile statically linked 64bit Linux binary called main
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
