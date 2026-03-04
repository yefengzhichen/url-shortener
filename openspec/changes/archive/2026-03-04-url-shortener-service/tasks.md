## 1. Setup & Dependencies

- [x] 1.1 Initialize Go module structure if missing and add Gin dependency
- [x] 1.2 Add Redis client dependency and configuration via `REDIS_URL`

## 2. Core Logic

- [x] 2.1 Implement Base62 short-code generator with 6-8 char output
- [x] 2.2 Implement URL validation for HTTP/HTTPS using `net/url`
- [x] 2.3 Define storage interface and in-memory `sync.Map` implementation
- [x] 2.4 Implement Redis storage with Get/Set and collision checks

## 3. HTTP Handlers

- [x] 3.1 Add `POST /shorten` handler with validation and JSON errors
- [x] 3.2 Add `GET /{short}` handler with 301 redirect and 404 JSON error
- [x] 3.3 Wire Gin router and middleware for JSON errors

## 4. Rate Limiting

- [x] 4.1 Add optional per-IP rate limiter (10/min) for shorten endpoint
- [x] 4.2 Make rate limiting configurable (enabled/disabled)

## 5. Tests

- [x] 5.1 Unit tests for code generation uniqueness and Base62 format
- [x] 5.2 Unit tests for URL validation and storage get/set behavior
- [x] 5.3 HTTP tests for shorten and resolve endpoints

## 6. Docker

- [x] 6.1 Add Dockerfile for service build and run
- [x] 6.2 Add docker-compose for local Redis and service
