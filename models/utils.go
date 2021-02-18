package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	// you can now parse b as thoroughly as you want
	str := string(b[:])
	//!FORMAT IS DD-MM-YYYY
	date, err := time.Parse(`"02-01-2006"`, str)
	if err != nil {
		panic(err)
	}
	*t = Time{date}
	return nil
}

func (t *Time) MarshalJSON() (b []byte, err error) {
	day, month, year := t.Date()
	data := fmt.Sprintf("%d-%d-%d", day, month, year)
	return []byte(data), nil
}

func (t *Time) Scan(b interface{}) (err error) {
	switch x := b.(type) {
	case time.Time:
		t.Time = x
	default:
		err = fmt.Errorf("unsupported scan type %T", b)
	}
	return
}

func (t Time) Value() (driver.Value, error) {
	// check if the date was not set..
	if t.Time.IsZero() {
		return nil, nil
	}
	return t.Time.Format("02-01-2006"), nil
}
