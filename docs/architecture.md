# Архитектура zolotoy.dev

## Основания

Проверено 2026-09-14: публичная консоль показывает четыре сервиса и `MOCK production`
/ Simulated. Это MSW demo, а не подтверждение работающих Go API. Backend проверен
по main `95c8285` (PR #17).

Планы прочитаны в frontend на коммите `df234de827f48f12a20db9654f42d517fddd466b`:

- [Общий план](https://github.com/Officialsayp/zolotoy-dev-frontend/blob/df234de827f48f12a20db9654f42d517fddd466b/docs/frontend/MASTER_FRONTEND_PLAN.md)
- [Order](https://github.com/Officialsayp/zolotoy-dev-frontend/blob/df234de827f48f12a20db9654f42d517fddd466b/docs/backend-specs/01_order_service.md)
- [Auth](https://github.com/Officialsayp/zolotoy-dev-frontend/blob/df234de827f48f12a20db9654f42d517fddd466b/docs/backend-specs/02_auth_service.md)
- [Notification](https://github.com/Officialsayp/zolotoy-dev-frontend/blob/df234de827f48f12a20db9654f42d517fddd466b/docs/backend-specs/03_notification_service.md)
- [Shortener](https://github.com/Officialsayp/zolotoy-dev-frontend/blob/df234de827f48f12a20db9654f42d517fddd466b/docs/backend-specs/04_url_shortener.md)

Перед реализацией сравнивать с актуальными версиями. Не копируем большие планы,
чтобы не создавать расходящиеся источники истины. MD описывают цель; канонического
OpenAPI здесь пока нет. Mock DTO не равен согласованному live-контракту.

## Решение о размещении

Сохраняем историю stockflow в `zolotoy-dev-backend`, frontend остаётся отдельно.
Исходные service MD предлагали отдельные репозитории: здесь осознанно меняется
только размещение кода. Для одного разработчика monorepo упрощает навигацию и
согласованные изменения. Цена — общий CI и необходимость модульных проверок.
Четыре отдельных backend-репозитория сейчас добавили бы сопровождение без готовых
сервисов; выделить их позже можно с сохранением истории.

Каждый реализованный сервис имеет свой go.mod и internal. Нет прямых запросов
в чужую БД и импорта чужих внутренних пакетов. Go workspace и общие библиотеки
добавляются при реальной необходимости. Пустых service-директорий не создаём.

| Сервис | Целевая ответственность | Не раздувать до |
| --- | --- | --- |
| order-service | Order/payment state machines, история, версии, идемпотентность, outbox | Магазина, каталога, платёжного процессинга |
| auth-service | Identity, сессии, refresh rotation, RBAC | Enterprise IAM и social login |
| notification-service | Durable inbox/jobs, workers, retries, recovery | Маркетинговой платформы |
| url-shortener | Ссылки, redirect, DB, cache-aside, аналитика | Шардинга и отдельных analytics-сервисов |

Inventory/API Gateway из старого README не нужны для начала текущего плана.
PostgreSQL/Kafka/Redis появляются с конкретным сценарием и тестами. Kubernetes
и полный observability stack не являются условием первого работающего сервиса.

## Реестр незакрытых контрактов

| Область | Решить до live-интеграции |
| --- | --- |
| Общая | Cookie/CORS/CSRF, error envelope, Auth user/admin и роли Order, health DTO |
| Order | /api/v1 DTO, list/history, fulfillment actions, отмена/refund, 400/409, доверие к buyer_id |
| Auth | Cookie/body login/refresh/logout, sessions DTO, регистрация с сессией или отдельный login |
| Notification | Pagination/filter, jobs/attempts DTO, manual retry states и история; Kafka DLQ отдельно от dead jobs |
| Shortener | Ownership, anonymous create, update/status/analytics DTO, 404/410; hostname из short_url |

Порядок: решение в backend MD → OpenAPI/контрактные проверки → frontend adapter
→ реальный сценарий. Учебный `{product}` не подменяет целевой контракт заказа.
