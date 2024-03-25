# mini-shop

Небольшое api для каталога товаров

### Stack
  * Golang 1.22
  * Postgres
  * Docker

## Возможные действия:
  * Получение списка всех категорий
  * Получение списка товаров в конкретной категории
  * Авторизация пользователей
  * Добавление/Редактирование/Удаление категории (для авторизованных пользователей)
  * Добавление/Редактирование/Удаление товара (для авторизованных пользователей)

## Запуск проекта
  Необходимо склонировать репозиторий:
  ```
     git clone https://github.com/sh1neqd/mini-shop
     git checkout main
```
  Далее запускаем проект с помощью docker-compose:
```
  docker-compose build
  docker-compose up
```
  Запускаем Postman и теституем)

## Auth
  Регистрация[post]
* http://localhost:5000/api/signup
```
{
    "username": username,
    "password": password,
    "email": email
}
```
  Авторизация[post]
  
* http://localhost:5000/api/login
```
{
    "username": username,
    "password": password
}
```
json answer:
```
{
    "message": access_token
}
```
## Category
* http://localhost:5000/api/category/ [get] getting all categories
* http://localhost:5000/api/category/ [post] creating category:
```
{
    "name": category_name
}
```
* http://localhost:5000/api/category/:id [get] getting category by id(int)
* http://localhost:5000/api/category/:id [patch] update category by id(int)
```
{
    "name": category_name
}
```
* http://localhost:5000/api/category/:id [delete] delete category by id(int)
## Item
* http://localhost:5000/api/item/ [get] getting all items
* http://localhost:5000/api/item/ [post] create item:
```
{
    "name": item_name,
    "price": price
}
```
* http://localhost:5000/api/item/:id [get] get item by id(int)
* http://localhost:5000/api/item/:id [patch] update item by id(int):
```
{
    "name": item_name,
    "price": price
}
```
* http://localhost:5000/api/item/:id [delete] delete item by id(int)
* http://localhost:5000/api/item/:id [post] add category to item:
```
{
    "category_id": category_id
}
```
