# Проект

Order Service — первый Go backend в zolotoy-dev-backend для технической консоли
zolotoy.dev. См. [README](../README.md) и [архитектуру](../../../docs/architecture.md).
В Service Layer пройдены DTO vs Domain, преобразование request DTO → `CreateOrderInput`
→ доменные позиции и создание `domain.Order`. `POST /orders` возвращает response DTO
с серверным UUID, начальными статусами, суммой позиции и временем создания.
Заказ пока не сохраняется; GET остаётся учебным. Следующий этап — `OrderRepository`
interface → memory repository → POST Save → GET by ID → затем PostgreSQL.
Точная точка продолжения и границы проверок — в [SESSION_STATE](SESSION_STATE.md).
Создание заказа не означает готовность всех бизнес-правил или live-интеграции с frontend.
