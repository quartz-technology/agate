package indexer

import "fmt"

func NewIndexerListenerError(err error) error {
	return fmt.Errorf("failed to listen to head events: %w", err)
}
