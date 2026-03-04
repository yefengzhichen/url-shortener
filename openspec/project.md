# URL Shortener Server

## Overview
This project provides a URL shortener backend service. It accepts long URLs, generates short codes, and redirects clients from short URLs to their original destinations.

## Goals
- Simple HTTP API for creating short URLs and resolving them.
- Reliable, low-latency redirects.
- Safe handling of invalid or expired short codes.

## Core Behavior
- Create a short code for a given URL.
- Resolve a short code to its original URL.
- Redirect clients to the original URL using standard HTTP redirects.

## Data Model (High Level)
- Short code
- Original URL
- Creation timestamp
- Optional expiration timestamp
- Optional metadata (e.g., creator, tags)

## API Surface (High Level)
- `POST /shorten` to create a short URL.
- `GET /{code}` to resolve and redirect.
- `GET /healthz` for health checks.

## Storage
- Persistent store for code-to-URL mappings.
- Unique index on short codes.

## Non-Goals
- User authentication and billing.
- Analytics dashboards.
- Custom domains (unless added later).

## Engineering Conventions
- Keep handlers small and focused.
- Validate inputs and normalize URLs.
- Prefer explicit error responses and status codes.
- Implementation uses Go with the Gin framework.
- Use in-memory storage for code-to-URL mappings.
- Add unit tests for important functions.
