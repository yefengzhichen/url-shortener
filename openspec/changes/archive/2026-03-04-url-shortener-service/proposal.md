## Why

The service needs a minimal, reliable URL shortener backend with clear API behavior, storage fallback, and validation. Defining this now enables a straightforward Go implementation and test coverage from the start.

## What Changes

- Add HTTP endpoints to create short URLs and redirect requests.
- Implement Base62 short-code generation with uniqueness guarantees.
- Add storage abstraction with Redis when configured, in-memory otherwise.
- Enforce URL validation and consistent JSON error responses.
- Add optional IP-based rate limiting.
- Add unit and HTTP-level tests.

## Capabilities

### New Capabilities
- `shorten-create`: Create short URLs via HTTP API with validation and JSON errors.
- `shorten-resolve`: Resolve short codes and redirect with HTTP 301.
- `shorten-storage`: Store and retrieve mappings with Redis or in-memory fallback.
- `shorten-rate-limit`: Optional per-IP rate limiting for shorten requests.
- `shorten-testing`: Unit and HTTP tests for core behaviors.

### Modified Capabilities
- 

## Impact

- New HTTP handlers and routing for `/shorten` and `/{short}`.
- New storage dependency on Redis when `REDIS_URL` is set.
- New core logic for code generation, validation, and resolution.
- Added test suites for API and business logic.
