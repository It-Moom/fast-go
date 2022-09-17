/*
 * @PackageName: handler
 * @Description:
 * @Author: gabbymrh
 * @Date: 2022-09-17 17:29:25
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-09-17 17:29:25
 */

package handler

import (
	"database/sql/driver"
	"time"
)

// NullTime 可为空的时间类型
type NullTime struct {
	Time  time.Time
	Valid bool
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

// ToNullTime 转换为可为空的时间类型
func ToNullTime(t time.Time) NullTime {
	return NullTime{Time: t, Valid: !t.IsZero()}
}
