# GhanaDataset Registry product definition

## Problem and users

Official Ghana data exists across catalogues with different search, licence and machine-access behavior. GhanaDataset Registry gives developers, journalists, researchers and civic-technology teams one provenance-first index of where datasets can be found and how access is governed.

## Public beta scope

- Ten source-linked discovery records across GSS StatsBank, GSS Microdata Catalog and Ghana Open Data.
- Stable registry IDs, publishers, topics, descriptions, source dates, licence labels and official access paths.
- REST, constrained GraphQL and a dependency-free TypeScript client.
- Search and publisher/topic/access-method filters.
- Checked-at status that describes link reachability, not dataset correctness or freshness.

## Non-goals and safety

- No copying or proxying source datasets, tables or microdata.
- No bypassing login, application, research-purpose or attribution conditions.
- No assumption that portal-wide policy overrides a record-specific exception.
- No claim that a reachable URL makes underlying data current, complete or accurate.
- No harvesting until robots, licences, rate limits and publisher expectations are reviewed.
- No publisher claim or privileged correction workflow in beta.

## Acceptance

- Every record exposes source, publisher, licence/access condition, checked date and at least one official distribution.
- IDs remain stable when titles or source URLs change.
- Unknown licence is visible and never treated as open.
- Search finds titles/topics and filters never invent matches.
- Registry stores metadata only; fixtures contain no downloaded source rows.
