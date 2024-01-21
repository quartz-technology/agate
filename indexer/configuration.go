package indexer

import "github.com/spf13/viper"

type Configuration struct {
	BeaconAPIURL               string
	RelayAPIURLs               []string
	DatabaseMigrationSourceURL string
	DatabaseConnURL            string
}

func NewConfigurationFromViper(v *viper.Viper) *Configuration {
	return &Configuration{
		BeaconAPIURL:               getBeaconAPIURL(v),
		RelayAPIURLs:               getRelayAPIURLs(v),
		DatabaseMigrationSourceURL: getDatabaseMigrationSourceURL(v),
		DatabaseConnURL:            getDatabaseConnectionURL(v),
	}
}
