FROM golang:1.25.0 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener ./cmd/server

FROM gcr.io/distroless/base-debian12

WORKDIR /
COPY --from=build /url-shortener /url-shortener

EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/url-shortener"]
