# ROADMAP.md

Этапы обучения.

- [x] net/http
- [x] GET handlers
- [x] POST handlers
- [x] JSON
- [x] Базовые request/response DTO и JSON-сериализация
- [x] Transport validation и отличие от business validation
- [ ] Service Layer
  - [x] Контракт ошибок и безопасное сопоставление с HTTP-ответами
  - [x] `ProductAvailability` interface рядом с потребителем
  - [x] Constructor injection через `NewOrderService`
  - [x] Memory `AvailabilityChecker` и wiring в `main`
  - [x] Ручная проверка через Bruno
  - [x] Unit tests `OrderService`: available, unavailable, technical error
  - [ ] DTO vs Domain — следующая тема
  - [ ] Настоящий `CreateOrder` с созданием `domain.Order`
- [ ] Repository
  - [ ] Интерфейс `OrderRepository`
  - [ ] Memory-хранилище: сохранение и получение заказа
- [ ] PostgreSQL — после работающего сценария с memory repository
- [ ] Middleware
- [ ] Context
- [ ] Logging
- [ ] Swagger
- [ ] Testing — дальнейшие domain, handler, repository и integration tests;
  базовые unit tests `OrderService` уже пройдены
- [ ] Docker
- [ ] Redis
- [ ] Kafka
- [ ] gRPC

Переходить к следующему этапу только после успешного завершения предыдущего.
Тесты добавлять вместе с проверяемым поведением, не откладывая их до отдельного этапа Testing.
