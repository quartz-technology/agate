package dto

import (
	"time"
)

type Submission struct {
	RelayURL string

	IsDelivered  bool
	IsOptimistic bool
	SubmittedAt  time.Time
}
