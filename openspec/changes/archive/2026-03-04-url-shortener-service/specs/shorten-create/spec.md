## ADDED Requirements

### Requirement: Shorten endpoint accepts valid URLs
The service SHALL accept a JSON payload with a `url` field containing a valid HTTP or HTTPS URL and return a JSON response with a `short_url` field.

#### Scenario: Successful shorten
- **WHEN** the client POSTs to `/shorten` with `{ "url": "https://example.com/path" }`
- **THEN** the service responds `200` with JSON containing a non-empty `short_url`

#### Scenario: Invalid URL rejected
- **WHEN** the client POSTs to `/shorten` with `{ "url": "ftp://example.com" }`
- **THEN** the service responds `400` with JSON `{ "error": "<message>" }`

#### Scenario: Missing url field rejected
- **WHEN** the client POSTs to `/shorten` with `{}`
- **THEN** the service responds `400` with JSON `{ "error": "<message>" }`

### Requirement: Error responses are JSON
The service SHALL return JSON error bodies in the form `{ "error": "<message>" }` for request validation failures.

#### Scenario: Malformed JSON
- **WHEN** the client POSTs to `/shorten` with invalid JSON
- **THEN** the service responds `400` with JSON `{ "error": "<message>" }`
