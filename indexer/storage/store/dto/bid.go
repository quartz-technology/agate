package dto

import (
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/holiman/uint256"
)

type Bid struct {
	Slot       phase0.Slot
	ParentHash phase0.Hash32
	BlockHash  phase0.Hash32

	FeeRecipient bellatrix.ExecutionAddress
	GasLimit     uint64
	GasUsed      uint64
	Value        *uint256.Int
	NumTx        uint

	Proposer phase0.BLSPubKey
	Builder  phase0.BLSPubKey
}
