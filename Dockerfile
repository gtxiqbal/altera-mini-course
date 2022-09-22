# step 1: build executable binary
FROM golang:1.18-alpine3.16 as builder
LABEL maintainer="Mohammad Iqbal Alhusain<moh.iqbal.alhusain@gmail.com>"
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /mini-course-service

# step 2: build a small image
FROM alpine:3.16.2
WORKDIR /app
COPY --from=builder mini-course-service .
COPY .env .
EXPOSE 8080
CMD ["./mini-course-service"]