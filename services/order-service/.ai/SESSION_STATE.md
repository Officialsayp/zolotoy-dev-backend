# SESSION_STATE.md

## Текущий этап

Service Layer: DTO vs Domain и создание `domain.Order` из HTTP-запроса пройдены.
`POST /orders` возвращает созданный заказ с серверным UUID, но пока не сохраняет его.
Следующая тема — `OrderRepository`: interface → memory repository → POST Save → GET by ID.
PostgreSQL — только после работающего сценария сохранения и чтения в памяти.

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
- transport validation на прежнем DTO с полем `product`: пустая строка и пробелы → `400 Bad Request`;
- Bruno assertions для прежнего `product` и некорректного JSON; примеры с `{product}` теперь исторические;
- различие transport validation и business validation;
- начальная связка `main -> handler -> OrderService` через передачу зависимости в `createOrderHandler`.
- контракт ошибок: `ErrProductUnavailable`, `errors.Is` и безопасные HTTP-ответы `400`/`500`;
- интерфейс `ProductAvailability` рядом с потребителем в пакете `service`;
- constructor injection через `NewOrderService(productAvailability)`;
- учебная memory-реализация `AvailabilityChecker`: `unavailable` недоступен, остальные товары доступны;
- wiring в `main`: `memory.AvailabilityChecker -> NewOrderService -> createOrderHandler`;
- ручная проверка доступного и недоступного товара через Bruno;
- unit tests `OrderService` с `fakeProductAvailability`: available, unavailable и technical error;
- передача технической ошибки зависимости вызывающему коду и проверка через `errors.Is`;
- DTO vs Domain: HTTP-формат, вход service и бизнес-модель имеют разные задачи;
- request DTO `createOrderV1Request` и `createOrderItemRequest` вместо прежнего `{product}`;
- mapper `createOrderV1Request.toServiceInput()` → `service.CreateOrderInput`;
- контракт `ProductAvailability.IsAvailable(productID string, quantity int64) (bool, error)`;
- domain mapping `CreateOrderInput.toDomainItems()` → `domain.NewMoney` → `domain.NewOrderItem`;
- `OrderItem`: `ProductID`, `ProductNameSnapshot`, `Quantity`, `UnitPrice`, `TotalPrice`;
- `PaymentMethod`: `prepaid` / `pay_on_receipt_online`, проверка через `ParsePaymentMethod`;
- серверная генерация UUID в service через `uuid.NewString()`;
- `OrderService.CreateOrder(input CreateOrderInput) (domain.Order, error)` создаёт и возвращает заказ;
- getters `Order` и mapper `newCreateOrderResponse(order domain.Order)` → `createOrderResponse`;
- `ErrInvalidOrder` и безопасный HTTP-ответ `400` для ошибок создания доменного заказа;
- успешная ручная проверка нового `POST /orders` в Bruno по итогам занятия.

## Состояние по коду на 2026-09-28

Рабочая ветка: `feature/service-error-contract`.

- `cmd/order-service/create_order_dto.go` — request/response DTO и оба transport mapper.
- `cmd/order-service/main.go` — wiring, transport validation и HTTP handlers.
- `internal/service/order_service.go` — `CreateOrderInput`, `ProductAvailability`,
  проверка доступности каждой позиции, разбор способа оплаты, UUID и создание `domain.Order`.
- `internal/service/create_order.go` — преобразование входных позиций в `Money` и `OrderItem`.
- `internal/domain/order_item.go` — snapshot имени товара и расчёт `TotalPrice = UnitPrice × Quantity`.
- `internal/domain/payments.go` — два допустимых способа оплаты и `ParsePaymentMethod`;
  отдельный `PaymentTiming` больше не используется.
- `internal/domain/order.go` — `NewOrder` задаёт `created`, `awaiting_payment` и `created_at`;
  getters `ID`, `BuyerID`, `Status`, `PaymentStatus`, `PaymentMethod`, `Items`,
  `DeliveryAddress`, `BuyerComment`, `CreatedAt` позволяют собрать response;
  `Items()` возвращает копию среза.
- `internal/availability/memory/checker.go` — учебная проверка: `product_id=unavailable`
  недоступен, остальные доступны; аргумент `quantity` пока не используется.
- `internal/service/order_service_test.go` — три теста адаптированы к `CreateOrderInput`
  и возврату `(domain.Order, error)`; проверяют только исходы ошибок, поля заказа не проверяют.

Текущий путь данных:

```text
JSON → createOrderV1Request → toServiceInput() → CreateOrderInput
     → toDomainItems() → Money → OrderItem
     → ParsePaymentMethod → ProductAvailability(productID, quantity)
     → серверный UUID → domain.NewOrder → domain.Order
     → getters → newCreateOrderResponse → createOrderResponse → 201 JSON
```

Handler проверяет обязательные поля и положительное количество. `ErrInvalidOrder`
и `ErrProductUnavailable` сопоставляются с безопасным `400`, технические ошибки — с `500`.
Название и цена товара пока приходят из request; сервер вычисляет сумму позиции,
но не получает цену из каталога и не резервирует остатки.

## Ручная проверка Bruno

По результату занятия, зафиксированному пользователем, запрос
`api/bruno/order-service-local/Post-request.yml` успешно вернул:

- `201 Created` и `Content-Type: application/json`;
- `id` — сгенерированный сервером UUID;
- `status=created`, `payment_status=awaiting_payment`, `payment_method=prepaid`;
- `items[0].total_price=998000` для `unit_price=499000` и `quantity=2`;
- `buyer_id=buyer-123`, `delivery_address=Krasnodar`, `buyer_comment=Call before delivery`;
- серверный `created_at`.

`total_price` — сумма позиции внутри `items`, не отдельное поле суммы всего заказа.
У `Post-request.yml` пока нет автоматических assertions; старые Bruno-запросы с
`{product}` требуют актуализации и не подтверждают новый контракт.

## Граница результата

Заказ **не сохраняется**: `CreateOrder` создаёт объект и handler отправляет его в ответе,
но после запроса заказ не остаётся в хранилище и получить его повторно нельзя.
`OrderRepository`, memory repository и PostgreSQL ещё не реализованы.
`memory.AvailabilityChecker` проверяет доступность, а не хранит заказы.

`GET /orders/{id}` всё ещё учебный: принимает положительный числовой ID, возвращает
текст и не обращается к repository. UUID из POST сейчас приведёт к `400 Bad Request`.
Создание `domain.Order` не означает завершение всех доменных правил, Service Layer,
Testing или live-интеграции с zolotoy.dev. Автоматических проверок полей нового ответа,
DTO mapper, UUID, суммы позиции и `ErrInvalidOrder` пока нет.

## Текущая задача обучения

Разобрать repository pattern на уже создаваемом `domain.Order`: сначала объяснение
и небольшой пример, затем одно самостоятельное задание на контракт `OrderRepository`
рядом с потребителем в пакете `service`. Реализацию будущих уроков заранее не писать.

Дальнейшая последовательность:

1. Интерфейс `OrderRepository` для сохранения и получения заказа.
2. Memory repository и передача зависимости в `OrderService`.
3. `POST /orders` → `CreateOrder` → `Save`.
4. `GET /orders/{id}` → получение по UUID → сохранённый `domain.Order` → JSON.
5. После работающего POST → Save → GET — PostgreSQL, миграции и проверка чтения после перезапуска.

Критерий memory-этапа: POST создаёт и сохраняет заказ, а GET по возвращённому UUID
возвращает тот же заказ в рамках работающего процесса. Потеря данных при перезапуске
memory-хранилища ожидаема; постоянное хранение относится к следующему этапу PostgreSQL.

## Переезд структуры

- Репозиторий: zolotoy-dev-backend (бывший stockflow).
- Модуль: services/order-service; запуск: go run ./cmd/order-service.
- Handler/main: cmd/order-service/main.go; service: internal/service/order_service.go.
- Доменные наработки: internal/domain; Bruno: api/bruno/order-service-local.
- Старые эксперименты: ../../pending-review относительно модуля.
- История LESSONS.md сохранена; её старые пути и PR относятся к прошлым занятиям.
