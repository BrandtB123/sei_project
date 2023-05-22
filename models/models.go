package models

import "time"

type PeerScore struct {
	Address string
	Score   int
}

type Block struct {
	Height     int64       `json:"height"`
	Timestamp  time.Time   `json:"timestamp"`
	TxCount    int         `json:"tx_count"`
	Proposer   string      `json:"proposer"`
	Validators []string    `json:"validators"`
	Peers      []PeerScore `json:"peers"`
}
