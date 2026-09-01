# ADR-0001: Independent product boundary

Status: Accepted

## Decision

GhanaDataset Registry owns its repository, release lifecycle, data, credentials, contracts and operational evidence. It may consume another Digital Ghana product only through a versioned public contract or pinned dataset artifact.

The canonical web hostname is `datasets.digitalghana.dev`. Additional API or operational hostnames require an explicit ADR and portfolio registry entry before deployment.

The canonical API hostname is `api-datasets.digitalghana.dev`, recorded centrally before deployment. The first release indexes descriptive metadata and official access paths only. It does not mirror resources, bypass study access conditions, normalize licences into a broader permission, or imply publisher endorsement.

## Consequences

Failures remain isolated, histories remain understandable, and a portfolio-wide platform outage is not created by convenience. Some configuration and small primitives may be repeated until two proven consumers justify a versioned shared package.
