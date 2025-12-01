FROM golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем все локальные файлы в /app внутри контейнера
COPY ./ ./

# install psql и другие зависимости
RUN apt-get update && apt-get install -y postgresql-client && rm -rf /var/lib/apt/lists/*

# Делаем wait-for-postgres.sh исполняемым
RUN chmod +x wait-for-postgres.sh

# Собираем приложение
RUN go mod download
RUN go build -o todo-app ./cmd/main.go

# Это из оригинального (todo-app) проекта, мой- todo-app1 Но это не влияет на сборку
CMD ["./todo-app"]