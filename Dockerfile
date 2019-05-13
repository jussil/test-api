FROM golang:1.11 as builder
WORKDIR /go/src/github.com/jussil/test-api
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/test-api .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/test-api .
EXPOSE 8080
CMD ["./test-api"] 