# zolotoy-dev-backend

**Go Backend Engineering · Four-service architecture · Work in progress**

[English](#english) · [Русский](#русский)  
**Demo:** [zolotoy.dev](https://zolotoy.dev/) · **Frontend:** [zolotoy-dev-frontend](https://github.com/Officialsayp/zolotoy-dev-frontend)

---

## English

### Overview

**zolotoy-dev-backend** is the Go backend repository for [zolotoy.dev](https://zolotoy.dev/), an engineering portfolio covering four service domains: **Order**, **Auth**, **Notification**, and **URL Shortener**. Previously named *StockFlow*, this project develops demonstrable backend skills incrementally: HTTP, domain modeling, persistence, transactions, security, events, and reliability.

The frontend demo uses MSW mocks and runs **without Go services**. Implemented UI flows do **not** imply implemented or publicly deployed backend APIs.

### Implementation status

- **Order — in development:** an HTTP handler maps request DTOs to `domain.Order`, generates a UUID, calculates item totals, and returns `201 Created`. Orders are **not persisted**.
- **Auth — planned:** identity, sessions, refresh-token rotation, RBAC. No implementation.
- **Notification — planned:** durable jobs, delivery attempts, retries, recovery. No implementation.
- **URL Shortener — planned:** redirects, Redis caching, analytics. No implementation.

Current Order endpoints:

- `GET /health` — `204 No Content`.
- `POST /orders` — `201 Created` with a generated order; validation or service failures return error responses.
- `GET /orders/{id}?details=true` — a **learning-only** text response for a positive numeric ID, not a repository lookup. It cannot retrieve the UUID returned by POST.

The availability checker is an in-memory stub: it does not reserve inventory. There is no order repository, PostgreSQL, Redis, Kafka, Auth, or live frontend integration. The current `/orders` routes are not the proposed `/api/v1/orders` contract.

### Technology and architecture

Go **1.26.6** (as declared in the service's `go.mod`), standard-library `net/http`, and `github.com/google/uuid`. The repository uses one Go module and `internal` package boundary per implemented service. Bruno supplies manual HTTP request scenarios; `gofmt`, `go vet`, `go build`, and `go test` provide the current Go checks.

```text
services/order-service/
  cmd/order-service/       HTTP handlers and request/response DTOs
  internal/domain/        Orders, payments, money and state models
  internal/service/       Order creation and availability checks
  internal/availability/  In-memory product availability stub
  api/bruno/              Manual HTTP requests
  .ai/                    Learning notes and progress
scripts/check.sh          Formatting, vet, build and tests
docs/                     Architecture and roadmap
pending-review/           Experiments awaiting owner review
```

Only implemented services receive directories under `services/`. No service imports another service's private packages or accesses another service's database.

### Local development

**Requirements:** Go no older than the version in [Order Service go.mod](services/order-service/go.mod). No Docker or database is required at this stage.

```bash
git clone https://github.com/Officialsayp/zolotoy-dev-backend.git
cd zolotoy-dev-backend
bash scripts/check.sh
cd services/order-service
go run ./cmd/order-service
```

In another terminal:

```bash
curl -i http://localhost:8080/health

curl -i -X POST http://localhost:8080/orders \
  -H 'Content-Type: application/json' \
  -d '{"buyer_id":"buyer-123","payment_method":"prepaid","items":[{"product_id":"keyboard-001","name":"Keyboard","quantity":2,"unit_price":499000,"currency":"RUB"}],"delivery_address":"Krasnodar","buyer_comment":"Call before delivery"}'
```

Expect `204` for health and `201` for POST: a UUID, `status=created`, `payment_status=awaiting_payment`, `payment_method=prepaid`, `created_at`, and `items[0].total_price=998000`. The positive POST was manually verified with Bruno; this is **not** evidence of persistence or comprehensive API coverage.

For details, see the [Order Service README](services/order-service/README.md).

### Build and checks

```bash
bash scripts/check.sh
git diff --check
```

`check.sh` checks `gofmt` and then runs `go vet ./...`, `go build ./...`, and `go test ./...` in the Order module. Three unit tests currently cover available, unavailable, and technically failing product availability. DTO mapping, order totals, UUIDs, and domain transitions still require tests.

GitHub Actions runs build, lint, test, and an informational `gosec -no-fail` security scan. The latter is **not** a blocking security gate.

### Roadmap

Next: **`OrderRepository` interface → in-memory repository → POST Save → GET by UUID**, then PostgreSQL migrations and persistence tests. Later: lifecycle and payment states, cancellation/refunds, idempotency, concurrency, transactional outbox, Auth, Notification, and URL Shortener. Real frontend integration requires agreed API contracts and contract tests. Do not present planned architecture as running infrastructure.

### Documentation and related repositories

- [Order Service](services/order-service/README.md) · [Architecture](docs/architecture.md) · [Roadmap](docs/roadmap.md)
- [Restructuring decisions](docs/restructure.md) · [Learning state](services/order-service/.ai/SESSION_STATE.md) · [Pending review](pending-review/README.md)
- [Frontend](https://github.com/Officialsayp/zolotoy-dev-frontend) · [System overview](https://github.com/Officialsayp/zolotoy-dev) · [Infrastructure](https://github.com/Officialsayp/zolotoy-dev-infra)

**License:** [MIT](LICENSE).

---

## Русский

### О проекте

**zolotoy-dev-backend** — Go backend для [zolotoy.dev](https://zolotoy.dev/), инженерного портфолио с четырьмя направлениями: **Order**, **Auth**, **Notification** и **URL Shortener**. Ранее проект назывался *StockFlow*. Работа ведётся поэтапно: HTTP, доменное моделирование, хранение данных, транзакции, безопасность, события и надёжность.

Демонстрационный frontend использует моки MSW и работает **без запущенных Go-сервисов**. Реализованные сценарии интерфейса **не означают**, что соответствующие backend API созданы или опубликованы.

### Состояние реализации

- **Order — в разработке:** HTTP-обработчик преобразует request DTO в `domain.Order`, создаёт UUID, рассчитывает суммы позиций и возвращает `201 Created`. Заказы **не сохраняются**.
- **Auth — запланирован:** пользователи, сессии, ротация refresh-токенов, RBAC. Реализации нет.
- **Notification — запланирован:** задания, попытки доставки, повторы и восстановление. Реализации нет.
- **URL Shortener — запланирован:** редиректы, кеширование Redis и аналитика. Реализации нет.

Текущие маршруты Order Service:

- `GET /health` — `204 No Content`.
- `POST /orders` — `201 Created` с созданным заказом; ошибки валидации и сервиса возвращают соответствующие ответы.
- `GET /orders/{id}?details=true` — **учебный** текстовый ответ по положительному числовому ID, а не поиск заказа в хранилище. UUID из POST получить через него нельзя.

Проверка доступности товара — in-memory заглушка без резервирования остатков. Пока отсутствуют repository заказов, PostgreSQL, Redis, Kafka, Auth и live-интеграция с frontend. Текущий `/orders` не равнозначен целевому контракту `/api/v1/orders`.

### Технологии и архитектура

**Go 1.26.6** согласно `go.mod` сервиса, стандартная библиотека `net/http` и `github.com/google/uuid`. Для реализуемых сервисов предусмотрены отдельные Go-модули и границы пакетов `internal`. Bruno используется для ручных HTTP-запросов; текущие проверки выполняются через `gofmt`, `go vet`, `go build` и `go test`.

```text
services/order-service/
  cmd/order-service/       HTTP-обработчики и DTO
  internal/domain/        Заказы, оплата, деньги и состояния
  internal/service/       Создание заказов и проверка доступности
  internal/availability/  In-memory заглушка доступности товаров
  api/bruno/              Ручные HTTP-запросы
  .ai/                    Учебные материалы и прогресс
scripts/check.sh          Форматирование, vet, сборка и тесты
docs/                     Архитектура и план развития
pending-review/           Эксперименты на рассмотрении
```

В `services/` появляются только действительно реализуемые сервисы. Сервисы не импортируют чужие закрытые пакеты и не обращаются напрямую к чужой базе данных.

### Локальный запуск

**Требования:** Go не ниже версии из [go.mod Order Service](services/order-service/go.mod). На данном этапе Docker и база данных не требуются.

```bash
git clone https://github.com/Officialsayp/zolotoy-dev-backend.git
cd zolotoy-dev-backend
bash scripts/check.sh
cd services/order-service
go run ./cmd/order-service
```

Во втором терминале:

```bash
curl -i http://localhost:8080/health

curl -i -X POST http://localhost:8080/orders \
  -H 'Content-Type: application/json' \
  -d '{"buyer_id":"buyer-123","payment_method":"prepaid","items":[{"product_id":"keyboard-001","name":"Keyboard","quantity":2,"unit_price":499000,"currency":"RUB"}],"delivery_address":"Krasnodar","buyer_comment":"Call before delivery"}'
```

Ожидаются `204` для health и `201` для POST: UUID, `status=created`, `payment_status=awaiting_payment`, `payment_method=prepaid`, `created_at` и `items[0].total_price=998000`. Положительный POST проверен вручную в Bruno, но это **не подтверждает** сохранение заказов или полноценное тестирование API.

Подробности — в [README Order Service](services/order-service/README.md).

### Сборка и проверки

```bash
bash scripts/check.sh
git diff --check
```

`check.sh` проверяет `gofmt`, затем выполняет `go vet ./...`, `go build ./...` и `go test ./...` внутри модуля Order Service. Три unit-теста покрывают доступный товар, недоступный товар и техническую ошибку проверки доступности. DTO mapper, суммы, UUID и доменные переходы требуют дополнительного тестирования.

GitHub Actions запускает build, lint, test и информационный security scan `gosec -no-fail`. Последний **не блокирует** merge.

### План развития

Следующий этап: **интерфейс `OrderRepository` → in-memory repository → сохранение при POST → получение по UUID через GET**, затем PostgreSQL, миграции и тесты сохранения. Далее — состояния заказа и оплаты, отмена/refund, идемпотентность, конкурентные операции, transactional outbox, Auth, Notification и URL Shortener. Для live-интеграции с frontend потребуются согласованные API-контракты и контрактные тесты. Запланированную архитектуру нельзя выдавать за работающую.

### Документация и связанные репозитории

- [Order Service](services/order-service/README.md) · [Архитектура](docs/architecture.md) · [План развития](docs/roadmap.md)
- [Реструктуризация](docs/restructure.md) · [Учебная точка](services/order-service/.ai/SESSION_STATE.md) · [На рассмотрении](pending-review/README.md)
- [Frontend](https://github.com/Officialsayp/zolotoy-dev-frontend) · [Общий проект](https://github.com/Officialsayp/zolotoy-dev) · [Инфраструктура](https://github.com/Officialsayp/zolotoy-dev-infra)

**Лицензия:** [MIT](LICENSE).
