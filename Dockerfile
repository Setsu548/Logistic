#build stage
FROM golang:1.20-alpine AS builder
RUN apk add --no-cache git
WORKDIR /logistic
COPY . .
RUN go get -d -v ./...
RUN go build -o logistic -v .


#final stage
FROM alpine:latest
LABEL Name=docucenter Version=0.0.1
RUN apk update
RUN apk --no-cache add ca-certificates
COPY --from=builder /logistic .
ENTRYPOINT [ "./logistic" ]
