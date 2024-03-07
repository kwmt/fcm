package fcm

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestMessageSetData(t *testing.T) {
	data := make(map[string]interface{})
	data["url"] = "https://example.com"

	message := NewMessage("regsitration token sample").
		SetData(data)

	out, err := json.MarshalIndent(message, "", "  ")
	if err != nil {
		t.Error(err)
	}
	expected := `{
  "message": {
    "topic": "regsitration token sample",
    "notification": {},
    "data": {
      "url": "https://example.com"
    }
  }
}`

	if string(out) != expected {
		t.Errorf("%s\n%s", strings.TrimSpace(string(out)), strings.TrimSpace(expected))
	}
}
