package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

const (
	// TimeFormat 时间格式化
	TimeFormat = "2006-01-02 15:04:05"
)

type Timestamp time.Time

func (this Timestamp) Value() (driver.Value, error) {
	var zero time.Time
	t := time.Time(this)
	// 判断给定时间是否和默认零时间的时间戳相同
	if t.UnixNano() == zero.UnixNano() {
		return nil, nil
	}
	return t, nil
}

func (this *Timestamp) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*this = Timestamp(value)
		return nil
	}
	return nil
}

func (this *Timestamp) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*this)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(TimeFormat))), nil
}

func (this *Timestamp) Format() string {
	tTime := time.Time(*this)
	return tTime.Format(TimeFormat)
}

type JSON map[string]interface{}

// Scan Scanner
func (this *JSON) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}

	return json.Unmarshal(b, &this)
}

// Value Valuer
func (this JSON) Value() (driver.Value, error) {
	if &this == nil {
		return nil, nil
	}
	return json.Marshal(this)
}

func ToScanJSON(src interface{}) JSON {
	jsonByte, err := json.Marshal(src)
	if err != nil {
		return nil
	}
	var scanJSON JSON
	if err = json.Unmarshal(jsonByte, &scanJSON); err != nil {
		return nil
	}
	return scanJSON
}

func JSONToANY[T any](src interface{}) T {
	var dest T
	if src == nil {
		return dest
	}
	jsonByte, err := json.Marshal(src)
	if err != nil {
		return dest
	}
	if err = json.Unmarshal(jsonByte, &dest); err != nil {
		return dest
	}
	return dest
}
