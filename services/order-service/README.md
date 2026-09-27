# Order Service

Учебная реализация первого backend-сервиса zolotoy.dev. В Service Layer пройдено
создание `domain.Order` из request DTO; следующий этап — `OrderRepository`.

## Сейчас работает

- `GET /health` → 204 без body.
- `GET /orders/{id}?details=true` → учебный текстовый ответ по положительному числовому ID.
- `POST /orders` → request DTO → `CreateOrderInput` → `Money` / `OrderItem` →
  `domain.Order` → response DTO с `201 Created`.
- UUID генерируется в service; начальные статусы и время создания задаёт `domain.NewOrder`.
- Способ оплаты проверяется через `ParsePaymentMethod`: `prepaid` / `pay_on_receipt_online`.
- Некорректный JSON, неверные обязательные поля, `ErrInvalidOrder` и недоступная
  позиция с `product_id=unavailable` → 400; техническая ошибка → безопасный 500.

Заказы **не сохраняются**: после запроса нет хранилища, из которого можно получить заказ.
`201` подтверждает создание доменного объекта и формирование ответа, а не запись в БД.
GET не читает repository и пока вернёт `400` для UUID из POST.

## Текущий POST и результат занятия

```json
{
  "buyer_id": "buyer-123",
  "payment_method": "prepaid",
  "items": [
    {
      "product_id": "keyboard-001",
      "name": "Keyboard",
      "quantity": 2,
      "unit_price": 499000,
      "currency": "RUB"
    }
  ],
  "delivery_address": "Krasnodar",
  "buyer_comment": "Call before delivery"
}
```

Ручная проверка пользователя в Bruno успешна: `201 Created`, серверный `id` (UUID),
`status=created`, `payment_status=awaiting_payment`, `payment_method=prepaid`,
`items[0].total_price=998000` (`499000 × 2`) и `created_at`. Покупатель, адрес и комментарий
переносятся в response. `total_price` — сумма позиции внутри `items`; имя и цена пока
приходят из request, сервер вычисляет произведение цены и количества.

`ProductAvailability.IsAvailable(productID, quantity)` вызывается для каждой позиции,
но memory checker пока игнорирует количество и не резервирует остатки.

## Запуск

Из этой директории, с Go не ниже версии из go.mod:

```bash
go run ./cmd/order-service
```

Сервер слушает `:8080`; это локальный учебный запуск, не production-конфигурация.
Bruno: открыть `api/bruno/order-service-local`, задать `baseUrl=http://localhost:8080`.
Актуальный положительный запрос — [Post-request.yml](api/bruno/order-service-local/Post-request.yml),
пока без автоматических assertions. Старые запросы с `{product}` относятся к предыдущему
уроку и требуют актуализации; их результаты не подтверждают новый DTO.

## Структура

- `cmd/order-service/main.go` — сборка приложения и текущие HTTP handlers.
- `cmd/order-service/create_order_dto.go` — request/response DTO и их mapper.
- `internal/service` — вход создания заказа, domain mapping, availability, UUID и контракт ошибок.
- `internal/domain` — создание заказа, деньги, позиции и переходы состояний; требуют дальнейших тестов.
- `internal/availability/memory` — учебный availability checker, не repository заказов.
- `.ai` — учебные правила, журнал и точка продолжения.

Разделять HTTP handlers и main стоит вместе с следующим осмысленным шагом транспорта.
Пустые repository-пакеты удалены; настоящий repository появится вместе с реализацией.

## Проверки и продолжение обучения

Из корня репозитория: `bash scripts/check.sh` и `git diff --check`.
Скрипт проверяет форматирование, затем выполняет `go vet`, `go build` и `go test`.
Три service unit tests проверяют available / unavailable / technical error, но не поля
созданного заказа; проверки DTO mapper, UUID, сумм и `ErrInvalidOrder` ещё не добавлены.

Следующая тема: `OrderRepository` interface → memory repository → POST Save → GET by ID.
Критерий: GET по UUID из POST возвращает тот же сохранённый заказ в рамках процесса.
Затем — PostgreSQL и чтение после перезапуска. Учебный цикл и точная точка продолжения —
в [SESSION_STATE](.ai/SESSION_STATE.md) и [ROADMAP](.ai/ROADMAP.md).

## Цель

Жизненный цикл заказа, независимые состояния оплаты и заказа, PostgreSQL,
идемпотентность, optimistic concurrency, история и transactional outbox.
Сначала согласовать API по [архитектуре](../../docs/architecture.md).
Текущий `/orders` не совместим с целевым `/api/v1/orders` и DTO frontend.
Inventory и API Gateway не являются обязательными сервисами текущего плана.
