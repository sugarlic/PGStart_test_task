# Используем базовый образ Golang
FROM golang:latest

# Устанавливаем переменную окружения GOPATH
ENV GOPATH /go

# Копируем все файлы проекта в /go/src/app внутри образа
COPY . /go/src/app

# Устанавливаем рабочую директорию внутри образа
WORKDIR /go/src/app

# Собираем исполняемый файл
RUN go build -o app ./cmd/web 

# Открываем порт, на котором будет работать сервер
EXPOSE 8080

# Запускаем сервер при запуске контейнера
CMD ["./app"]
