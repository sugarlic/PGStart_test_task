FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./cmd/web 

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN apk add --no-cache postgresql-client


WORKDIR /root/
RUN mkdir /root/log
RUN touch /root/log/info.log
RUN touch /root/log/error.log

COPY --from=builder /app/myapp .

# Порт, на котором будет доступно приложение
EXPOSE 8080
