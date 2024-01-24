package indexer

import "fmt"

func NewIndexerListenerError(err error) error {
	return fmt.Errorf("indexer failed to start listening to head events: %w", err)
}
