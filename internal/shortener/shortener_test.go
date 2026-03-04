package shortener

import (
	"net/http"
	"net/url"
	"regexp"
	"testing"
)

func TestGenerateCodeUniqueness(t *testing.T) {
	seen := make(map[string]struct{})
	pattern := regexp.MustCompile("^[0-9A-Za-z]{6,8}$")

	for i := 0; i < 200; i++ {
		code, err := GenerateCode()
		if err != nil {
			t.Fatalf("GenerateCode failed: %v", err)
		}
		if !pattern.MatchString(code) {
			t.Fatalf("unexpected code format: %s", code)
		}
		if _, ok := seen[code]; ok {
			t.Fatalf("duplicate code generated: %s", code)
		}
		seen[code] = struct{}{}
	}
}

func TestValidateURL(t *testing.T) {
	tests := []struct {
		value string
		ok    bool
	}{
		{value: "https://example.com", ok: true},
		{value: "http://example.com/path", ok: true},
		{value: "ftp://example.com", ok: false},
		{value: "example.com", ok: false},
	}

	for _, test := range tests {
		err := ValidateURL(test.value)
		if test.ok && err != nil {
			t.Fatalf("expected valid for %s, got error %v", test.value, err)
		}
		if !test.ok && err == nil {
			t.Fatalf("expected error for %s", test.value)
		}
	}
}

func TestBuildShortURL(t *testing.T) {
	req := &http.Request{Host: "example.com", URL: &url.URL{}}
	short := BuildShortURL(req, "abc123")
	if short != "http://example.com/abc123" {
		t.Fatalf("unexpected short url: %s", short)
	}
}
