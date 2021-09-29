package timetrap

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type Timestamp time.Time

var ErrInvalidTimestampType = errors.New("Timestamp.Scan expected time.Time")

// Scan fulfills database/sql.Scanner interface
// https://pkg.go.dev/database/sql#Scanner
func (lt *Timestamp) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("got %T: %w", value, ErrInvalidTimestampType)
	}
	*lt = (Timestamp)(time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local))
	return nil
}

// Value fulfills database/sql/driver.Valuer interface
// https://pkg.go.dev/database/sql/driver#Valuer
func (lt *Timestamp) Value() (driver.Value, error) {
	return (*time.Time)(lt), nil
}

// String fulfills fmt.Stringer interface
// https://pkg.go.dev/fmt#Stringer
func (lt Timestamp) String() string {
	return time.Time(lt).String()
}

// GormDataType fulfills gorm.io/gorm/schema.GormDataTypeInterface interface
// https://pkg.go.dev/gorm.io/gorm@v1.21.15/schema#GormDataTypeInterface
func (Timestamp) GormDataType() string {
	return "timestamp"
}
