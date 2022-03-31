package lyy

import (
	"database/sql/driver"
	"encoding/json"
)

type StringSlice []string

func (a *StringSlice) Scan(value interface{}) error {
	byte := value.([]byte)
	return json.Unmarshal(byte, a)
}

func (a StringSlice) Value() (driver.Value, error) {
	if a == nil {
		return "[]", nil
	}
	byte, err := json.Marshal(a)
	if err != nil {
		return "[]", err
	}
	return string(byte), nil
}
