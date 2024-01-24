package cmd

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"github.com/quartz-technology/agate/indexer"
	"github.com/quartz-technology/agate/indexer/data_aggregator"
	"github.com/quartz-technology/agate/indexer/data_preprocessor"
	"github.com/quartz-technology/agate/indexer/head_listener"
	"github.com/quartz-technology/agate/indexer/storage_manager"
	"github.com/quartz-technology/agate/indexer/storage_manager/store/dto"
	"github.com/quartz-technology/agate/indexer/storage_manager/store/postgres"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// indexerCmd represents the indexer command.
//
//nolint:exhaustruct,gochecknoglobals
var indexerCmd = &cobra.Command{
	Use:   "indexer",
	Short: "The main service of Agate",
	Long: `
The indexer service listens for new heads on the Ethereum network and aggregates data from a set
of relays. Once preprocessed, this data is stored in a database.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := make(chan error, 1)
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		ctx, cancel := context.WithCancel(context.Background())

		configuration := indexer.NewConfigurationFromViper(viper.GetViper())
		wg := sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer wg.Done()

			err <- run(ctx, configuration)
		}()

		defer func() {
			cancel()
			wg.Wait()
		}()

		for {
			select {
			case <-interrupt:
				log.Warn().Msg("agate indexer has been interrupted, stopping all services...")

				return nil
			case e := <-err:
				log.Warn().Msg("agate indexer encountered an error, stopping all services...")

				return e
			}
		}
	},
}

//nolint:wrapcheck,funlen
func run(ctx context.Context, configuration *indexer.Configuration) error {
	// Performs the database migration.
	migrator := storage_manager.NewDefaultDatabaseMigrator()

	if err := migrator.Init(
		configuration.DatabaseMigrationSourceURL,
		configuration.DatabaseConnURL,
	); err != nil {
		// TODO: Wrap error.
		return err
	}

	log.Info().Msg("applying database migrations..")

	if err := migrator.Migrate(); err != nil {
		// TODO: Wrap error.
		return err
	}

	log.Info().Msg("database migrations applied!")

	// Sets up beacon API client.
	beaconAPIClient := head_listener.NewAgateBeaconAPIClient()
	if err := beaconAPIClient.Init(ctx, configuration.BeaconAPIURL); err != nil {
		// TODO: Wrap error.
		return err
	}

	// Sets up head listener.
	listener := head_listener.NewAgateHeadListener()
	listener.Init(beaconAPIClient)

	// Sets up relay API clients.
	relayAPIClients := make([]data_aggregator.RelayAPIClient, 0)
	relaysDTOs := make([]*dto.Relay, 0)

	for _, relayAPIURL := range configuration.RelayAPIURLs {
		relayAPIClient := data_aggregator.NewAgateRelayAPIClient(relayAPIURL)

		if err := relayAPIClient.Init(); err != nil {
			// TODO: Wrap error.
			return err
		}

		relayAPIClients = append(relayAPIClients, relayAPIClient)
		relaysDTOs = append(relaysDTOs, &dto.Relay{URL: relayAPIURL})
	}

	// Sets up data aggregator.
	aggregator := data_aggregator.NewAgateDataAggregator()
	aggregator.Init(relayAPIClients...)

	// Creates data preprocessor.
	preprocessor := data_preprocessor.NewDataPreprocessor()

	// Sets up the database store.
	store := postgres.NewDefaultStore()
	if err := store.Init(ctx, configuration.DatabaseConnURL); err != nil {
		// TODO: Wrap error.
		return err
	}

	// Sets up storage manager and stores provided relays.
	storage := storage_manager.NewDefaultStorageManager()
	storage.Init(store)

	if err := storage.StoreRelays(ctx, relaysDTOs); err != nil {
		// TODO: Wrap error.
		return err
	}

	// Sets up main indexer service.
	service := indexer.NewIndexer(listener, aggregator, preprocessor, storage)

	log.Info().Msg("indexer service starting..")
	// Starts the indexer service. Blocks until context is done.
	if err := service.Start(ctx); err != nil {
		return err
	}

	service.Stop()
	log.Info().Msg("indexer service stopped!")

	return nil
}

func init() {
	rootCmd.AddCommand(indexerCmd)
	indexer.Flags(viper.GetViper(), indexerCmd.Flags())
}
