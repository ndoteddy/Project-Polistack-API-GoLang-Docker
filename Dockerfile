FROM golang:1.8-alpine
ADD . /go/src/simple-api-restaurant
RUN go install simple-api-restaurant

FROM alpine:latest
COPY --from=0 /go/bin/simple-api-restaurant .
ENV PORT 8080
CMD ["./simple-api-restaurant"]
