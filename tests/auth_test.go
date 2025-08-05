package tests

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header2 := http.Header{}
	header2.Set("Authorization", "ApiKey ABCDEFGH")

	header1 := http.Header{}
	header1.Set("Authorization", "ABCDEFGH")

	tests := map[string]struct {
		input     http.Header
		wantKey   string
		wantError bool
	}{
		"simple": {
			input:     header1,
			wantKey:   "ABCDEFGH",
			wantError: false,
		},
		"malformed": {
			input:     header2,
			wantKey:   "",
			wantError: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			apiKey, err := auth.GetAPIKey(tc.input)
			if (err != nil) != tc.wantError {
				t.Errorf("GetAPIKey() error = %v, wantError = %v", err, tc.wantError)
			}

			if apiKey != tc.wantKey {
				t.Errorf("GetAPIKey() apiKey = %v, wantKey = %v", apiKey, tc.wantKey)
			}
		})
	}
}
