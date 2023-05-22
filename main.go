package main

import (
	"context"
	"fmt"
	"sei_project/models"
	"time"

	"github.com/go-pg/pg"
	tmClient "github.com/tendermint/tendermint/rpc/client/http"
)

func saveBlockData(db *pg.DB, block models.Block) error {
	_, err := db.Model(&block).OnConflict("(height) DO UPDATE").Set("tx_count = ?tx_count").Insert()
	if err != nil {
		return err
	}
	return nil
}

func getLatestHeight(db *pg.DB) (int64, error) {
	var block models.Block
	err := db.Model(&block).Order("height DESC").Limit(1).Select()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve the latest block height: %v", err)
	}

	return block.Height, nil
}

func collectData(ctx context.Context, rpcClient *tmClient.HTTP, db *pg.DB, prevBlock int64) error {
	status, err := rpcClient.Status(ctx)
	if err != nil {
		return err
	}
	latestBlockHeight := status.SyncInfo.LatestBlockHeight
	for i := prevBlock + 1; i <= latestBlockHeight; i++ {
		fmt.Println(i)
		block, err := rpcClient.Block(context.Background(), &i)
		if err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond)

		var parsedBlock models.Block
		parsedBlock.Height = block.Block.Height
		parsedBlock.Timestamp = block.Block.Header.Time
		parsedBlock.TxCount = len(block.Block.Txs)

		validators := make([]string, 0)
		for _, val := range block.Block.LastCommit.Signatures {
			validators = append(validators, val.ValidatorAddress.String())
		}
		parsedBlock.Validators = validators

		peers := make([]models.PeerScore, 0)

		v, err := rpcClient.Validators(ctx, &i, nil, nil)
		if err != nil {
			return err
		}

		if v != nil {
			for _, peer := range v.Validators {
				peers = append(peers, models.PeerScore{
					Address: peer.Address.String(),
					Score:   int(peer.VotingPower),
				})
			}
		}
		parsedBlock.Peers = peers
		err = saveBlockData(db, parsedBlock)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	db := pg.Connect(&pg.Options{
		Addr:     "sei-database-1.cs96vtk3gnbe.us-east-2.rds.amazonaws.com:5432",
		User:     "postgres",
		Password: "seipassword",
		Database: "postgres",
	})
	rpcClient, _ := tmClient.New("https://rpc.osmosis.zone/")
	ctx := context.Background()
	latestHeight, err := getLatestHeight(db)
	if err != nil {
		fmt.Println("error:", err)
	}
	for {
		err := collectData(ctx, rpcClient, db, latestHeight)
		if err != nil {
			fmt.Println("Error collecting data:", err)
		}

		time.Sleep(30 * time.Second)
	}

}
