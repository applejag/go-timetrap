package timetrap

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var DefaultConfigPath string

func init() {
	if env, ok := os.LookupEnv("TIMETRAP_CONFIG_FILE"); ok {
		DefaultConfigPath = env
	} else if home, err := os.UserHomeDir(); err == nil {
		DefaultConfigPath = filepath.Join(home, ".timetrap.yml")
	}
}

type Config struct {
	DatabaseFile         string      `yaml:"database_file"`
	RoundInSeconds       int         `yaml:"round_in_seconds"`
	FormatterSearchPaths []string    `yaml:"formatter_search_paths"`
	DefaultFormatter     string      `yaml:"default_formatter"`
	AutoSheet            string      `yaml:"auto_sheet"`
	AutoSheetSearchPaths []string    `yaml:"auto_sheet_search_paths"`
	AutoCheckout         bool        `yaml:"auto_checkout"`
	RequireNote          bool        `yaml:"require_note"`
	NoteEditor           interface{} `yaml:"note_editor"`
	WeekStart            WeekStart   `yaml:"week_start"`
	DayLengthHours       int         `yaml:"day_length_hours"`
}

type WeekStart string

const (
	Monday    WeekStart = "Monday"
	Tuesday   WeekStart = "Tuesday"
	Wednesday WeekStart = "Wednesday"
	Thursday  WeekStart = "Thursday"
	Friday    WeekStart = "Friday"
	Saturday  WeekStart = "Saturday"
	Sunday    WeekStart = "Sunday"
)

func NewConfigLocal() (Config, error) {
	return NewConfigFile(DefaultConfigPath)
}

func NewConfigFile(path string) (Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	return NewConfigBytes(b)
}

func NewConfigBytes(b []byte) (Config, error) {
	var config Config
	if err := yaml.Unmarshal(b, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}
