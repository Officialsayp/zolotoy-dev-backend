# Order Service

Учебная реализация первого backend-сервиса zolotoy.dev. Текущий этап — Service Layer.

## Сейчас работает

- `GET /health` → 204 без body.
- `GET /orders/{id}?details=true` → учебный текстовый ответ; проверка положительного ID.
- `POST /orders` с JSON `{"product":"keyboard"}` → 201 и нормализованное имя товара.
- Некорректный JSON, пустой product и `unavailable` → 400; внутренняя ошибка → безопасный 500.

Заказы **не сохраняются**. 201 демонстрирует HTTP/service flow и не означает запись в БД.
GET не читает заказ из хранилища. Доменная модель пока не подключена к handler.

## Запуск

Из этой директории, с Go не ниже версии из go.mod:

```bash
go run ./cmd/order-service
```

Сервер слушает `:8080`; это локальный учебный запуск, не production-конфигурация.
Bruno: открыть `api/bruno/order-service-local`, задать `baseUrl=http://localhost:8080`.

## Структура

- `cmd/order-service/main.go` — сборка приложения и текущие HTTP handlers.
- `internal/service` — существующая бизнес-проверка и контракт ошибки.
- `internal/domain` — сохранённые модели и переходы состояний, требуют тестов и интеграции.
- `.ai` — учебные правила, журнал и точка продолжения.

Разделять HTTP handlers и main стоит вместе с следующим осмысленным шагом транспорта.
Пустые repository-пакеты удалены; настоящий repository появится вместе с реализацией.

## Цель

Жизненный цикл заказа, независимые состояния оплаты и заказа, PostgreSQL,
идемпотентность, optimistic concurrency, история и transactional outbox.
Сначала согласовать API по [архитектуре](../../docs/architecture.md).
Текущий `/orders` не совместим с целевым `/api/v1/orders` и DTO frontend.
Inventory и API Gateway не являются обязательными сервисами текущего плана.
