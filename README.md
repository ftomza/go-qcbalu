# go-qcbalu
The microservice: "The balance of the user"
### Стек технологий.
- Golang 1.15
- [ent.](https://entgo.io/) - Entity framework, обработка и хранение  кошелька
- [ampq](github.com/streadway/amqp) - Клиент для RabbitMQ
- [zap](https://github.com/uber-go/zap) - Логирование
- mockery - Mock для тестов
- testify - Тулкит для тестов
### Окружение.
- docker - среда выполнения инстансов сервиса и вспомогательных сервисов
- rabbitmq - брокер сообщений
- postgresql - хранение данных
### Сборка и доставка.
- github actions
### Доступ клиентов.
Доступ осуществляется на стандартный порт AMPQ RabbitMQ с аутентификацией и авторизацией RabbitMQ.	

### Дополнительно
 - [Спецификация](https://github.com/ftomza/go-qcbalu/blob/main/spec.md) обмена данными
 - Описание [модели](https://github.com/ftomza/go-qcbalu/blob/main/domain/model.md)

### Использование

**Сборка**:

- Сервиса: `go build -o service ./wallet/service/app.go`
- Клиента: `go build -o client ./cmd/client/cmd.go`
- Слушателя: `go build -o sub ./cmd/sub/cmd.go`
- Конфигуратор MBS: `go build -o config-mbs ./cmd/config_mbs/cmd.go`


**Конфиг файл сервиса и переменные окружения**:

`qcb_wallet_config.yaml` Дефолтные значения:
```yaml
node_name: 0                                            #ENV: QCB_WALLET_NODE_NAME
ent_url: postgresql://root:root@127.0.0.1/qcb_wallet    #ENV: QCB_WALLET_ENT_URL
mbs_url: amqp://localhost/qcbalu                        #ENV: QCB_WALLET_MBS_URL
rpc_queue: rpc.wallet.main                              #ENV: QCB_WALLET_RPC_QUEUE
pub_exchange: pub.wallet                                #ENV: QCB_WALLET_PUB_EXCHANGE
mbs_timeout: 10000                                      #ENV: QCB_WALLET_MBS_TIMEOUT
```

**Подготова**

- Создать БД: `CREATE DATABASE qcb_wallet`
- Создать VHost RabbitMQ: `qcbalu`
- Сконфигурировать VHost RabbitMQ: `./config-mbs`, если адрес MBS отличается от дефолтного, то нужно указать параметр: `--uri=[Новый адрес]`

**Запуск сервиса**:

`QCB_WALLET_NODE_NAME=0 ./service`

Для следующих нод запуск аналогичен, только меняем `QCB_WALLET_NODE_NAME`

**Запуск слушателя**:

`./sub` 

Если адрес MBS отличается от дефолтного, то нужно указать параметр: `--uri=[Новый адрес]`

**Запуск интерактивного клиента**:

`./client`

Если адрес MBS отличается от дефолтного, то нужно указать параметр: `--uri=[Новый адрес]`

**Запуск клиента и сценарий выполнения команд**:

`./client < ./cmd/client/commands.txt`