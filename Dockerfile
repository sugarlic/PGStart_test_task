# Используйте официальный образ Golang как базовый
FROM golang:1.21 as builder

# Установите рабочую директорию в контейнере
WORKDIR /app

# Копируйте go.mod и go.sum для управления зависимостями
COPY go.mod go.sum ./

# Загружайте зависимости
RUN go mod download

# Копируйте исходный код проекта
COPY . .

# Соберите исполняемый файл вашего приложения
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Вторая стадия сборки, используйте чистый образ для уменьшения размера
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируйте исполняемый файл из предыдущей стадии
COPY --from=builder /app/myapp .

# Порт, на котором будет доступно приложение
EXPOSE 8000

# Запуск приложения
CMD ["./myapp"]