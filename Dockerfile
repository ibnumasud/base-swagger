FROM golang:1.10-alpine3.8 as builder
WORKDIR /go/src/github.com/ibnumasud/base-swagger
RUN apk --update add git
RUN rm -f /var/cache/apk/*

COPY . ./
RUN go get -u github.com/go-openapi/loads
RUN go get -u github.com/jessevdk/go-flags
RUN go get -u github.com/go-openapi/runtime
RUN go get -u github.com/go-openapi/runtime/flagext
RUN go get -u github.com/go-openapi/runtime/middleware
RUN go get -u github.com/go-openapi/runtime/security
RUN go get -u github.com/minio/minio-go

ADD ca-certificates.crt /etc/ssl/certs/
WORKDIR /go/src/github.com/ibnumasud/base-swagger/cmd/upload-s3-server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch 
COPY --from=builder /go/src/github.com/ibnumasud/base-swagger/cmd/upload-s3-server/main /
EXPOSE 9000

ENTRYPOINT ["./main","--port","9000","--host","0.0.0.0"]
 
