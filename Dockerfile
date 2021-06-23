#builder
FROM golang:alpine as builder
WORKDIR /home
COPY . .
RUN go build -o boilerplate-api app/main.go

#final image
FROM alpine
RUN apk add tzdata
COPY --from=builder /home/boilerplate-api .
EXPOSE 6001
CMD ./boilerplate-api
