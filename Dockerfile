FROM golang:1.22.5-alpine AS build

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN go build -o main ./main.go

FROM alpine:latest

RUN apk --no-cache add \
    bash \
    ca-certificates \
    && mkdir /assets \
    && mkdir /templates \
    && mkdir /static

WORKDIR /app

COPY --from=build /app/main .

COPY ./assets ./assets
COPY ./templates ./templates
COPY ./static ./static

EXPOSE 8080

CMD ["./main"]
