package config

import "os"

type Config struct {
	RedisURL         string
	RateLimitEnabled bool
	RateLimitPerMin  int
	RateLimitWindowS int
}

func Load() Config {
	rateLimitEnabled := parseBoolEnv("RATE_LIMIT_ENABLED")

	return Config{
		RedisURL:         os.Getenv("REDIS_URL"),
		RateLimitEnabled: rateLimitEnabled,
		RateLimitPerMin:  10,
		RateLimitWindowS: 60,
	}
}

func parseBoolEnv(key string) bool {
	value := os.Getenv(key)
	switch value {
	case "1", "true", "TRUE", "yes", "YES":
		return true
	default:
		return false
	}
}
