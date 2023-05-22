package models

import (
	"time"

	"github.com/tendermint/tendermint/types"
)

type PeerScore struct {
	Address string
	Score   int
}

type Block struct {
	Height     int64       `json:"height,omitempty"`
	Timestamp  time.Time   `json:"timestamp,omitempty"`
	Txs        []types.Tx  `json:"txs,omitempty"`
	TxCount    int         `json:"tx_count,omitempty"`
	Proposer   string      `json:"proposer,omitempty"`
	Validators []string    `json:"validators,omitempty"`
	Peers      []PeerScore `json:"peers,omitempty"`
}
