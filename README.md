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
| Order | Заказы, оплата, переходы состояний, идемпотентность и outbox | Локальный HTTP/service flow; отдельные доменные модели; без хранилища |
| Auth | Пользователи, сессии, refresh rotation и RBAC | План; кода пока нет |
| Notification | Событие → задание → попытки доставки, retries и dead jobs | План; кода пока нет |
| URL Shortener | Короткие ссылки, redirect, Redis cache и аналитика | План; кода пока нет |

Это работа в процессе, не готовая production-платформа. Текущий учебный этап Order —
Service Layer. PostgreSQL, Kafka, Redis, авторизация и live-интеграция с frontend
ещё не подключены. Kubernetes и API Gateway не нужны для начала разработки.

## Быстрый старт

Нужен Go версии не ниже указанной в [go.mod](services/order-service/go.mod).
Текущий сервис использует только стандартную библиотеку; Docker и БД не требуются.

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
  -d '{"product":"keyboard"}'
```

Ожидаются 204 и 201 с `{"product":"keyboard"}` соответственно.
POST пока **не сохраняет заказ**, а демонстрирует валидацию и вызов service.
Подробности и Bruno — в [README Order Service](services/order-service/README.md).

## Структура

```text
services/order-service/   единственный реализованный Go-модуль
  cmd/order-service/      запуск сервера и текущие handlers
  internal/domain/        модели заказа, денег и оплаты
  internal/service/       бизнес-проверки
  api/bruno/              актуальная ручная коллекция
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

Сначала довести один проверяемый Order-сценарий до хранилища и согласованного API,
затем подключать остальные темы по этапам. Не добавлять инфраструктуру ради списка
технологий и не выдавать целевые требования за реализованные возможности.

## Проверки

`bash scripts/check.sh` запускает gofmt, go vet, go build и go test для Order Service.
В исходном проекте отсутствовали автоматические тесты: успешная сборка сама по себе
не доказывает корректность доменных переходов. GitHub Actions отдельно выполняет
build, lint, test и информационный security scan; его текущий `gosec -no-fail`
не является блокирующим security gate. Интеграционный контур добавляется с реальными
хранилищами и integration-тестами.

## Лицензия

[MIT](LICENSE).
