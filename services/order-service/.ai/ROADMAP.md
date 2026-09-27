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
  - [x] DTO vs Domain
  - [x] Request DTO `createOrderV1Request` и mapper → `CreateOrderInput`
  - [x] `ProductAvailability.IsAvailable(productID, quantity)`
  - [x] Domain mapping `CreateOrderInput` → `Money` → `OrderItem`
  - [x] `OrderItem`: `ProductID`, `ProductNameSnapshot`, `Quantity`, `UnitPrice`, `TotalPrice`
  - [x] `PaymentMethod`: `prepaid` / `pay_on_receipt_online` и `ParsePaymentMethod`
  - [x] Серверный UUID; `CreateOrder` создаёт и возвращает `domain.Order`
  - [x] Getters `Order` и mapper → `createOrderResponse`
  - [x] Bruno: `201`, UUID, начальные статусы, `created_at`, `items[0].total_price=998000`
- [ ] Repository — следующий этап
  - [ ] Интерфейс `OrderRepository` рядом с потребителем
  - [ ] Memory repository и wiring в `OrderService`
  - [ ] `POST /orders` сохраняет заказ через `Save`
  - [ ] `GET /orders/{id}` возвращает сохранённый заказ по UUID
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

Текущая граница: создание заказа работает, хранения ещё нет; GET остаётся учебным.
Три service unit tests проверяют исходы ошибок, но не поля созданного заказа.
Service Layer и Testing целиком не завершены; актуальная точка — [SESSION_STATE](SESSION_STATE.md).

Переходить к следующему этапу только после успешного завершения предыдущего.
Тесты добавлять вместе с проверяемым поведением, не откладывая их до отдельного этапа Testing.
