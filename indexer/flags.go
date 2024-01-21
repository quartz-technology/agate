package indexer

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Flags(v *viper.Viper, f *pflag.FlagSet) {
	indexerBeaconAPIURLFlag(v, f)
	indexerRelayAPIURLsFlag(v, f)
	indexerDatabaseMigrationSourceURLFlag(v, f)
	indexerDatabaseConnectionURLFlag(v, f)
}

func bindKeyToFlagToEnv(v *viper.Viper, f *pflag.FlagSet, key, flag, env string) {
	err := v.BindPFlag(key, f.Lookup(flag))
	cobra.CheckErr(err)

	err = v.BindEnv(key, env)
	cobra.CheckErr(err)
}

const (
	beaconAPIURLViperKey string = "indexer.beacon-api-url"
	beaconAPIURLFlag     string = "beacon-api-url"
	beaconAPIURLEnv      string = "AGATE_INDEXER_BEACON_API_URL"
	beaconAPIURLDefault  string = "http://localhost:5051"
)

func indexerBeaconAPIURLFlag(v *viper.Viper, f *pflag.FlagSet) {
	f.String(
		beaconAPIURLFlag,
		beaconAPIURLDefault,
		"The API URL of the beacon node the indexer connects to when subscribing to new head events",
	)

	bindKeyToFlagToEnv(v, f, beaconAPIURLViperKey, beaconAPIURLFlag, beaconAPIURLEnv)
}

func getBeaconAPIURL(v *viper.Viper) string {
	return v.GetString(beaconAPIURLViperKey)
}

var (
	relayAPIURLsViperKey string   = "indexer.relay-api-urls"
	relayAPIURLsFlag     string   = "relay-api-urls"
	relayAPIURLsEnv      string   = "AGATE_INDEXER_RELAY_API_URLS"
	relayAPIURLsDefault  []string = []string{
		"https://boost-relay.flashbots.net",
		"https://bloxroute.max-profit.blxrbdn.com",
		"https://bloxroute.regulated.blxrbdn.com",
		"https://relay.ultrasound.money",
		"https://relay.edennetwork.io",
		"https://mainnet.aestus.live",
		"https://agnostic-relay.net",
	}
)

func indexerRelayAPIURLsFlag(v *viper.Viper, f *pflag.FlagSet) {
	f.StringArray(
		relayAPIURLsFlag,
		relayAPIURLsDefault,
		"The API URLs of the relays the indexer aggregates data from.",
	)

	bindKeyToFlagToEnv(v, f, relayAPIURLsViperKey, relayAPIURLsFlag, relayAPIURLsEnv)
}

func getRelayAPIURLs(v *viper.Viper) []string {
	return v.GetStringSlice(relayAPIURLsViperKey)
}

const (
	dbMigrationSourceURLViperKey string = "indexer.database-migration-source-url"
	dbMigrationSourceURLFlag     string = "database-migration-source-url"
	dbMigrationSourceURLEnv      string = "AGATE_INDEXER_DATABASE_MIGRATION_SOURCE_URL"
	dbMigrationSourceURLDefault  string = "file://db/migrations"
)

func indexerDatabaseMigrationSourceURLFlag(v *viper.Viper, f *pflag.FlagSet) {
	f.String(
		dbMigrationSourceURLFlag,
		dbMigrationSourceURLDefault,
		"",
	)

	bindKeyToFlagToEnv(v, f, dbMigrationSourceURLViperKey, dbMigrationSourceURLFlag,
		dbMigrationSourceURLEnv)
}

func getDatabaseMigrationSourceURL(v *viper.Viper) string {
	return v.GetString(dbMigrationSourceURLViperKey)
}

const (
	dbConnectionURLViperKey string = "indexer.database-connection-url"
	dbConnectionURLFlag     string = "database-connection-url"
	dbConnectionURLEnv      string = "AGATE_INDEXER_DATABASE_CONNECTION_URL"
	dbConnectionURLDefault  string = "postgres://postgres:password@localhost:5432/agatedb?sslmode=disable"
)

func indexerDatabaseConnectionURLFlag(v *viper.Viper, f *pflag.FlagSet) {
	f.String(
		dbConnectionURLFlag,
		dbConnectionURLDefault,
		"",
	)

	bindKeyToFlagToEnv(v, f, dbConnectionURLViperKey, dbConnectionURLFlag,
		dbConnectionURLEnv)
}

func getDatabaseConnectionURL(v *viper.Viper) string {
	return v.GetString(dbConnectionURLViperKey)
}
