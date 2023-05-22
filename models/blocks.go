package models

import (
	"fmt"
	"sort"

	"github.com/go-pg/pg"
)

var DB *pg.DB

func NewDB() {
	db := pg.Connect(&pg.Options{
		Addr:     "sei-database-1.cs96vtk3gnbe.us-east-2.rds.amazonaws.com:5432",
		User:     "postgres",
		Password: "seipassword",
		Database: "postgres",
	})
	DB = db
}

func SaveBlockData(block Block) error {
	_, err := DB.Model(&block).OnConflict("(height) DO UPDATE").Set("tx_count = ?tx_count").Insert()
	if err != nil {
		return err
	}
	return nil
}

func GetLatestHeight() (int64, error) {
	var block Block
	err := DB.Model(&block).Order("height DESC").Limit(1).Select()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve the latest block height: %v", err)
	}

	return block.Height, nil
}

func GetTransactionsInPastNBlocks(n int) (int, error) {
	var block Block
	var totalTxs int
	err := DB.Model(&block).ColumnExpr("SUM(tx_count)").Limit(n).Select(&totalTxs)
	if err != nil {
		return 0, err
	}

	return totalTxs, nil
}

func GetProposedBlocksByValidator(proposer string) ([]int64, error) {
	var heights []int64
	_, err := DB.Query(&heights, `
		SELECT height
		FROM blocks
		WHERE proposer = ?
	`, proposer)
	if err != nil {
		return nil, err
	}
	return heights, nil
}

func GetTopNPeersByScore(n int) ([]PeerScore, error) {
	var blocks []Block
	err := DB.Model(&blocks).
		Order("height DESC").
		Limit(n).
		Select()

	if err != nil {
		return nil, err
	}

	peers := make(map[string]int)
	for _, block := range blocks {
		for _, score := range block.Peers {
			peers[score.Address] += score.Score
		}
	}

	var topPeers []PeerScore

	for address, score := range peers {
		topPeers = append(topPeers, PeerScore{
			Address: address,
			Score:   score,
		})
	}

	sort.Slice(topPeers, func(i, j int) bool {
		return topPeers[i].Score > topPeers[j].Score
	})

	if len(topPeers) > n {
		topPeers = topPeers[:n]
	}
	return topPeers, nil
}
