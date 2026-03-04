# URL Shortener Service

Simple URL shortener backend built with Go and Gin. Supports Redis storage when `REDIS_URL` is set, otherwise falls back to in-memory storage.

## Features

- `POST /shorten` returns a short URL
- `GET /{short}` redirects to the original URL (301)
- JSON error responses: `{ "error": "message" }`
- Optional rate limiting (10 requests per minute per IP)

## Configuration

- `REDIS_URL`: Redis connection string (e.g. `redis://localhost:6379/0`). If empty, uses in-memory storage.
- `RATE_LIMIT_ENABLED`: Set to `true` to enable rate limiting on `POST /shorten`.
- `PORT`: HTTP port (default `8080`).

## Run Locally

```bash
go run ./cmd/server
```

## Usage

Create a short URL:

```bash
curl -X POST "http://localhost:8080/shorten" \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com/path"}'
```

Example response:

```json
{"short_url":"http://localhost:8080/Abc123"}
```

Resolve the short URL:

```bash
curl -I "http://localhost:8080/Abc123"
```

```
curl -X POST "http://localhost:8080/shorten" \
  -H "Content-Type: application/json" \
  -d '{"url":"https://kubernetes.io/docs/concepts/overview/"}'
{"short_url":"http://localhost:8080/kaQax2J"}%

curl -v "http://localhost:8080/kaQax2J"
> GET /kaQax2J HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.88.1
> Accept: */*
>
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: https://kubernetes.io/docs/concepts/overview/
< Date: Wed, 04 Mar 2026 06:32:42 GMT
< Content-Length: 80
<
<a href="https://kubernetes.io/docs/concepts/overview/">Moved Permanently</a>.

* Connection #0 to host localhost left intact
```


## Tests

Run all tests:

```bash
go test ./...
```
