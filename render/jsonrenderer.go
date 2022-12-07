package render

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONRenderer struct {
	indent bool
}

func NewJSONRenderer(indent bool) *JSONRenderer {
	return &JSONRenderer{indent}
}

func (jr *JSONRenderer) Render(w http.ResponseWriter, v interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	if jr.indent {
		enc.SetIndent("", "  ")
	}
	if err := enc.Encode(v); err != nil {
		return fmt.Errorf("")
	}
	return nil
}
