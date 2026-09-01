# Release evidence — public beta — 2026-09-01

## Immutable release

- Source commit: `dfc3e6d43932a51a94aa2af1e3cd20ef7908c94c`.
- GitHub Quality run: `33525713204` — success.
- Dataset version: `2026.09.01-beta.1`; 10 metadata-only records, 1 visibly under link review.
- Vercel project: `ghanadatasets`; current restored deployment `dpl_7DNotCeQfjj422UESt9TL5ecUVHi`.
- Render service: `srv-dabetjcs728c739m9ti0`; current restore deploy `dep-dabeveu10ojc73a4cjj0`.
- Render custom domain: `cdm-dabetl710e5c73ad1sag`; DNS record `rec_f32e22172d0ecb73726cd341`; verified.

## Verification

- `pnpm check`: governance validation, Go tests/vet, TypeScript, fixture tests and production build passed.
- `https://datasets.digitalghana.dev`: TLS 200, canonical metadata, JSON-LD, favicon, manifest, robots and sitemap passed.
- OG response is PNG at exactly 1200x630; Open Graph and Twitter large-image tags resolve to the canonical host.
- Production browser: search `trade` returned one record; custom Radix `Population` filter returned three records.
- Production UI: Outfit body, Geist Mono labels, Newsreader accent title; zero native select/dialog/checkbox/radio/date/time controls and zero horizontal overflow at 1489px.
- `https://api-datasets.digitalghana.dev/health`: 200 with version `2026.09.01-beta.1`, 10 records and 1 review state.
- REST search, constrained GraphQL, canonical-origin CORS and denied untrusted-origin CORS passed.
- Known limitation: the Ghana Open Data agriculture collection timed out during source preflight and remains explicitly labelled `REVIEW`; this does not affect the other nine checked access paths.

## Rollback proof

- Vercel: deployed `dpl_7DNotCeQfjj422UESt9TL5ecUVHi`, rolled back to `dpl_GxF4GojVWoBo7RXb4kSqLvCWGN7k`, smoked the canonical host, then restored `dpl_7DNotCeQfjj422UESt9TL5ecUVHi`; all steps passed.
- Render: deployed `dep-dabeus4s728c739mebs0`, rolled back to `dep-dabetk4s728c739m9vdg` producing `dep-dabev75cqm1c73dcofpg`, smoked `/health`, then restored the newer release producing `dep-dabeveu10ojc73a4cjj0`; all steps passed.
