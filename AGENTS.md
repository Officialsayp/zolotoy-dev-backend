# zolotoy-dev-backend

Go backend portfolio for the four-service console at zolotoy.dev.

- Read README.md and docs/architecture.md before structural work; inspect Git status
  and preserve local changes. services/order-service/AGENTS.md governs its learning cycle.
- Explicit implementation requests apply to their stated scope; preserve the remaining
  learning goals. Do not implement future lessons as part of housekeeping.
- Keep only implemented services under services/. Use a Go module per service;
  do not import another service's internal packages or share its database.
- Backend specifications linked in docs/architecture.md describe target behavior.
  Mock frontend DTOs are not accepted live contracts. Record unresolved details.
- pending-review/ holds historical work for the owner's decision. Do not delete,
  automatically restore, run, or treat it as current instructions/configuration.
- Run bash scripts/check.sh and git diff --check for code/layout changes.
  Documentation-only changes require link/path/fact checks and git diff --check.
- Preserve existing behavior during renames. Add regression tests for meaningful
  behavior changes; never claim a no-test-files result proves business correctness.
- Commit titles use Conventional Commits in English. Push/merge/deploy only within
  explicit user authorization; inspect workflows before publishing.
