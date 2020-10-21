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