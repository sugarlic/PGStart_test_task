# Используем базовый образ Golang
FROM golang:1.21 as builder

# Устанавливаем переменную окружения GOPATH
ENV GOPATH /go

# Копируем все файлы проекта в /go/src/app внутри образа
COPY . .
COPY go.mod go.sum ./

# Устанавливаем рабочую директорию внутри образа
WORKDIR /go/src/app

RUN go mod download

# Собираем исполняемый файл
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .


# чистый образ для уменьшения размера
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируйте исполняемый файл из предыдущей стадии
COPY --from=builder /app/app .

# Порт, на котором будет доступно приложение
EXPOSE 8000

# Запуск приложения
CMD ["./app"]
