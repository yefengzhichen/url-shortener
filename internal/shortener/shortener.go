package shortener

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strings"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenerateCode() (string, error) {
	length, err := randomLength(6, 8)
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	builder.Grow(length)
	for i := 0; i < length; i++ {
		idx, err := randomIndex(len(alphabet))
		if err != nil {
			return "", err
		}
		builder.WriteByte(alphabet[idx])
	}

	return builder.String(), nil
}

func ValidateURL(raw string) error {
	parsed, err := url.Parse(raw)
	if err != nil {
		return errors.New("invalid url")
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return errors.New("url must use http or https")
	}
	if parsed.Host == "" {
		return errors.New("url must include a host")
	}
	return nil
}

func BuildShortURL(r *http.Request, code string) string {
	scheme := requestScheme(r)
	host := r.Host
	if host == "" {
		return "/" + code
	}
	return fmt.Sprintf("%s://%s/%s", scheme, host, code)
}

func requestScheme(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-Proto")
	if forwarded != "" {
		parts := strings.Split(forwarded, ",")
		return strings.TrimSpace(parts[0])
	}
	if r.TLS != nil {
		return "https"
	}
	return "http"
}

func randomLength(min int, max int) (int, error) {
	if max <= min {
		return min, nil
	}
	delta := max - min + 1
	idx, err := randomIndex(delta)
	if err != nil {
		return 0, err
	}
	return min + idx, nil
}

func randomIndex(max int) (int, error) {
	if max <= 0 {
		return 0, errors.New("invalid max")
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}
