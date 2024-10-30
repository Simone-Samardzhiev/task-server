package task

import (
	"database/sql"
	"encoding/json"
)

type NullTime struct {
	sql.NullTime
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	if t.Valid {
		return json.Marshal(t.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (t *NullTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t.Valid = false
		return nil
	}

	t.Valid = true
	return json.Unmarshal(data, &t.Time)
}

func (t *NullTime) String() string {
	if t.Valid {
		return t.Time.String()
	} else {
		return "null"
	}
}
