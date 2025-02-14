# Используем официальный образ Go
FROM golang:1.22

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем файлы в контейнер
COPY . .

# Скачиваем зависимости
RUN go mod tidy

# Собираем бинарник
RUN go build -o notification-service .

# Запускаем сервис
CMD [ "./notification-service" ]
