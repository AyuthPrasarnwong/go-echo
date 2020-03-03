package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Array []interface{}

//Value json Marshal to byte
func (a Array) Value() (driver.Value, error) {
	bytes, err := json.Marshal(a)
	return string(bytes), err
}

//Scan string or byte Unmarshal to json
func (a *Array) Scan(src interface{}) error {
	switch value := src.(type) {
	case string:
		return json.Unmarshal([]byte(value), a)
	case []byte:
		return json.Unmarshal(value, a)
	default:
		return errors.New("not supported")
	}
}
