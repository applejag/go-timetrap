package timetrap

import (
	"fmt"
	"time"
)

// CREATE TABLE `entries` (
// 	`id`    integer NOT NULL PRIMARY KEY AUTOINCREMENT,
// 	`note`  varchar(255),
// 	`start` timestamp,
// 	`end`   timestamp,
// 	`sheet` varchar(255)
// );
// CREATE TABLE `meta` (
// 	`id`    integer NOT NULL PRIMARY KEY AUTOINCREMENT,
// 	`key`   varchar(255),
// 	`value` varchar(255)
// );

type Entry struct {
	ID    int     `gorm:"primaryKey"`
	Note  *string `gorm:"size:255"`
	Start *Timestamp
	End   *Timestamp
	Sheet *string `gorm:"size:255"`
}

func (Entry) TableName() string {
	return "entries"
}

func (e Entry) String() string {
	var (
		start = "     "
		end   = "     "
		note  = "<null>"
	)
	if e.Start != nil {
		start = time.Time(*e.Start).Format("15:04")
	}
	if e.End != nil {
		end = time.Time(*e.End).Format("15:04")
	}
	if e.Note != nil {
		note = *e.Note
	}
	return fmt.Sprintf("(%d) %s - %s %s", e.ID, start, end, note)
}

type Meta struct {
	ID    int     `gorm:"primaryKey"`
	Key   *string `gorm:"size:255"`
	Value *string `gorm:"size:255"`
}

func (Meta) TableName() string {
	return "meta"
}

func (m Meta) String() string {
	var (
		key   = "<null>"
		value = "<null>"
	)
	if m.Key != nil {
		key = *m.Key
	}
	if m.Value != nil {
		value = *m.Value
	}
	return fmt.Sprintf("(%d) %s=%s", m.ID, key, value)
}
