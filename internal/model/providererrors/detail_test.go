package providererrors

import (
	"encoding/json"
	"testing"
)

func TestDetailMessage(t *testing.T) {
	tests := []struct {
		name string
		raw  json.RawMessage
		want string
	}{
		{name: "nil", raw: nil, want: ""},
		{name: "empty", raw: json.RawMessage(``), want: ""},
		{name: "string", raw: json.RawMessage(`"Invalid API Key"`), want: "Invalid API Key"},
		{name: "string with spaces", raw: json.RawMessage(`"  padded  "`), want: "padded"},
		{name: "array", raw: json.RawMessage(`[{"msg":"field required"},{"msg":"not a valid string"}]`), want: "field required; not a valid string"},
		{name: "array with empty messages", raw: json.RawMessage(`[{"loc":["body"]},{"msg":" "}]`), want: ""},
		{name: "null", raw: json.RawMessage(`null`), want: ""},
		{name: "number", raw: json.RawMessage(`42`), want: ""},
		{name: "object", raw: json.RawMessage(`{"msg":"nested"}`), want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetailMessage(tt.raw); got != tt.want {
				t.Fatalf("DetailMessage(%s) = %q, want %q", tt.raw, got, tt.want)
			}
		})
	}
}
