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


```shell
sudo apt  install protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latests

protoc --go_out=./pkg/queues/basketconfirmedpb ./api/proto/basket_confirmed.proto
protoc --go_out=./pkg/queues/orderstatuschangedpb ./api/proto/order_status_changed.proto
```