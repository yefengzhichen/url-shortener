## ADDED Requirements

### Requirement: Storage selection via environment
The service SHALL use Redis storage when `REDIS_URL` is set; otherwise it SHALL use in-memory storage.

#### Scenario: Redis configured
- **WHEN** `REDIS_URL` is set
- **THEN** the service stores and resolves short codes using Redis

#### Scenario: Redis not configured
- **WHEN** `REDIS_URL` is not set
- **THEN** the service stores and resolves short codes using in-memory storage

### Requirement: Storage guarantees uniqueness
The service SHALL prevent short-code collisions by ensuring stored codes are unique.

#### Scenario: Collision retry
- **WHEN** a generated short code already exists in storage
- **THEN** the service generates a new code and retries until unique
