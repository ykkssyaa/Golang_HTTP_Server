# Тестовое задание Golang Developer

## Описание
Данный сервис получает по апи ФИО, из открытых апи обогащает ответ наиболее вероятными возрастом, 
полом и национальностью и сохраняет данные в БД. 
По запросу выдаёт информацию о найденных людях.

1. Реализованы REST методы:
   1. Добавление пользователя
   2. Удаление пользователя
   3. Обновление информации о пользователе
   4. Получение списка пользователей с фильтрами и пагинацией
2. В качестве базы данных используется Postgres. БД поднимается в контейнере с помощью docker-compose
3. Структура базы данных создается с помощью [миграций](https://github.com/golang-migrate/migrate)
4. Вся важная информация хранится в .env файле

**.env добавлен в .gitignore, поэтому в репозитории он не отображается.
.env файл располагается в директории /config и хранит:**

```
PORT=8080
POSTGRES_USER=user
POSTGRES_PASSWORD=userpass
POSTGRES_PORT=5432
POSTGRES_DBNAME=postgres
POSTGRES_DSN="host=localhost port=5432 user=user password=userpass dbname=postgres sslmode=disable"
```

5. Содержимое конфига читается с помощью [Viper](https://github.com/spf13/viper)
6. Пример использования API сервиса можно посмотреть в [Postman](https://www.postman.com/joint-operations-operator-99149269/workspace/golang-public/request/28284200-cf68b3b3-1999-4439-9634-982c55c7dc1c).
Тут настроенны 4 запроса с возможностью добавить JSON в тело запроса, настроить параметры.
7. В структуре проекта используется чистая архитектура. Весь функционал разделен на 3 слоя: 

**Server** - транспортный слой, который принимает и обрабатывает http запросы от клиента\
**Service** - слой бизнес логики\
**Gateway** - слой с внешними подключениями - Postgres и внешнее API

8. Реализовано gracefull shutdown - безопасный выход из программы с 
обработкой всех оставлшихся потоков и подключений и завершением их работы
9. Реализована фильтрация и пагинация списка пользователей.\
Пагинация работает через limit и offset.\
Фильтрация задается в параметрах запроса и каждое значение обозначает единственное допустимое значение заданного поля. 
Между этими параметрами можно поставить логический оператор AND.
К примеру:

При отправке запроса`/api/user?limit=3&name=Egor&age=20&country=RU&offset=5` 
Будет формироваться запрос 
```
SELECT * FROM users 
WHERE name="Egor" AND age=20 AND country="RU"
LIMIT=3 OFFSET 5
```

## Технологии
- Golang
- PostgreSQL
- Docker
- [Viper](https://github.com/spf13/viper)
- [Gorilla/mux](https://github.com/gorilla/mux)
- [SQLX](https://github.com/jmoiron/sqlx)

## Эндпойнты и методы
- ``/api/user`` (GET) - Получение списка сохраненных пользователей\
**Используемые параметры:**
1. limit
2. offset
3. name
4. surname
5. patronymic
6. country
7. age
8. gender

Если не задать limit или offset они получают значения limit=10 и offset = 0

- ``/api/user`` (POST) - Создание пользователя\
Задается JSON объект в теле запроса вида:
```json
{
  "name": "name",
  "surname": "surname",
  "patronymic": "patronymic"
}
```
name и surname - обязательные поля

- ``/api/user`` (DELETE) - Удаление пользователя\
Задается параметр id удаляемого пользователя
- ``/api/user`` (PATCH) - Обновление пользователя\
Задается JSON объект в теле запроса вида:
```json
{
  "id": "id",
  "name": "name",
  "surname": "surname",
  "patronymic": "patronymic",
  "country": "country",
  "age": 20,
  "gender": "male"
}
```
Поле id - обязательное. Также в объекте должно быть хотя бы одно ненулевое поле

## Примеры запросов и ответов
### Получение пользователей(GET)
Использование всех параметров:
http://localhost:8080/api/user?limit=3&offset=5&name=Egor&age=43&surname=Faleev&patronymic=Bbb&country=RU&gender=male

Использование только пагинации
http://localhost:8080/api/user?limit=10&offset=0

### Создание пользователя(POST)
http://localhost:8080/api/user
``` json
{
    "name": "Ivan",
    "surname": "Ivanov",
    "patronymic": "Ivanovich"
}
```


### Удаление пользователя(DELETE)

http://localhost:8080/api/user?id=1

### Обновление пользователя(PATCH)

http://localhost:8080/api/user

```
{
    "id":12,
    "age":100
}
```