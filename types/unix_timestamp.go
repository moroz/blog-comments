package types

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type UnixTimestamp struct {
	time.Time
}

var (
	_ driver.Valuer = &UnixTimestamp{}
	_ sql.Scanner   = &UnixTimestamp{}
)

func (u *UnixTimestamp) Value() (driver.Value, error) {
	return u.Unix(), nil
}

func (u *UnixTimestamp) Scan(src any) error {
	switch s := src.(type) {
	case int64:
		u.Time = time.Unix(s, 0).UTC()

	case string:
		return u.parse(s)

	case []byte:
		return u.parse(string(s))

	default:
		return fmt.Errorf("cannot scan %T into UnixTimestamp", src)
	}

	return nil
}

func (u *UnixTimestamp) parse(src string) error {
	parsed, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		return fmt.Errorf("UnixTimestamp.Scan: %w", err)
	}
	u.Time = time.Unix(parsed, 0).UTC()
	return nil
}
