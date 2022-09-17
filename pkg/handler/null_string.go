/*
 * @PackageName: handler
 * @Description: 空字符串处理器
 * @Author: gabbymrh
 * @Date: 2022-09-17 17:23:16
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-09-17 17:23:16
 */

package handler

import (
	"database/sql"
	"encoding/json"
)

// NullString is a wrapper around sql.NullString that implements the json.Marshaler and json.Unmarshaler interfaces.
type NullString struct {
	sql.NullString
}

// MarshalJSON implements the json.Marshaler interface.
func (v *NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v NullString) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.String = *s
	} else {
		v.Valid = false
	}
	return nil
}
