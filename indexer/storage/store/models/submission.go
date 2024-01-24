package models

import (
	"time"

	"github.com/attestantio/go-eth2-client/spec/phase0"
)

type Submission struct {
	ID           uint64
	RelayID      uint64
	BidBlockHash phase0.Hash32

	IsDelivered  bool
	IsOptimistic bool
	SubmittedAt  time.Time
}
