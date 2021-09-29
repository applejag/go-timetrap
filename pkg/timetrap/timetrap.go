package timetrap

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound
var ErrNullMetaValue = errors.New("meta value was null")

type DB interface {
	GetCurrentSheet() (string, error)
	GetLastCheckoutID() (int, error)
	GetActiveEntry() (Entry, error)
	GetEntriesTimeRange(start, end time.Time) ([]Entry, error)
}

func NewDB(dbPath string) (DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &gormDB{DB: db}, nil
}

type gormDB struct {
	*gorm.DB
}

func (db *gormDB) getMeta(key string) (string, error) {
	var dbMeta Meta
	if err := db.Where(&Meta{Key: &key}).
		First(&dbMeta).Error; err != nil {
		return "", fmt.Errorf("meta key '%s': %w", key, err)
	}
	if dbMeta.Value == nil {
		return "", fmt.Errorf("meta key '%s': %w", key, ErrNullMetaValue)
	}
	return *dbMeta.Value, nil
}

func (db *gormDB) GetCurrentSheet() (string, error) {
	return db.getMeta("current_sheet")
}

func (db *gormDB) GetLastCheckoutID() (int, error) {
	value, err := db.getMeta("last_checkout_id")
	if err != nil {
		return 0, err
	}
	i64, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("meta key 'last_checkout_id': %w", err)
	}
	return int(i64), nil
}

func (db *gormDB) GetActiveEntry() (Entry, error) {
	var dbEntry Entry
	if err := db.Where(&Entry{End: nil}, "End").
		Order("id desc").
		First(&dbEntry).Error; err != nil {
		return Entry{}, fmt.Errorf("get active entry: %w", err)
	}
	return dbEntry, nil
}

func (db *gormDB) GetEntriesTimeRange(start, end time.Time) ([]Entry, error) {
	var dbEntries []Entry
	if err := db.
		Where("start BETWEEN ? AND ?", start, end).
		Order("start asc").
		Find(&dbEntries).Error; err != nil {
		return nil, fmt.Errorf("get entries: '%s' to '%s': %w",
			start.Format(time.RFC3339),
			end.Format(time.RFC3339),
			err)
	}
	return dbEntries, nil
}
