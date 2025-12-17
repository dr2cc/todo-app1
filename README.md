## API:

### POST /auth/sign-up

Creates new user 

##### Example Input: 
```
{
    "name":"Const",
    "username":"drk",
    "password":"qwerty"
}
```
##### Example Response:
```
{
    "id": 1
}
```

### POST /auth/sign-in

Request to get JWT Token based on user credentials
Запрос на получение токена JWT на основе учетных данных пользователя

##### Example Input: 
```
{
	"username":"drk",
    "password":"qwerty"
} 
```

##### Example Response: 
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjYwMzcwMDEsImlhdCI6MTc2NTk5MzgwMSwidXNlcl9pZCI6MX0.6SU7hcVREFNQABbGGzlp5TLrh3hjaQZnhpCGf4CgbPE"
} 
```

### POST /api/lists

Создает новый список, если передается Autorization header.
Autorization header создается из токена введенного во вкладке Authorization (Type: Bearer Token)

##### Example Input: 
```
{
    "title":"Список деталей"
}
```

##### Example Response:
```
{
    "id": 1
}
```

### GET /api/lists

Returns all lists
or
### GET /api/lists/{number of the list}

Returns list with this number

##### Example Response: 
```
{
    "data": [
        {
            "id": 7,
            "title": "Список деталей",
            "description": ""
        },
        {
            "id": 8,
            "title": "Список покупок",
            "description": "очень важно!"
        }
    ]
}
```
or
```
{
    "id": 7,
    "title": "Список деталей",
    "description": ""
}
```

- Дальше из другого примера. Для образца 
### DELETE /api/bookmarks

Deletes bookmark by ID:

##### Example Input: 
```
{
	"id": "5da2d8aae9b63715ddfae856"
} 
```


Комментарий автора:
- Про первый комментарий не знаю, а make migrate при первом запуске- обязательно!
# Для запуска приложения:
make build && make run
# Если приложение запускается впервые, необходимо применить миграции к базе данных:
make migrate


# Run project

Use ```make run``` to build and run docker containers with application itself and mongodb instance

# Применение схемы миграции (миграции соответствуют системе контроля версий для db)
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

## Stop the project and destroy the containers. Также необходимо для очистки кэша docker-compose и сборки проекта с новой сигнатурой

`docker-compose down`

## Rebuild the project

`make build`
`make run`

## (при запуске в Linux) wait-for-postgres.sh должен иметь LF  (Line Feed, перевод строки) последовательность конца строки. 
Последовательность конца строки в стиле Windows (CRLF- Carriage Return + Line Feed, возврат каретки и перевод строки), вызывает проблемы в Linux-контейнерах (которые ожидают LF). Интерпретатор /bin/sh видит символ возврата каретки как часть имени команды, и в результате файл не может быть выполнен.
- Решение: Преобразовать файл в формат LF.
- В редакторе, например VS Code, и в правом нижнем углу выбрать формат окончания строк LF.
Git не сохраняет такое изменение, надо проводить во всех экземплярах (поможет .gitattributes ?)

## Если запускаем "вручную": 
### Запуск контейнера (временного) с db (образ postgres должен уже быть скачан) с внешним портом 5436.
docker run --name=todo-db -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d --rm postgres
### Запуск проекта (в config.yml активировать настройки "Использовать при запуске приложения из cmd")
	go run .\cmd\main.go
### Применяем схему миграции из 2.


# Настройка Swagger для проекта на Golang
https://youtu.be/DBZgt9iIWzk?t=641
9:30

[Документация swagger в нашем проекте](http://localhost:8000/swagger/index.html)

# Установка swagger
go install github.com/swaggo/swag/cmd/swag@latest

# После запускаем из корня, но ссылаясь на main.go
# Так создается документация.
swag init -g .\cmd\main.go

# Была ошибка в import файла handler.go . Два варианта решения:
# Написать:
import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

# или в начале закомментировать хендлер 
# // router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
# Написать :
import (
	"github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
)
# а после снять комментарий с хендлера и сохранить.


# 5 
## Импорт (godotenv). Позволяет читать .env (зачем упомянут здесь, не знаю)
go get -u github.com/joho/godotenv

## Создать в корне файл .env с данными:
DB_PASSWORD=qwerty

# 4
## Если нужен откат таблиц db с текущей версии на предыдущую
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down


## Состояние docker, из него берем CONTAINER ID, в этот раз 6f02c8402d0b
docker ps
## Подключение к db в контейнере 6f02c8402d0b
docker exec -it 6f02c8402d0b /bin/bash

## Комманды linux
psql -U postgres

select * from users;
\d

exit
exit

# =====-*-=======
## (скачивание образа postgres в docker)
docker pull postgres

# About  the migration utilite (https://github.com/golang-migrate/migrate/) on Debian
wget http://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.deb
sudo dpkg -i migrate.linux-amd64.deb
rm migrate.linux-amd64.deb
(последнее- удаление)

# About  the migration utilite on windows.

<a href="https://github.com/golang-migrate/migrate/tree/master/cmd/migrate">migrate</a>

# Install scoop:

Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression

scoop install migrate

# Делается если в репозитории нет схем миграции. Создание структур миграции, утилита migrate к этому моменту должна быть установлена
migrate create -ext sql -dir ./schema -seq init


# REST API Для Создания TODO Списков на Go

## <a href="https://www.youtube.com/playlist?list=PLbTTxxr-hMmyFAvyn7DeOgNRN8BQdjFm8">Видеокурс на YouTube</a>

## В курсе разобранны следующие концепции:
- Разработка Веб-Приложений на Go, следуя дизайну REST API.
- Работа с фреймворком <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.
- Подход Чистой Архитектуры в построении структуры приложения. Техника внедрения зависимости.
- Работа с БД Postgres. Запуск из Docker. Генерация файлов миграций. 
- Конфигурация приложения с помощью библиотеки <a href="https://github.com/spf13/viper">spf13/viper</a>. Работа с переменными окружения.
- Работа с БД, используя библиотеку <a href="https://github.com/jmoiron/sqlx">sqlx</a>.
- Регистрация и аутентификация. Работа с JWT. Middleware.
- Написание SQL запросов.
- Graceful Shutdown