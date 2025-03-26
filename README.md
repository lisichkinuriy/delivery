# Сервис Delivery

## Use Cases

### Менеджер
- получить незавершенные заказы
- получить курьеров

### Курьер
- Переместить курьеров
- Назначить заказ на курьера

### Basket
- создать заказ


## Запуск через докер

```shell
docker compose up
```

## Билд

```shell
make build
```

## Mockery


```shell
docker pull vektra/mockery
```

```shell
docker run -v "$PWD":/src -w /src vektra/mockery --all --case=underscore
```

