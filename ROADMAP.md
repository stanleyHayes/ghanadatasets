# Roadmap

This roadmap is **directional, not a commitment**. It contains no delivery dates and no promises. Dates that appear below are verification dates already recorded elsewhere in this repository, never targets. It is a readable summary of the state and stated gates recorded elsewhere in this repository; the authoritative, machine-readable state lives in [`agent_plan.md`](agent_plan.md), and the boundary rules it must respect are in [`docs/adr/0001-product-boundary.md`](docs/adr/0001-product-boundary.md) and [`docs/product-definition.md`](docs/product-definition.md).

Lifecycle words used across Digital Ghana: proposed, building, beta, stable, externally blocked, deferred, retired. GhanaDataset Registry is currently in **public beta**, verified 2026-09-01.

## Now — shipped in the current public beta

Every item below is marked Done on the task board in [`agent_plan.md`](agent_plan.md), with evidence recorded in [`docs/runbooks/release-evidence.md`](docs/runbooks/release-evidence.md).

- [x] **Product definition and source review** (`P-0.1`) — GSS StatsBank, GSS Microdata Catalog and Ghana Open Data reviewed, with metadata-link-only publication decisions and record-specific access warnings.
- [x] **Domain contracts and fixtures** (`P-0.2`) — ten stable metadata records with licence, distribution, review-state and no-copied-data invariants.
- [x] **Metadata catalogue implementation** (`P-1.1`) — Go REST and constrained GraphQL, a dependency-free TypeScript client, a custom topic picker, and the Next.js catalogue, all passing the full local quality gate.
- [x] **Production release** (`P-2.1`) — canonical web and API TLS, CI, REST/GraphQL/CORS, browser UI and SEO, and provider rollback/restore evidence recorded.

Concretely, that means: `datasets.digitalghana.dev` and `api-datasets.digitalghana.dev` are live; dataset version `2026.09.01-beta.1` serves ten records, one visibly under link review; `pnpm check` is enforced in CI by [`.github/workflows/quality.yml`](.github/workflows/quality.yml); and rollback has been exercised on both Vercel and Render rather than assumed.

## Next — the gates that must clear before scope grows

There is no additional task claimed on the board. What follows are the conditions this repository has already written down as prerequisites. Each is a gate, not a scheduled feature.

| Gate | What it blocks | Blocking dependency |
|---|---|---|
| Resolve the record under link review | Retiring the `REVIEW` state for `godi-agriculture-catalogue` | The Ghana Open Data agriculture collection link timed out at release preflight and must be re-checked at source ([release evidence](docs/runbooks/release-evidence.md)) |
| Harvesting review | Any automated collection beyond hand-verified records | Robots, licences, rate limits and publisher expectations reviewed first ([product definition](docs/product-definition.md), non-goals) |
| Per-source licence and authority review | Adding any record from a new authority | An approved entry in [`docs/governance/source-register.json`](docs/governance/source-register.json) with an explicit publication decision; `blocked` and `unknown` decisions fail [`scripts/validate.rb`](scripts/validate.rb) |
| Record-count change | Growing the catalogue beyond ten records | Matching updates to the count invariants in `scripts/validate.rb`, [`tests/dataset.test.mjs`](tests/dataset.test.mjs) and [`internal/catalog/catalog_test.go`](internal/catalog/catalog_test.go), plus a `datasetVersion` bump |
| New hostname ADR | Any additional API or operational hostname | An explicit ADR and a portfolio registry entry before deployment ([ADR-0001](docs/adr/0001-product-boundary.md)) |
| Stable-release security baseline | Moving the lifecycle word from beta to stable | Dependency review, secret scanning, least-privilege credentials, rate limits where applicable, security headers, auditability for privileged changes and a tested rollback path ([`SECURITY.md`](SECURITY.md)) |
| Writable production state | Any correction workflow that writes at runtime | The beta has no writable production state and therefore no backup surface; a backup and restore plan would be a precondition ([operations runbook](docs/runbooks/operations.md)) |

Open engineering gaps that are safe to close without a new gate are listed as good first contributions in [`CONTRIBUTING.md`](CONTRIBUTING.md) — response schemas on the OpenAPI contract, a `get(id)` method on the client, the uncovered API test paths, a link-checker script, and documenting the GraphQL constraint.

## Later — deferred scope

Recorded in the repository as deliberately not part of beta, and not yet gated for work:

- **A publisher claim workflow.** Beta has "no publisher claim or privileged correction workflow". Any such workflow introduces identity, authority and write paths that the current boundary does not support.
- **A privileged correction path.** Corrections are reviewed in the open, by pull request, with a human approving publication. A privileged path would need the writable-state and auditability gates above.
- **Shared portfolio packages.** ADR-0001 accepts that "some configuration and small primitives may be repeated until two proven consumers justify a versioned shared package". Extraction waits for that second consumer.
- **Integration through the Digital Ghana gateway.** Cross-product integration happens only through versioned public contracts or pinned dataset artifacts. The unified gateway is deferred at the portfolio level, not scheduled here.

## Explicitly out of scope

These are product boundaries, not backlog. They are not expected to change, and a contribution that assumes otherwise will be declined. Drawn verbatim in substance from the non-goals in [`docs/product-definition.md`](docs/product-definition.md) and [`docs/adr/0001-product-boundary.md`](docs/adr/0001-product-boundary.md):

- Copying, mirroring or proxying source datasets, tables or microdata.
- Bypassing login, application, research-purpose or attribution conditions.
- Assuming portal-wide policy overrides a record-specific exception.
- Normalising differing source licences into a single broader permission.
- Claiming or implying that a reachable URL makes underlying data current, complete or accurate.
- Implying publisher endorsement, or any government affiliation or official status.
- Silently removing a failed official link or rewriting its provenance to imply availability.
- Becoming a comprehensive national database. This is a deliberately small, verified, source-linked subset.

## How this roadmap changes

By pull request, with evidence, alongside the change it describes. If this file and [`agent_plan.md`](agent_plan.md) ever disagree, `agent_plan.md` is correct and this file is stale.
