# LESSONS

## 2026-08-21

### Тема

GET handlers на стандартном `net/http`.

### Освоено

- запуск сервера, `ServeMux` и регистрация маршрута `GET /orders/{id}`;
- извлечение `id` через `r.PathValue`, преобразование в число и transport validation;
- response с корректными HTTP status codes;
- различие между path parameter (`/orders/42`) и query parameter (`?details=true`);
- обработка необязательного параметра `details` как значения `true` или `false`.

### Важные выводы

- `fmt.Printf` пишет в терминал сервера, а не клиенту; для response нужен `http.ResponseWriter`.
- В handler нельзя вызывать `log.Fatal` для ошибки одного request: это останавливает весь server.
- Все значения из URL сначала являются строками; их нужно явно преобразовывать и проверять.
- Отсутствующий необязательный query parameter — нормальная ситуация, не обязательно ошибка.
- В одном Go package не может быть две функции `main`; независимые приложения должны находиться в разных директориях.

### Следующее занятие

Завершить `GET /health` с `204 No Content`, затем разобрать `http.Handler` и `http.HandlerFunc`.

## 2026-08-22

### Завершено

- `GET /health` с `204 No Content` и без body;
- регистрация handler через `mux.Handle` и `http.HandlerFunc`.

### Важный вывод

`mux.HandleFunc` принимает обычную функцию, а `mux.Handle` ожидает `http.Handler`. `http.HandlerFunc` преобразует функцию в handler.

### Следующее занятие

Базовый `POST /orders`: method pattern и `201 Created` до разбора request body и JSON.

## 2026-08-22 — POST и Bruno

### Завершено

- базовый `POST /orders` с `201 Created`;
- проверка правильного method через `405 Method Not Allowed`;
- Bruno collection в проекте с двумя request и assertions на `201` и `405`.

### Важный вывод

Один Bruno request должен проверять один ожидаемый сценарий: нельзя одновременно ожидать `201` и `405` от одного response.

### Следующее занятие

Разобрать `r.Body` на простом текстовом POST request, затем перейти к JSON и DTO.

## 2026-08-22 — текстовый request body

### Завершено

- отправка text body из Bruno;
- чтение `r.Body` через `io.ReadAll`;
- преобразование байтов request body в строку и отправка response `received: keyboard`.

### Важный вывод

`r.Body` содержит входящие данные request. `io.ReadAll` удобно использовать в учебном примере, но в production body нужно ограничивать по размеру, чтобы клиент не мог отправить чрезмерно большой payload.

### Следующее занятие

JSON request body и DTO для `POST /orders`.

## 2026-08-22 — JSON request DTO

### Завершено

- DTO `createOrderRequest` для поля `product`;
- декодирование JSON из `r.Body` через `json.Decoder`;
- `400 Bad Request` при ошибке разбора JSON;
- проверка JSON request через Bruno.

### Важный вывод

JSON сначала должен быть декодирован в Go-struct. После этого handler работает с понятным полем `req.Product`, а не с исходным текстом JSON.

### Следующее занятие

JSON response: response DTO, `Content-Type: application/json` и `json.Encoder`.

## 2026-08-22 — JSON response

### Завершено

- response DTO `createOrderResponse`;
- заголовок `Content-Type: application/json`;
- сериализация response через `json.NewEncoder`;
- `201 Created` и JSON response в Bruno.

### Важный вывод

Для JSON response header `Content-Type` нужно установить до `WriteHeader`. Затем `json.Encoder` записывает body в `http.ResponseWriter`.

### Следующее занятие

Transport validation: JSON может быть синтаксически корректным, но обязательное поле `product` может быть пустым.

## 2026-08-22 — transport validation и начало Service Layer

### Завершено

- `POST /orders` возвращает `400 Bad Request`, если `product` пустой или состоит только из пробелов;
- `strings.TrimSpace` убирает пробелы в начале и конце строки для проверки обязательного поля;
- Bruno requests и assertions на `400` созданы для пустого `product`, `product` из пробелов и некорректного JSON;
- transport validation отделена от business validation;
- создан пакет `service` и `OrderService`;
- `OrderService` передаётся в `createOrderHandler` из `main` как зависимость;
- добавлено временное бизнес-правило: товар `unavailable` нельзя заказать.

### Важные выводы

- Handler отвечает за HTTP: чтение JSON, HTTP-коды и response.
- Service отвечает за правила приложения: например, можно ли заказать товар.
- Некорректный JSON и пустой `product` — ошибки формы входящего HTTP-запроса; отсутствие товара на складе — бизнес-правило.
- Повторный запуск failed GitHub Actions проверяет тот же старый коммит. Для исправления после merge нужен новый PR с новым коммитом.

### Точка остановки

- PR №10 с началом Service Layer уже смёржен, но его lint упал из-за форматирования `service/order_service.go`.
- Форматирование исправлено в отдельном коммите `6fdab38` (`minimal changes`), который ещё нужно отправить отдельным PR.
- После зелёного lint нужно вернуть безопасный внешний текст ошибки `product cannot be ordered` вместо `err.Error()` и исправить Bruno assertion `Create unavailable order`.

### Следующее занятие

Создать PR с форматированием, проверить зелёный lint и завершить первый сценарий business validation в Service Layer.

## 2026-09-27 — ProductAvailability и unit tests OrderService

### Завершено

- контракт ошибок `OrderService`: недоступный товар и техническая ошибка зависимости;
- `ProductAvailability` interface в пакете `service`, рядом с потребителем;
- constructor injection через `NewOrderService(productAvailability)`;
- учебный `memory.AvailabilityChecker` с правилом для товара `unavailable`;
- wiring в `main`: создание checker, передача в service и передача service в handler;
- ручная проверка доступного и недоступного товара через Bruno по итогам занятия;
- три unit-теста `OrderService` с fake-зависимостью: available → `nil`,
  unavailable → `ErrProductUnavailable`, technical error → ошибка зависимости.

### Важные выводы

- Service зависит от поведения интерфейса, конкретную реализацию выбирает `main`.
- Конструктор явно принимает зависимость; в тесте вместо memory checker можно передать fake.
- `false, nil` означает недоступный товар, а ненулевая ошибка — сбой самой проверки.
- Техническая ошибка не подменяется бизнес-ошибкой; handler отвечает за безопасный HTTP-ответ.
- Unit tests service проверяют его логику без HTTP-сервера и БД.

### Граница текущего результата

`CreateOrder` пока не создаёт `domain.Order` и не сохраняет заказ. Memory checker
проверяет доступность товара; будущий memory repository будет хранить заказы.
Доменные модели уже есть, но их интеграция с текущим HTTP/service flow ещё впереди.
Service Layer и Testing целиком не объявлены завершёнными.

### Следующее занятие

DTO vs Domain на текущих request/response DTO и `domain.Order`.
Затем настоящий `CreateOrder -> domain.Order -> OrderRepository -> memory -> PostgreSQL`;
с memory-хранилищем довести сценарий до сохранения и получения заказа.

## 2026-09-28 — DTO vs Domain и создание domain.Order

### Завершено

- разделены HTTP DTO, вход service и доменная модель;
- request DTO `createOrderV1Request` / `createOrderItemRequest` и mapper
  `toServiceInput()` → `service.CreateOrderInput`;
- `ProductAvailability.IsAvailable(productID string, quantity int64) (bool, error)`;
- domain mapping `CreateOrderInput.toDomainItems()` → `NewMoney` → `NewOrderItem`;
- `OrderItem` обновлён до `ProductID`, `ProductNameSnapshot`, `Quantity`, `UnitPrice`, `TotalPrice`;
- `PaymentMethod` обновлён до `prepaid` / `pay_on_receipt_online`, добавлен `ParsePaymentMethod`;
- UUID создаётся в service через `uuid.NewString()`;
- `OrderService.CreateOrder(input CreateOrderInput) (domain.Order, error)` создаёт и возвращает заказ;
- getters `Order`, копия среза в `Items()` и mapper `newCreateOrderResponse`;
- handler отправляет `createOrderResponse` с `201 Created`; `ErrInvalidOrder` даёт безопасный `400`;
- три прежних service unit tests адаптированы к новому входу и возвращаемому значению.

### Проверка в Bruno

По результату ручной проверки пользователя, `Post-request.yml` вернул `201 Created`,
серверный UUID, `status=created`, `payment_status=awaiting_payment`, `payment_method=prepaid`,
`created_at` и `items[0].total_price=998000` для `499000 × 2`.
В ответе также сохранены `buyer_id=buyer-123`, `delivery_address=Krasnodar`
и `buyer_comment=Call before delivery`.

Это проверка положительного HTTP-сценария. Три service unit tests проверяют исходы
ошибок, а не поля заказа; у нового Bruno request ещё нет автоматических assertions.
Старые Bruno-примеры с `{product}` относятся к предыдущему контракту.

### Важные выводы

- DTO описывает JSON; service input не зависит от HTTP; `domain.Order` хранит бизнес-состояние.
- Клиент не задаёт UUID, начальные статусы и время создания: их формирует сервер.
- `total_price` — сумма одной позиции; её считает domain, но имя и цена пока приходят из request.
- Availability checker принимает количество, но учебная реализация его пока игнорирует.
- Создание объекта и отправка JSON не означают сохранение заказа.

### Точка остановки и следующее занятие

Заказ пока **не сохраняется** и после запроса недоступен для повторного получения.
GET остаётся учебным ответом по числовому ID; запрос с UUID из POST сейчас вернёт `400`.
Repository и БД отсутствуют; весь Service Layer и Testing не объявлены завершёнными.

Далее: `OrderRepository` interface → memory repository → POST Save → GET by ID → PostgreSQL.
Начать с объяснения repository pattern, небольшого примера и одного задания на интерфейс;
не реализовывать будущие этапы заранее. Memory-сценарий должен возвращать по UUID тот же
заказ в рамках процесса; чтение после перезапуска проверяется на этапе PostgreSQL.
