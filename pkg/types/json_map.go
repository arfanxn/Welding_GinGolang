package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONMap map[string]any

// ---------------------------
// GORM SCANNER + VALUER
// ---------------------------
func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}

	if len(m) == 0 {
		return []byte("{}"), nil
	}

	return json.Marshal(m)
}

func (m *JSONMap) Scan(value any) error {
	if value == nil {
		*m = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("jsonmap: %T", value)
	}

	if len(bytes) == 0 {
		*m = JSONMap{}
		return nil
	}

	var tmp map[string]any
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	*m = tmp
	return nil
}

// ---------------------------
// SAFE GETTERS
// ---------------------------

func (m JSONMap) GetString(key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

func (m JSONMap) GetBool(key string) bool {
	if v, ok := m[key].(bool); ok {
		return v
	}
	return false
}

func (m JSONMap) GetInt(key string) int {
	switch v := m[key].(type) {
	case int:
		return v
	case float64:
		return int(v) // JSON numbers come as float64
	default:
		return 0
	}
}

func (m JSONMap) GetFloat(key string) float64 {
	if v, ok := m[key].(float64); ok {
		return v
	}
	return 0
}

func (m JSONMap) GetMap(key string) JSONMap {
	if v, ok := m[key].(map[string]any); ok {
		return JSONMap(v)
	}
	if v, ok := m[key].(JSONMap); ok {
		return v
	}
	return JSONMap{}
}

func (m JSONMap) GetList(key string) []any {
	if v, ok := m[key].([]any); ok {
		return v
	}
	return []any{}
}

// ---------------------------
// SAFE SETTERS (OPTIONAL)
// ---------------------------

func (m JSONMap) Set(key string, value any) {
	m[key] = value
}

func (m JSONMap) SetString(key, value string) {
	m[key] = value
}

func (m JSONMap) SetBool(key string, value bool) {
	m[key] = value
}

func (m JSONMap) SetInt(key string, value int) {
	m[key] = value
}

func (m JSONMap) SetFloat(key string, value float64) {
	m[key] = value
}
