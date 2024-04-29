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
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./cmd/web 

# Вторая стадия сборки, используйте чистый образ для уменьшения размера
FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN apk add --no-cache postgresql-client


WORKDIR /root/
RUN mkdir /root/log
RUN touch /root/log/info.log
RUN touch /root/log/error.log

# Копируйте исполняемый файл из предыдущей стадии
COPY --from=builder /app/myapp .

# Порт, на котором будет доступно приложение
EXPOSE 8080

# # # # Запуск приложения
# CMD ["ls"]