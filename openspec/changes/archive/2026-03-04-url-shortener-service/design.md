## Context

The service is a Go-based URL shortener with a minimal HTTP API. It must validate input URLs, generate unique Base62 short codes, and redirect requests. Storage should use Redis when configured via `REDIS_URL`, otherwise fall back to in-memory storage. Optional per-IP rate limiting is needed for shorten requests.

## Goals / Non-Goals

**Goals:**
- Provide `POST /shorten` and `GET /{short}` endpoints with consistent JSON error responses.
- Use a storage abstraction that supports Redis and in-memory `sync.Map`.
- Ensure short-code uniqueness with 6-8 character Base62 values.
- Validate only HTTP/HTTPS URLs.
- Add unit and HTTP-level tests for core behaviors.

**Non-Goals:**
- User authentication, billing, analytics, or custom domains.
- Persistent storage beyond Redis and in-memory.
- Advanced rate limiting or distributed throttling.

## Decisions

- **Use Gin for HTTP routing.** It is lightweight, common in Go services, and aligns with project direction.
  - Alternatives: net/http directly (more boilerplate), Echo (similar but not required).
- **Storage interface with two implementations.** Define a small interface (Get/Set) to allow Redis or in-memory `sync.Map`.
  - Alternatives: global map without abstraction (harder to test and swap).
- **Base62 short code generation with collision retry.** Generate 6-8 chars and retry on collision by checking storage.
  - Alternatives: hash input URL (can leak predictability, harder to ensure uniqueness).
- **Validation via net/url.** Enforce scheme is http/https and host is present.
  - Alternatives: regex (error-prone), third-party validator (extra dependency).
- **Optional rate limiting per IP on POST.** Keep it simple with an in-memory token bucket when enabled.
  - Alternatives: Redis-backed rate limit (more complex, not required for v1).

## Risks / Trade-offs

- In-memory storage and rate limiting are not shared across instances → acceptable for local/dev and fallback; Redis path mitigates for production.
- Short code collisions possible with small space → mitigate via retry and allow 8 chars.
- Optional rate limiting can be bypassed in distributed setups → acceptable for basic protection.
