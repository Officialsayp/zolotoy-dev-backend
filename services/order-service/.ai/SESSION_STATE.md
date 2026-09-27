# SESSION_STATE.md

## Текущий этап

Service Layer: ProductAvailability и базовые unit tests пройдены.
Следующая тема — DTO vs Domain.

## Изучено

- запуск HTTP-сервера через `http.ListenAndServe`;
- маршрутизация через `http.ServeMux` и method pattern;
- handler с сигнатурой `func(http.ResponseWriter, *http.Request)`;
- path parameter через `r.PathValue`;
- query parameter через `r.URL.Query().Get`;
- преобразование строковых параметров через `strconv.Atoi` и `strconv.ParseBool`;
- `200 OK`, `204 No Content`, `400 Bad Request`, `405 Method Not Allowed`;
- отправка response и заголовка `Content-Type`.
- `GET /health` и статус `204 No Content` без body;
- разница между `mux.HandleFunc` и `mux.Handle` с `http.HandlerFunc`.
- `POST /orders` и статус `201 Created`;
- ручная проверка endpoints и assertions в Bruno.
- отправка текстового request body из Bruno;
- чтение текстового body через `r.Body` и `io.ReadAll`.
- JSON request body, `json.Decoder` и request DTO.
- JSON response через response DTO и `json.Encoder`.
- transport validation поля `product`: пустая строка и строка только из пробелов возвращают `400 Bad Request`;
- Bruno assertions для `400` при пустом `product`, пробелах и некорректном JSON;
- различие transport validation и business validation;
- начальная связка `main -> handler -> OrderService` через передачу зависимости в `createOrderHandler`.
- контракт ошибок: `ErrProductUnavailable`, `errors.Is` и безопасные HTTP-ответы `400`/`500`;
- интерфейс `ProductAvailability` рядом с потребителем в пакете `service`;
- constructor injection через `NewOrderService(productAvailability)`;
- учебная memory-реализация `AvailabilityChecker`: `unavailable` недоступен, остальные товары доступны;
- wiring в `main`: `memory.AvailabilityChecker -> NewOrderService -> createOrderHandler`;
- ручная проверка доступного и недоступного товара через Bruno;
- unit tests `OrderService` с `fakeProductAvailability`: available, unavailable и technical error;
- передача технической ошибки зависимости вызывающему коду и проверка через `errors.Is`.

## Состояние по коду на 2026-09-27

Рабочая ветка: `feature/service-error-contract`.

- `internal/service/order_service.go` — `ProductAvailability`, конструктор и проверка доступности.
- `internal/availability/memory/checker.go` — учебная реализация интерфейса.
- `cmd/order-service/main.go` — сборка зависимостей и HTTP handlers.
- `internal/service/order_service_test.go` — три unit-теста исходов `CreateOrder`.
- Handler нормализует `product`, сопоставляет `ErrProductUnavailable` с безопасным `400`,
  а техническую ошибку — с безопасным `500`.

`CreateOrder(product string) error` пока только проверяет доступность товара:
`domain.Order` не создаётся и не сохраняется. `memory.AvailabilityChecker` —
проверка доступности, а не хранилище заказов. `GET /orders/{id}` остаётся учебным
ответом без чтения из repository. Этот этап ещё не заменяет mock на zolotoy.dev.

## Текущая задача обучения

Разобрать DTO vs Domain на существующих `createOrderRequest`, `createOrderResponse`
и `internal/domain.Order`: данные HTTP-запроса/ответа, бизнес-состояние и правила,
границы преобразования между ними. Сначала объяснение и небольшой пример,
затем одно самостоятельное задание.

Дальнейшая последовательность:

1. Настоящий `CreateOrder`, который создаёт `domain.Order`.
2. Интерфейс `OrderRepository` и memory-хранилище: сохранить заказ и получить его обратно.
3. PostgreSQL-реализация хранения после освоения memory-сценария.

Не считать весь Service Layer или Testing завершёнными по этим трём тестам.
Создание заказа, repository и PostgreSQL остаются будущими учебными задачами;
их реализацию заранее не писать.

## Переезд структуры

- Репозиторий: zolotoy-dev-backend (бывший stockflow).
- Модуль: services/order-service; запуск: go run ./cmd/order-service.
- Handler/main: cmd/order-service/main.go; service: internal/service/order_service.go.
- Доменные наработки: internal/domain; Bruno: api/bruno/order-service-local.
- Старые эксперименты: ../../pending-review относительно модуля.
- История LESSONS.md сохранена; её старые пути и PR относятся к прошлым занятиям.
