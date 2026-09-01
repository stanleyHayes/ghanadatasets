# GhanaDataset Registry

`GhanaDataset Registry` is an independent, metadata-only Digital Ghana discovery product. Its beta indexes official access paths, publishers, licence/access conditions and link status without copying source data. Canonical web/API surfaces remain pre-release until production evidence supports a lifecycle transition.

## Before implementation

1. Record the problem, users, non-goals, source rights and acceptance evidence in `agent_plan.md`.
2. Replace the placeholder source-register record only after authority and licence review.
3. Add domain contracts before transport or UI code.
4. Keep deployments fail-closed until required provider values exist.

## Verification

Run `pnpm install` and `pnpm check` for governance validation, Go tests/vet, metadata invariants, TypeScript checks and the production web build.
