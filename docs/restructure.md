# Разбор и реорганизация — 2026-09-14

Переписывать проект с нуля не нужно. Основная проблема — устаревшее описание,
зависимости эксперимента, пустые пакеты и неверная точка продолжения обучения.

| Было | Стало / причина |
| --- | --- |
| stockflow | zolotoy-dev-backend: backend технической консоли |
| order-service/ | services/order-service/: модуль первого сервиса |
| handlers/go.go | cmd/order-service/main.go: здесь действительно main |
| service/order-service.go | internal/service/order_service.go |
| domain/ | internal/domain/: все содержательные модели сохранены |
| Bruno/Stockflow local | api/bruno/order-service-local |
| github.com/MaximZolotoy/stockflow/... | github.com/Officialsayp/zolotoy-dev-backend/... |
| Echo/GORM dependencies | Убраны из активного модуля; snapshot рядом с экспериментом |
| repository/{memory,postgres,redis}, domain/errors.go | Удалены только package-заглушки без реализации |
| CRUD, 2 старых Bruno request, bootstrap config | pending-review/: решение об удалении за владельцем |
| README, SESSION_STATE, пустой PROJECT.md | Актуальное назначение и учебный контекст |

CI и check.sh переведены на новый путь. История Git и LESSONS сохранена; старые
названия в архиве и историческом журнале намеренно не подменены новым именем.

## Что доработать отдельно

- Domain не подключён к HTTP и не имеет автоматических тестов. Это наработки,
  а не мусор. Например, Cancel требует refunded даже для неоплаченного заказа:
  согласовать матрицу отмены и проверить её тестами.
- POST пока не сохраняет данные; нужны repository и согласованный контракт.
- Main/handlers разделить вместе с развитием транспорта и тестов.
- Сервер :8080 без production lifecycle/timeouts; перед внешним запуском нужны
  конфигурация http.Server, graceful shutdown и проверка отказов.
- Security использует gosec@latest и -no-fail: версия невоспроизводима, findings
  не блокируют CI. Нужны отдельный разбор findings и согласованный blocking gate.
  Уборка эту существующую проверку не ослабляет и не выдаёт за гарантию безопасности.

## Сохранность и продолжение

См. [список на разбор](../pending-review/README.md). Переносы проверяются через
`git diff --find-renames`; исходные файлы также остаются в Git history.
Восстанавливать выбранный материал отдельным коммитом, не весь старый каркас.
После GitHub rename обновить origin и открыть новый каталог в IDE. Старая ветка
сохраняется, разработка продолжается с main после fast-forward. Сайт и deployment
этой реорганизацией не изменяются.
