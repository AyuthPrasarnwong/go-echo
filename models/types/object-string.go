package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type ObjectStr map[string]string

//Value json Marshal to byte
func (a ObjectStr) Value() (driver.Value, error) {
	bytes, err := json.Marshal(a)
	return string(bytes), err
}

//Scan string or byte Unmarshal to json
func (a *ObjectStr) Scan(src interface{}) error {
	switch value := src.(type) {
	case string:
		return json.Unmarshal([]byte(value), a)
	case []byte:
		return json.Unmarshal(value, a)
	default:
		return errors.New("not supported")
	}
}
