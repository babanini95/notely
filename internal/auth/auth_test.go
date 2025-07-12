package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		inputAuth      string
		expectedOutput string
	}{
		"should pass":         {inputAuth: "ApiKey theapikey", expectedOutput: "theapikey"},
		"Authorization empty": {inputAuth: "", expectedOutput: ""},
		"wrong auth header":   {inputAuth: "JWT thejwt", expectedOutput: ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			header := make(http.Header)
			header.Set("Authorization", tc.inputAuth)

			apiKey, _ := GetAPIKey(header)
			if !reflect.DeepEqual(apiKey, tc.expectedOutput) {
				t.Fatalf("expected: %v, got: %v", tc.expectedOutput, apiKey)
			}
		})
	}
}
