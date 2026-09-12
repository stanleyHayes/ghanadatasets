# GhanaDataset Registry

A provenance-first, metadata-only catalogue of official Ghanaian public datasets — what exists, who publishes it, what the licence and access conditions are, and the exact machine-readable path to reach it at source.

[![Licence: Apache-2.0](https://img.shields.io/badge/licence-Apache--2.0-blue.svg)](LICENSE)
[![Web](https://img.shields.io/badge/web-datasets.digitalghana.dev-0b7285)](https://datasets.digitalghana.dev)
[![API](https://img.shields.io/badge/API-api--datasets.digitalghana.dev-0b7285)](https://api-datasets.digitalghana.dev/health)
[![Quality](https://img.shields.io/github/actions/workflow/status/stanleyHayes/ghanadatasets/quality.yml?branch=main&label=quality)](https://github.com/stanleyHayes/ghanadatasets/actions/workflows/quality.yml)
[![Go](https://img.shields.io/badge/Go-1.23-00ADD8)](go.mod)
[![Next.js](https://img.shields.io/badge/Next.js-16-000000)](package.json)

## Live now

| Surface | URL | State |
|---|---|---|
| Catalogue web app | <https://datasets.digitalghana.dev> | Live |
| Read-only REST/GraphQL API | <https://api-datasets.digitalghana.dev> | Live |
| Health and dataset version | <https://api-datasets.digitalghana.dev/health> | Live |

The API is hosted on Render's free plan and sleeps when idle. **The first request after a quiet period can take 30–60 seconds.** Subsequent requests are fast.

```sh
# Search the catalogue. Real response, verbatim.
curl -s "https://api-datasets.digitalghana.dev/v1/datasets?q=trade"
```

```json
{"count":1,"coverage":"metadata-only verified subset","data":[{"id":"gss-trade-statsbank","title":"Trade StatsBank","publisher":"Ghana Statistical Service","topic":"Trade","description":"Official aggregated trade tables exposed through the StatsBank catalogue.","licence":"Source terms apply","accessCondition":"Public aggregated-table access; verify source terms.","sourceId":"gss-statsbank","checkedAt":"2026-09-01","status":"REACHABLE","distributions":[{"format":"PXWEB_API","accessUrl":"https://statsbank.statsghana.gov.gh/api/v1/en/Trade"}]}],"datasetVersion":"2026.09.01-beta.1"}
```

That `accessUrl` is the live PxWeb endpoint at Ghana Statistical Service. This registry hands you the path; the data stays at source.

## The problem this solves

### For developers

Ghana's official data is public, but it is not discoverable in any way a program can use.

If you want, say, aggregated trade figures or the 2021 census public-use microdata, the work before this registry existed looked like this:

- Open [statsbank.statsghana.gov.gh](https://statsbank.statsghana.gov.gh/pxweb/en/) and click through a PxWeb tree until you find the database. PxWeb exposes a REST API, but the database identifier is a human-typed display name that has to be URL-encoded by hand — `Ghana%20Census%20of%20Agriculture%20(GCA)`, `Annual%20Household%20Income%20and%20Expenditure%20Survey%20(AHIES)`. Nothing tells you these path segments exist; you reverse-engineer them from the UI.
- Open the [GSS Microdata Catalog](https://microdata.statsghana.gov.gh/index.php/catalog/central) separately, because it is a different platform with different identifiers and, critically, **per-study access conditions** that are not the same as the aggregated-table terms next door.
- Open [data.gov.gh](https://data.gov.gh/) separately again, where portal policy generally points at ODC Attribution but individual records carry exceptions — so the portal-wide statement is not a safe answer for any single dataset.
- Then keep the result alive in a spreadsheet, a `constants.ts`, or a comment, and re-click every link by hand whenever something breaks.

The cost is not theoretical. It is hours per dataset of manual archaeology, an untracked copy of licence text that drifts from the source, hardcoded URLs that rot silently, and — the expensive defect — reuse decisions made against the wrong terms because the aggregated-table conditions were quietly applied to microdata.

What this project gives you instead:

- **Stable registry IDs** (`gss-phc-2021-microdata`, `gss-trade-statsbank`) that do not change when a title or a source URL changes.
- **Licence and access condition on the same record as the link**, per dataset, never normalised upward into a broader permission.
- **Typed machine access paths** with an explicit `format` — `PXWEB_API`, `CATALOG_RECORD`, `CATALOG_COLLECTION`, `WEB` — so you can filter for the ones a program can actually consume.
- **A `checkedAt` date and a `status`** so a broken official link is visible as `REVIEW` rather than silently deleted.
- **Three interfaces over one fixture**: REST, a constrained GraphQL query, and a dependency-free TypeScript client in [`sdk/typescript/index.ts`](sdk/typescript/index.ts).

### For the community and public interest

Discovery infrastructure decides what gets analysed. A dataset nobody can find is, for practical purposes, unpublished — and journalists, researchers and civic technologists in Ghana routinely re-do the same portal archaeology in parallel, privately, with no shared record of what they found or under what terms.

Making that record open, source-linked and independent matters for reasons that outlive this codebase:

- **Provenance over convenience.** Every record names its source authority and review date. You can check our work against the original; we cannot quietly become the authority.
- **Reproducibility.** The catalogue is a versioned, immutable fixture ([`data/datasets.json`](data/datasets.json), version `2026.09.01-beta.1`). A citation of a record resolves to a specific state of the registry.
- **No lock-in.** Apache-2.0, no account, no key, no rate-limited tier, no proprietary schema. The client is a single dependency-free file built on `fetch`. If this project stops, your integration still points at government endpoints you now know about.
- **Honesty as a feature.** One of the ten records is publicly marked `REVIEW` because its official link timed out during release preflight. It stays visible with its provenance intact rather than being dropped to make the numbers look clean.

### What this is not

Drawn from the stated non-goals in [`docs/product-definition.md`](docs/product-definition.md) and [`docs/adr/0001-product-boundary.md`](docs/adr/0001-product-boundary.md):

- **Not a data mirror.** No source tables, rows, observations or microdata are copied or proxied. The fixture is asserted to be metadata-only by a test in [`tests/dataset.test.mjs`](tests/dataset.test.mjs).
- **Not a way around access conditions.** Login, application, research-purpose and attribution requirements stay with the publisher and are reproduced, not removed.
- **Not a licence authority.** Portal-wide policy never overrides a record-specific exception here, and an unknown licence is shown as unknown — never as open.
- **Not a freshness guarantee.** A reachable URL means the link resolved on `checkedAt`. It says nothing about whether the underlying data is current, complete or correct.
- **Not comprehensive.** Ten deliberately verified records across three source authorities — a subset, not a national data index.
- **Not a harvester.** No crawling until robots, licences, rate limits and publisher expectations have been reviewed.
- **Not governmental.** No publisher claim, no privileged correction workflow, no endorsement.

## Quickstart

Prerequisites: Node.js 22, pnpm 10.17.1, Go 1.23, Ruby 3.4 (for the governance validator).

```sh
git clone https://github.com/stanleyHayes/ghanadatasets.git
cd ghanadatasets
pnpm install
pnpm dev          # catalogue web app on http://localhost:3000
```

To run the API locally against the same fixture:

```sh
go build -trimpath -o bin/ghanadatasets-api ./cmd/api
./bin/ghanadatasets-api     # listens on :8080
```

| Environment variable | Default | Purpose |
|---|---|---|
| `PORT` | `8080` | API listen port |
| `DATASETS_DATA_PATH` | `data/datasets.json` | Metadata fixture to serve |
| `ALLOWED_ORIGIN` | `https://datasets.digitalghana.dev` | The single CORS origin permitted |

## Usage

The full REST surface is defined in [`contracts/openapi.yaml`](contracts/openapi.yaml); the GraphQL types in [`contracts/graphql/schema.graphql`](contracts/graphql/schema.graphql).

| Method | Path | Purpose |
|---|---|---|
| `GET` | `/health` | Service state, dataset version, coverage, record and review counts |
| `GET` | `/v1/datasets` | Search and filter — `q`, `publisher`, `topic`, `format` |
| `GET` | `/v1/datasets/{id}` | One record by stable ID; `404` with `DATASET_NOT_FOUND` otherwise |
| `POST` | `/graphql` | The single constrained `datasets(q, topic)` query |

### Record fields

| Field | Meaning |
|---|---|
| `id` | Stable registry identifier; survives title and URL changes |
| `title`, `publisher`, `topic`, `description` | Descriptive metadata |
| `licence` | Licence label exactly as it applies to this record |
| `accessCondition` | What you must accept or do before use |
| `sourceId` | Key into [`docs/governance/source-register.json`](docs/governance/source-register.json) |
| `checkedAt` | Date this record's access path was last checked |
| `status` | `REACHABLE` or `REVIEW`; `statusNote` explains a `REVIEW` |
| `distributions[]` | `{ format, accessUrl }` — every `accessUrl` is HTTPS and at the official host |

### Health

```sh
curl -s https://api-datasets.digitalghana.dev/health
```

```json
{"checkedAt":"2026-09-12T09:21:54Z","coverage":"metadata-only verified subset","datasetVersion":"2026.09.01-beta.1","datasets":10,"review":1,"status":"ok"}
```

### GraphQL

```sh
curl -s -X POST https://api-datasets.digitalghana.dev/graphql \
  -H 'Content-Type: application/json' \
  -d '{"query":"query($topic:String){datasets(topic:$topic){id title status}}","variables":{"topic":"Agriculture"}}'
```

Returns the two Agriculture records, one `REACHABLE` and one `REVIEW`, with `extensions.datasetVersion` echoing the fixture version. A request whose query does not reference `datasets` is rejected with `400`, and request bodies are capped at 32 KiB. The beta gate is a substring match rather than a GraphQL parser, so it constrains the surface but does not validate the document — tightening it is tracked as a good first contribution in `CONTRIBUTING.md`. In beta the resolver returns the complete record for every match — the selection set is accepted but not applied, so a response carries every field documented in the Record fields table above.

### TypeScript client

```ts
import { GhanaDatasetsClient } from "./sdk/typescript/index";

const client = new GhanaDatasetsClient();
const page = await client.search({ topic: "Population", format: "PXWEB_API" });
```

Zero dependencies, `fetch`-based, `AbortSignal` supported.

## Data and provenance

Every record traces to an authority reviewed in [`docs/governance/source-register.json`](docs/governance/source-register.json), whose shape is described by [`source-register.schema.json`](docs/governance/source-register.schema.json). The schema is declarative: `scripts/validate.rb` enforces only that the register is non-empty and that no source carries a `blocked` or `unknown` publication decision.

| Source ID | Authority | Publication decision | Reviewed |
|---|---|---|---|
| `gss-statsbank` | Ghana Statistical Service — StatsBank / PxWeb API | metadata-link-only | 2026-09-01 |
| `gss-microdata` | Ghana Statistical Service — Microdata Catalog | metadata-link-only | 2026-09-01 |
| `ghana-open-data` | Ghana Open Data Initiative / NITA | metadata-link-only | 2026-09-01 |
| `portfolio-brief` | Digital Ghana project owner | approved-for-requirements | 2026-09-01 |

- **Licence boundary.** Original code and configuration here are Apache-2.0. Referenced source documents and third-party datasets retain their own rights and are **not** relicensed by being indexed. See [`NOTICE`](NOTICE).
- **Versioning.** The fixture carries `datasetVersion` (currently `2026.09.01-beta.1`) and `reviewedAt`. Every successful API response echoes the version so a result can be pinned. Error bodies — the `DATASET_NOT_FOUND` 404 and the GraphQL `400` — carry only their error payload.
- **Verification dates.** Per-record `checkedAt` records when the access path was last checked. Release-time evidence is in [`docs/runbooks/release-evidence.md`](docs/runbooks/release-evidence.md).
- **Corrections.** Open an issue or pull request with the stable ID, the current value, the proposed value, the authoritative source URL and its publication date. Automation may draft a correction; a human reviewer approves publication. See [`CONTRIBUTING.md`](CONTRIBUTING.md).
- **Failed links stay visible.** Per [`docs/runbooks/operations.md`](docs/runbooks/operations.md), a broken official link is marked `REVIEW` with a note. It is never silently removed, and its provenance is never rewritten to imply availability.

## Project layout

```
app/           Next.js 16 catalogue UI (App Router, custom Radix topic picker, SEO/OG routes)
cmd/api/       Go HTTP service: REST, constrained GraphQL, CORS, security headers
internal/      catalog package — fixture parsing, invariants, search and filter logic
data/          datasets.json — the immutable metadata-only fixture
contracts/     OpenAPI 3.1 and GraphQL schema, the public interface of record
sdk/typescript Dependency-free client
docs/          product definition, ADR, governance source register, runbooks
scripts/       validate.rb — governance and required-file validator
tests/         Node fixture invariant tests
infra/         Vercel configuration for the web surface
```

## Verification

The single gate is `pnpm check`, which is exactly what [`.github/workflows/quality.yml`](.github/workflows/quality.yml) runs on every pull request and push to `main`:

```sh
pnpm install
pnpm check
```

It runs, in order:

```sh
ruby scripts/validate.rb   # required files, source-register decisions, ten-record invariant, secret scan
go test ./...              # catalogue invariants, search/filter, REST, GraphQL, CORS
go vet ./...
pnpm typecheck             # tsc --noEmit
pnpm test                  # node --test tests/*.test.mjs
pnpm build                 # production Next.js build
```

## Status and roadmap

**Public beta**, verified 2026-09-01. Ten metadata-only records, one visibly under link review. The web app and API are both live on their canonical hostnames with recorded rollback proof.

Directional plans are in [`ROADMAP.md`](ROADMAP.md). The authoritative, machine-readable task board is [`agent_plan.md`](agent_plan.md).

## Contributing and policies

- [`CONTRIBUTING.md`](CONTRIBUTING.md) — setup, verification, data corrections, review expectations
- [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md) — Contributor Covenant v2.1; reports to the private channel documented there
- [`SECURITY.md`](SECURITY.md) — report privately to the private channel documented there, never in a public issue
- [`AGENTS.md`](AGENTS.md) — coordination rules for automated contributors
- [`LICENSE`](LICENSE) — Apache License 2.0 · [`NOTICE`](NOTICE) — third-party rights

## Independence

GhanaDataset Registry is an independent open-source project. It is **not** operated by, affiliated with, or endorsed by the Government of Ghana, the Ghana Statistical Service, NITA, or any other public body. It indexes publicly available metadata and links to official sources. Authoritative data, access decisions and terms of use remain with those publishers.

Part of [Digital Ghana](https://digitalghana.dev).
