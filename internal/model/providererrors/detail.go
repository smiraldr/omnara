package providererrors

import (
	"encoding/json"
	"strings"
)

// DetailMessage extracts a human-readable message from a FastAPI-style
// "detail" error field. FastAPI renders some errors as a plain string
// ({"detail":"Invalid API Key"}) and validation failures (HTTP 422) as an
// array of objects ({"detail":[{"msg":"field required","loc":[...]}]}).
// It returns "" when the field is absent or carries no readable message.
func DetailMessage(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	var message string
	if err := json.Unmarshal(raw, &message); err == nil {
		return strings.TrimSpace(message)
	}
	var items []struct {
		Msg string `json:"msg"`
	}
	if err := json.Unmarshal(raw, &items); err != nil {
		return ""
	}
	parts := make([]string, 0, len(items))
	for _, item := range items {
		if msg := strings.TrimSpace(item.Msg); msg != "" {
			parts = append(parts, msg)
		}
	}
	return strings.Join(parts, "; ")
}
