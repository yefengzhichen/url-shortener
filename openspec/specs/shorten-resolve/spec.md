## ADDED Requirements

### Requirement: Resolve short code redirects
The service SHALL redirect `GET /{short}` requests to the original URL with HTTP status `301` when the short code exists.

#### Scenario: Successful redirect
- **WHEN** the client GETs `/abc123` for an existing short code
- **THEN** the service responds `301` with `Location` set to the original URL

### Requirement: Unknown short code returns JSON error
The service SHALL return a JSON error body when a short code does not exist.

#### Scenario: Short code not found
- **WHEN** the client GETs `/unknown` for a non-existent short code
- **THEN** the service responds `404` with JSON `{ "error": "<message>" }`
