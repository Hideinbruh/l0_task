FROM golang:latest

WORKDIR /go/src/app

# Копируем go.mod и go.sum и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы из текущей директории в контейнер
COPY . .

# Копируем статические файлы
COPY static/ /go/src/app/static/

# Компилируем приложение
RUN go build -o my_service cmd/main/main.go

# Запускаем приложение
CMD ["./my_service"]