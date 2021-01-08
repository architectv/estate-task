# :hotel: Тестовое задание Avito.Недвижимость

[![Build Status](https://travis-ci.com/architectv/estate-task.svg?branch=main)](https://travis-ci.com/architectv/estate-task)
![Go Report](https://goreportcard.com/badge/github.com/architectv/estate-task)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/architectv/estate-task)
![Lines of code](https://img.shields.io/tokei/lines/github/architectv/estate-task)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/architectv/estate-task)

<!-- ToC start -->
# Содержание

1. [Запуск](#Запуск)
1. [Юнит-тесты](#Юнит-тесты)
1. [API](#API)
1. [Реализация](#Реализация)
<!-- ToC end -->

# Запуск

```
make build
make run
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
make migrate_up
```

Для миграций используется [golang-migrate/migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation).

# Юнит-тесты

```
make run_test
```

# API

> 1) Тело запроса (ответа) - в формате JSON.
> 2) В случае ошибки возвращается необходимый HTTP код, в теле содержится описание ошибки. Пример: ```{"error": "something went wrong"}```.

## POST /rooms/

Добавление номера отеля.

- Параметры тела запроса:
    - description - текстовое описание,
    - price - цена за ночь.
- Тело ответа:
    - room_id - идентификатор номера отеля.

**Пример**

Запрос:

```
curl -X POST localhost:9000/rooms/ \
-H "Content-Type: application/json" \
-d '{
	"description": "The Best Room",
	"price": 9000
}'
```

Ответ:

```
{
    "room_id": 144
}
```

## DELETE /rooms/:id

Удаление номера отеля.

- Параметры пути запроса:
    - id - идентификатор номера отеля.

**Пример**

Запрос:

```
curl -X DELETE localhost:9000/rooms/144
```

## GET /rooms/

Получение списка номеров отеля.

- Параметры строки запроса:
    - sort - поле, по которому производится сортировка:
        - id - по идентификатору (по дате добавления),
        - price - по цене.
- Тело ответа:
    - список номеров отеля.

> 1) Для сортировки по убыванию необходимо добавить знак минус перед значением поля (-id или -price).
> 2) По умолчанию (если параметр sort пуст или отсутствует) сортировка осуществляется по id по возрастанию.

**Пример**

Запрос:

```
curl -X GET localhost:9000/rooms/?sort=-price
```

Ответ:

```
[
    {
        "room_id": 2,
        "description": "description2",
        "price": 5000
    },
    {
        "room_id": 3,
        "description": "description3",
        "price": 3000
    },
    {
        "room_id": 1,
        "description": "description1",
        "price": 1000
    },
]
```

## POST /bookings/

Добавление бронирования номера отеля.

- Параметры тела запроса:
    - room_id - идентификатор номера отеля,
    - date_start - дата начала бронирования,
    - date_end - дата окончания бронирования.
- Тело ответа:
    - booking_id - идентификатор бронирования.

> Ограничения (из условия): нет проверки на доступность номера отеля в выбранное время.

**Пример**

Запрос:

```
curl -X POST localhost:9000/bookings/ \
-H "Content-Type: application/json" \
-d '{
	"room_id": 144,
	"date_start": "2021-12-30",
	"date_end": "2022-01-02"
}'
```

Ответ:

```
{
    "booking_id": 121
}
```

## DELETE /bookings/:id

Удаление бронирования номера отеля.

- Параметры запроса:
    - id - идентификатор бронирования.

**Пример**

Запрос:

```
curl -X DELETE localhost:9000/bookings/121
```
 
## GET /bookings/

Получение списка бронирований номера отеля

- Параметры строки запроса:
    - room_id - идентификатор номера отеля.
- Тело ответа:
    - список бронирований.

> Список сортируется по дате начала (date_start).

**Пример**

Запрос:

```
curl -X GET localhost:9000/bookings/?room_id=144
```

Ответ:

```
[
    {
        "booking_id": 289,
        "date_start": "2021-01-04",
	    "date_end": "2021-01-08"
    },
    {
        "booking_id": 121,
        "date_start": "2021-12-30",
	    "date_end": "2022-01-02"
    },
    {
        "booking_id": 256,
        "date_start": "2022-03-01",
	    "date_end": "2022-03-12"
    },
]
```

# Реализация

- Следование дизайну REST JSON API.
- Подход "Чистой Архитектуры" и техника внедрения зависимости.
- Работа с фреймворком [fiber](https://github.com/gofiber/fiber).
- Работа с БД Postgres с использованием библиотеки [sqlx](https://github.com/jmoiron/sqlx) и написанием SQL запросов.
- Конфигурация приложения - библиотека [viper](https://github.com/spf13/viper).
- Реализация Graceful Shutdown.
- Запуск из Docker.
- Юнит-тестирование с помощью моков - библиотеки [testify](https://github.com/stretchr/testify), [mock](https://github.com/golang/mock).
- Непрерывная интеграция, запуск тестов в Travis CI.

**Структура проекта**
```
.
├── pkg
│   ├── model       // основные структуры
│   ├── handler     // обработчики запросов
│   ├── service     // бизнес-логика
│   └── repository  // взаимодействие с БД
├── cmd             // точка входа в приложение
├── scripts         // SQL файлы с миграциями
└── configs         // файлы конфигурации
```
