## ADDED Requirements

### Requirement: Optional rate limiting for shorten
The service SHALL apply a per-IP rate limit of 10 requests per minute to `POST /shorten` when rate limiting is enabled.

#### Scenario: Rate limit exceeded
- **WHEN** a client exceeds 10 POST requests within a minute from the same IP
- **THEN** the service responds `429` with JSON `{ "error": "<message>" }`

#### Scenario: Rate limit disabled
- **WHEN** rate limiting is disabled
- **THEN** the service accepts requests without rate limiting enforcement
