package provider

import "time"

// Provider represents time data provider. Allows fetch of time data, and other
// provider-related functions.
type Provider interface {
	FetchAll(start time.Time, end time.Time) []TimeEntry
}

// TimeEntry represents one time entry, which can be ie. duplicated multiple times for few
// days of doing one task.
type TimeEntry struct {
	Billable     bool
	Description  string
	Id           string
	IsLocked     bool
	Tags         []string
	TaskId       string
	TimeInterval timeInterval
}

// timeInterval represents time interval.
type timeInterval struct {
	Duration time.Duration
	End      time.Time
	Start    time.Time
}

// Config represents base provider config.
type Config struct {
	ProviderName string `yaml:"provider_name"`
	ApiKey       string `yaml:"api_key"`
	AuthMethod   string `yaml:"auth_method"`
	FetchAllUri  string `yaml:"fetch_all_uri"`
}

func NewConfig(providerName string, apiKey string, authMethod string, fetchAllUri string) *Config {
	return &Config{
		ProviderName: providerName,
		ApiKey:       apiKey,
		AuthMethod:   authMethod,
		FetchAllUri:  fetchAllUri,
	}
}
