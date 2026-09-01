# Operations baseline

## Production inventory

- Web: `https://datasets.digitalghana.dev` on Vercel project `ghanadatasets`.
- API: `https://api-datasets.digitalghana.dev` on Render service `srv-dabetjcs728c739m9ti0`.
- Health: `GET /health` returns service state, dataset version, coverage and record/review counts.
- Source: immutable metadata fixture `data/datasets.json`; the beta has no writable production state and therefore no backup surface.
- Owner: Digital Ghana maintainers via the public repository.

## Release and incident procedure

1. Run `pnpm check` and require the GitHub Quality workflow to pass.
2. Confirm every source record has an explicit publication decision and every metadata record exposes licence/access conditions.
3. Deploy the web from `main` to Vercel and the API from the same commit to Render.
4. Smoke canonical TLS, `/health`, REST, GraphQL, allowed-origin CORS and denied-origin CORS.
5. Check favicon, manifest, canonical metadata, robots, sitemap and the 1200x630 OG image.
6. If web health regresses, use `vercel rollback <known-good-deployment> --yes`.
7. If API health regresses, use Render's `POST /v1/services/{serviceId}/rollback` with a known-good `deployId`.
8. Keep failed official links visible as `REVIEW`; never silently remove the provenance or imply source availability.

Requests are stateless and logs must not include credentials or source payloads. The API allows only the canonical web origin and caps GraphQL request bodies at 32 KiB.
