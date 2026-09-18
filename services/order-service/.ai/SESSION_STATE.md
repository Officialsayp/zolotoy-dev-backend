# SESSION_STATE.md

## Текущий этап

Service Layer

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

## Состояние по коду на 2026-09-14

В main включён PR #17 (`95c8285`). Handler использует errors.Is для
ErrProductUnavailable, возвращает безопасные 400/500 и нормализованный product.
Старое поручение повторно создать PR после #10 больше не актуально.

## Текущая задача обучения

Продолжить Service Layer: разобрать контракт ошибок и границу HTTP/business logic.
Не считать весь этап завершённым только по наличию кода. Следующее упражнение
согласовать в учебном диалоге; PostgreSQL и другие будущие уроки заранее не решены.

## Переезд структуры

- Репозиторий: zolotoy-dev-backend (бывший stockflow).
- Модуль: services/order-service; запуск: go run ./cmd/order-service.
- Handler/main: cmd/order-service/main.go; service: internal/service/order_service.go.
- Доменные наработки: internal/domain; Bruno: api/bruno/order-service-local.
- Старые эксперименты: ../../pending-review относительно модуля.
- История LESSONS.md сохранена; её старые пути и PR относятся к прошлым занятиям.
