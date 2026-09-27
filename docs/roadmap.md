# План развития

Реорганизация не означает завершение Service Layer. Продуктовые этапы согласуются
с учебным циклом: упражнения не решаются заранее.

1. **Order HTTP/service.** Пройдены DTO vs Domain, request DTO → `CreateOrderInput`
   → `Money` / `OrderItem` → `domain.Order` → response DTO, способы оплаты и серверный UUID.
   Bruno подтвердил `201`, начальные статусы, `created_at` и сумму позиции `499000 × 2 = 998000`.
   Три service unit tests проверяют только исходы ошибок; нужны дальнейшие handler,
   service и domain tests. Создание заказа ещё не означает сохранение.
2. **Сохраняемый Order — следующий этап.** `OrderRepository` interface → memory repository
   → POST Save → GET by ID. Критерий memory-этапа: GET по UUID из POST возвращает тот же
   заказ в рамках процесса. Сейчас GET учебный, хранилища нет. Затем — PostgreSQL,
   миграции и integration create/read после перезапуска. Compose — вместе с БД;
   OpenAPI/DTO согласовать до live-интеграции.
3. **Lifecycle.** Матрица order/payment, отмена/refund, версии, идемпотентность,
   история. Критерий — тесты повторных и конкурирующих запросов.
4. **Auth и live frontend.** Credentials, сессии, rotation/reuse detection, RBAC;
   cookies/CORS/CSRF. Проверить реальные login/refresh/logout и доступ к заказам.
5. **Order events → Notification.** Outbox/event contract, durable inbox/jobs,
   worker/retry/recovery. Проверить дубли, crash windows и отказы provider.
6. **Shortener независимо.** DB/redirect → Redis/fallback → singleflight и измерения
   → ограниченная аналитика. Проверить invalidation/expiry и недоступность cache.
7. **Демонстрация и эксплуатация.** Runbook, метрики/логи, backup/restore, integration
   CI, реальные browser checks. Публиковать измерения с методикой, не обещанные SLA.

Точная учебная точка и порядок занятия — в
[SESSION_STATE](../services/order-service/.ai/SESSION_STATE.md).

Новые пакеты и зависимости вводятся вместе с поведением. Полный переход сайта
с mock на real — отдельная задача после согласования контрактов.
