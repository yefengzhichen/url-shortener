## ADDED Requirements

### Requirement: Unit tests cover core logic
The service SHALL include unit tests for short-code generation, URL validation, and storage get/set behavior.

#### Scenario: Code generation uniqueness
- **WHEN** multiple codes are generated in a test
- **THEN** each code is unique and Base62 encoded

#### Scenario: URL validation
- **WHEN** a valid HTTP URL is validated
- **THEN** validation succeeds

#### Scenario: URL validation fails
- **WHEN** a non-HTTP(S) URL is validated
- **THEN** validation fails

### Requirement: HTTP tests cover API behavior
The service SHALL include HTTP tests for `/shorten` and `/{short}` endpoints.

#### Scenario: Shorten success
- **WHEN** a valid shorten request is posted
- **THEN** the response includes `short_url`

#### Scenario: Resolve success
- **WHEN** a stored short code is requested
- **THEN** the response is a `301` redirect
