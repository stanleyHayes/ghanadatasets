# Contributing to GhanaDataset Registry

Thank you for considering a contribution. This registry is public-interest infrastructure: its value is entirely in whether people can trust what it says. Contributions to code, contracts, documentation and — especially — data provenance are all welcome.

Please read [`README.md`](README.md), [`docs/product-definition.md`](docs/product-definition.md) and [`docs/adr/0001-product-boundary.md`](docs/adr/0001-product-boundary.md) before making a substantive change. Automated contributors must also read [`AGENTS.md`](AGENTS.md).

## Before you start

1. Check [`agent_plan.md`](agent_plan.md) — it is the authoritative task board and lifecycle record — and [`ROADMAP.md`](ROADMAP.md) for direction.
2. For data work, identify the source authority, licence, retrieval date and permitted use **before** adding anything. If the authority is not already in [`docs/governance/source-register.json`](docs/governance/source-register.json), the source review comes first and is its own pull request.
3. Open an issue for anything that changes a public contract, a stable ID, the record count, or the product boundary. Small fixes can go straight to a pull request.

## Prerequisites

| Tool | Version | Where it is pinned |
|---|---|---|
| Node.js | 22 | [`.github/workflows/quality.yml`](.github/workflows/quality.yml) |
| pnpm | 10.17.1 | `packageManager` in [`package.json`](package.json) |
| Go | 1.23 | [`go.mod`](go.mod) |
| Ruby | 3.4 | [`.github/workflows/quality.yml`](.github/workflows/quality.yml) |

Ruby is required only for the governance validator, which runs as the first step of `pnpm check`.

## Local setup

```sh
git clone https://github.com/stanleyHayes/ghanadatasets.git
cd ghanadatasets
pnpm install
```

Run the catalogue web app:

```sh
pnpm dev          # http://localhost:3000
```

Run the API against the same fixture, in a second terminal:

```sh
go build -trimpath -o bin/ghanadatasets-api ./cmd/api
./bin/ghanadatasets-api
```

`PORT` defaults to `8080`, `DATASETS_DATA_PATH` to `data/datasets.json`, and `ALLOWED_ORIGIN` to `https://datasets.digitalghana.dev`. Set `ALLOWED_ORIGIN=http://localhost:3000` if you need the local web app to call the local API across origins.

## Verification

Every pull request must pass the same single gate that CI runs:

```sh
pnpm check
```

That is the composite command. Its parts, which you can run individually while iterating:

```sh
ruby scripts/validate.rb   # required files, source-register decisions, record count, secret scan
go test ./...
go vet ./...
pnpm typecheck             # tsc --noEmit
pnpm test                  # node --test tests/*.test.mjs
pnpm build                 # production Next.js build
```

`pnpm check` is deliberately strict. [`scripts/validate.rb`](scripts/validate.rb) asserts that required governance files exist, that no source carries a `blocked` or `unknown` publication decision, that the fixture holds exactly the agreed number of records, and that nothing resembling a private key has been committed.

## Commit and pull-request conventions

Commits follow Conventional Commits, lowercase, imperative mood, no trailing full stop — matching this repository's history:

```
feat: build metadata-only GhanaDataset beta
docs: record GhanaDataset public beta evidence [skip ci]
chore: ignore local Vercel state [skip ci]
```

Use `feat`, `fix`, `docs`, `chore`, `test`, `refactor`. Append `[skip ci]` only to commits that touch nothing the quality gate can meaningfully check.

A pull request should state:

- **Scope** — what changed and which surface it affects (web, API, contracts, data, docs).
- **Source or ADR references** — for any data or boundary change.
- **Verification performed** — paste the `pnpm check` result.
- **Migration and rollback impact** — especially for anything touching stable IDs, `datasetVersion` or a published contract.
- **Known external gates** — anything blocked on a publisher, provider or upstream decision.

Keep changes path-bounded. Do not bundle a data correction with an unrelated refactor.

## Proposing a data correction

Corrections need evidence, not confidence. Open an issue or pull request containing:

| Field | What to provide |
|---|---|
| Stable ID | e.g. `godi-agriculture-catalogue` — never propose changing an ID to fix a title |
| Current value | The field and its value as published today |
| Proposed value | The corrected value, exactly as it should appear |
| Authoritative source | A URL at the publishing authority, not a secondary report |
| Source publication date | When the authority published or last updated it |
| Historical impact | Whether the correction changes a previously published record |

Rules the review will apply:

- Metadata and official links only. No source rows, tables, observations or microdata may enter [`data/datasets.json`](data/datasets.json); [`tests/dataset.test.mjs`](tests/dataset.test.mjs) enforces this.
- Every `accessUrl` must be HTTPS and on an official host already approved in the source register.
- Licence and access conditions are recorded per record. Never normalise a record up to portal-wide policy, and never render an unknown licence as open.
- A broken official link becomes `status: "REVIEW"` with a `statusNote`. It is not deleted, and its provenance is not rewritten to imply the source is available.
- Changing the record count means updating the count assertions in [`scripts/validate.rb`](scripts/validate.rb), [`tests/dataset.test.mjs`](tests/dataset.test.mjs) and [`internal/catalog/catalog_test.go`](internal/catalog/catalog_test.go) in the same pull request, and bumping `datasetVersion`.

Automation may draft a correction; a human reviewer approves canonical publication.

## Review expectations

- Reviewers check evidence before code. A data change with a weak source will be held regardless of how clean the diff is.
- Status language stays honest: proposed, building, beta, stable, externally blocked, deferred, retired. Do not label something stable to close a task.
- Public contracts in [`contracts/`](contracts/) change with a version note and, for anything breaking, a migration path.
- Tests should be proportional to risk. A new invariant on the fixture belongs in the Go or Node test suite, not only in a reviewer's head.
- Never commit credentials, private exports, personal data or provider environment files. Security issues go privately to the private channel in `SECURITY.md` per [`SECURITY.md`](SECURITY.md) — never in a public issue.
- Never imply government endorsement or official status anywhere in code, copy, metadata or commit messages.

Participation is governed by [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md).

## Good first contributions

These are real, currently open gaps in this repository:

1. **Re-check the record under review.** `godi-agriculture-catalogue` is marked `REVIEW` because the Ghana Open Data agriculture collection timed out during release preflight. Re-check the official link, and propose either a status change with evidence or an updated `statusNote` and `checkedAt`.
2. **Give the OpenAPI spec real response schemas.** [`contracts/openapi.yaml`](contracts/openapi.yaml) currently describes its `200` responses in prose only. Add `components.schemas` for `DatasetRecord`, `Distribution` and the search envelope, derived from [`internal/catalog/catalog.go`](internal/catalog/catalog.go) and [`sdk/typescript/index.ts`](sdk/typescript/index.ts), so the two stay provably in step.
3. **Add a `get(id)` method to the TypeScript client.** The SDK exposes `search()` only. Add single-record retrieval against `GET /v1/datasets/{id}`, including a typed result for the `DATASET_NOT_FOUND` 404.
4. **Close the API test gaps.** [`cmd/api/main_test.go`](cmd/api/main_test.go) covers REST search, allowed-origin CORS and GraphQL parity. It does not yet cover `/health`, the `404 DATASET_NOT_FOUND` path, the GraphQL rejection path for undocumented queries, or the denied-origin CORS case that the release runbook smokes by hand.
5. **Write a link-checker script.** [`docs/runbooks/operations.md`](docs/runbooks/operations.md) relies on manual smoke checks. A script that re-checks each `accessUrl`, reports status changes and *proposes* a `checkedAt` update — without auto-committing to `data/datasets.json` — would make the review state maintainable.
6. **Document the GraphQL constraint.** [`contracts/README.md`](contracts/README.md) says the query is constrained but does not say how. Document the supported `datasets(q, topic)` query, the 32 KiB body cap, and the fact that anything else is rejected with `400`.

## Licensing

Unless explicitly stated otherwise, contributions intentionally submitted for inclusion are provided under the Apache License 2.0 on the licence's inbound=outbound terms — see [`LICENSE`](LICENSE) and [`NOTICE`](NOTICE).

Contributions must not include third-party data or documents without recorded permission. Referenced source documents and third-party datasets retain their own rights and are not relicensed by being indexed here.
