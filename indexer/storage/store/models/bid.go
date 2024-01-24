package models

import (
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/holiman/uint256"
	"github.com/quartz-technology/agate/indexer/storage/store/dto"
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

func (bid *Bid) FromBidDTO(dto *dto.Bid) *Bid {
	bid.Slot = dto.Slot
	bid.ParentHash = dto.ParentHash
	bid.BlockHash = dto.BlockHash
	bid.FeeRecipient = dto.FeeRecipient
	bid.GasLimit = dto.GasLimit
	bid.GasUsed = dto.GasUsed
	bid.Value = dto.Value
	bid.NumTx = dto.NumTx
	bid.Proposer = dto.Proposer
	bid.Builder = dto.Builder

	return bid
}
