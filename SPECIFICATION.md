# API

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
- CRUD для списков

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

### DELETE /api/lists/{number of the list}
```
Deletes list with this number
```
### PUT /api/lists/{number of the list}
Обновление данных в списке с данным номером
##### Example Input: 
```
{
    "title": "Список деталей и расходников",
    "description": "маленький"
}
```
- CRUD для записей в списках

### POST /api/lists/10/items

Создает новую запись, списске с данным номером

##### Example Input: 
```
{
    "title":"Хомут"
}
```

##### Example Response:
```
{
    "id": 1
}
```

### GET /api/lists/10/items

Returns all items from list 10

##### Example Response: 
```
[
    {"id":1,"title":"Хомут","description":"","done":false},
    {"id":2,"title":"Хомут","description":"","done":false},
    {"id":3,"title":"Гайка","description":"","done":false}
]
```
### GET /api/items/4

Возвращает запись №4 (в проекте записи имеют сквозную нумерацию, независимую от списков)

##### Example Response: 
```
{
    "id": 4,
    "title": "Reading",
    "description": "",
    "done": false
}
```

### DELETE /api/items/4
```
Delete items with this number
```
### PUT /api/items/3
Обновление данных в записи с данным номером
##### Example Input: 
```
{
    "title": "Цапа",
    "description": "ржавая",
    "done":true
}
```