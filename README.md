# zolotoy-dev-backend

Go backend для [zolotoy.dev](https://zolotoy.dev) — технической консоли четырёх
портфолио-сервисов: **Order, Auth, Notification и URL Shortener**.
Репозиторий ранее назывался StockFlow. Цель — последовательно построить и уметь
объяснить проверяемые backend-сценарии: от HTTP и бизнес-правил до транзакций,
безопасности, событий и производительности.

Frontend живёт отдельно в [zolotoy-dev-frontend](https://github.com/Officialsayp/zolotoy-dev-frontend).
Его mock-демонстрация работает без Go backend. Наличие экранов и моков не означает,
что соответствующие backend API уже реализованы или подключены к сайту.

## Состояние

| Сервис | Назначение | Реализация в этом репозитории |
| --- | --- | --- |
| Order | Заказы, оплата, переходы состояний, идемпотентность и outbox | POST создаёт и возвращает domain.Order; GET учебный; без хранилища |
| Auth | Пользователи, сессии, refresh rotation и RBAC | План; кода пока нет |
| Notification | Событие → задание → попытки доставки, retries и dead jobs | План; кода пока нет |
| URL Shortener | Короткие ссылки, redirect, Redis cache и аналитика | План; кода пока нет |

Это работа в процессе, не готовая production-платформа. В Service Layer пройдены
DTO vs Domain и создание заказа; следующий этап — `OrderRepository` и memory repository.
PostgreSQL, Kafka, Redis, авторизация и live-интеграция с frontend ещё не подключены.
Kubernetes и API Gateway не нужны для начала разработки.

## Быстрый старт

Нужен Go версии не ниже указанной в [go.mod](services/order-service/go.mod).
HTTP использует стандартную библиотеку, UUID — `github.com/google/uuid`;
Docker и БД для текущего этапа не требуются.

```bash
git clone https://github.com/Officialsayp/zolotoy-dev-backend.git
cd zolotoy-dev-backend
bash scripts/check.sh
cd services/order-service
go run ./cmd/order-service
```

В другом терминале:

```bash
curl -i http://localhost:8080/health
curl -i -X POST http://localhost:8080/orders \
  -H 'Content-Type: application/json' \
  -d '{"buyer_id":"buyer-123","payment_method":"prepaid","items":[{"product_id":"keyboard-001","name":"Keyboard","quantity":2,"unit_price":499000,"currency":"RUB"}],"delivery_address":"Krasnodar","buyer_comment":"Call before delivery"}'
```

Ожидаются 204 для health и 201 с созданным заказом для POST: серверный UUID,
`status=created`, `payment_status=awaiting_payment`, `payment_method=prepaid`,
`items[0].total_price=998000` и `created_at`. Этот POST вручную проверен пользователем в Bruno.
POST пока **не сохраняет заказ**: после запроса его нельзя получить повторно.
GET остаётся учебным ответом по числовому ID и не читает repository.
Подробности и Bruno — в [README Order Service](services/order-service/README.md).

## Структура

```text
services/order-service/   единственный реализованный Go-модуль
  cmd/order-service/      запуск сервера и текущие handlers
  internal/domain/        модели заказа, денег и оплаты
  internal/service/       создание заказа и бизнес-проверки
  api/bruno/              ручная коллекция; текущий POST в Post-request.yml
  .ai/                    обучение и журнал
scripts/check.sh          форматирование, vet, build, test
docs/                     архитектура и план развития
pending-review/           сохранённые эксперименты на решение владельца
```

Новые service-директории появятся вместе с кодом. Сервисы сохраняют собственные
модули и границы данных; frontend остаётся отдельным репозиторием.

## Документация и продолжение

- [Архитектура и источники планов](docs/architecture.md)
- [Порядок развития и критерии готовности](docs/roadmap.md)
- [Разбор структуры: что сохранено, убрано и отложено](docs/restructure.md)
- [Материалы на разбор владельцу](pending-review/README.md)
- [Текущий учебный контекст](services/order-service/.ai/SESSION_STATE.md)

Ближайшая последовательность: `OrderRepository` interface → memory repository →
POST Save → GET by ID → затем PostgreSQL. Сначала проверить сохранение и получение
одного заказа по UUID в памяти, затем постоянное хранение и согласованный API. Не добавлять инфраструктуру ради списка
технологий и не выдавать целевые требования за реализованные возможности.

## Проверки

`bash scripts/check.sh` проверяет gofmt и запускает go vet, go build и go test для Order Service.
Три service unit tests проверяют исходы available / unavailable / technical error;
поля создаваемого заказа, DTO mapper и доменные переходы ими не покрыты.
Ручная проверка Bruno подтверждает положительный POST, но не заменяет эти тесты.
GitHub Actions отдельно выполняет
build, lint, test и информационный security scan; его текущий `gosec -no-fail`
не является блокирующим security gate. Интеграционный контур добавляется с реальными
хранилищами и integration-тестами.

## Лицензия

[MIT](LICENSE).
