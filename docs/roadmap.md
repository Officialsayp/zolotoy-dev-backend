# План развития

Реорганизация не означает завершение Service Layer. Продуктовые этапы согласуются
с учебным циклом: упражнения не решаются заранее.

1. **Order HTTP/service.** Закрепить errors.Is и границу слоёв; добавить handler,
   service и domain tests на позитивные/отрицательные сценарии.
2. **Сохраняемый Order.** Согласовать OpenAPI/DTO, добавить repository, PostgreSQL,
   миграции и integration create/read после перезапуска. Compose — вместе с БД.
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

Новые пакеты и зависимости вводятся вместе с поведением. Полный переход сайта
с mock на real — отдельная задача после согласования контрактов.
