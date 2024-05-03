# Используем официальный образ Go как базовый
FROM golang:1.21.0-alpine as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем исходники приложения в рабочую директорию
COPY . .
# Скачиваем все зависимости
RUN go mod tidy

# Собираем приложение
RUN go build -o main ./cmd

# Начинаем новую стадию сборки на основе минимального образа
FROM alpine:latest

# Добавляем исполняемый файл из первой стадии в корневую директорию контейнера
COPY --from=builder /app/main /main
COPY --from=builder /app/.env /.env

# Открываем порт
EXPOSE 30003

# Запускаем приложение
CMD ["/main"]