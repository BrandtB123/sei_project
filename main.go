package main

import (
	"context"
	"fmt"
	"net/http"
	"sei_project/models"
	"sei_project/router"
	"time"

	tmClient "github.com/tendermint/tendermint/rpc/client/http"
)

func collectData(ctx context.Context, rpcClient *tmClient.HTTP, prevBlock int64) (int64, error) {
	status, err := rpcClient.Status(ctx)
	if err != nil {
		return prevBlock, err
	}
	latestBlockHeight := status.SyncInfo.LatestBlockHeight
	for i := prevBlock; i < latestBlockHeight; i++ {
		block, err := rpcClient.Block(context.Background(), &i)

		for err != nil || block == nil {
			time.Sleep(2 * time.Second)

			block, err = rpcClient.Block(context.Background(), &i)

			if err != nil {
				return i, err
			}
		}
		fmt.Println(block.Block.Height)
		time.Sleep(500 * time.Millisecond)
		var parsedBlock models.Block
		parsedBlock.Height = block.Block.Height
		parsedBlock.Timestamp = block.Block.Header.Time
		parsedBlock.Txs = block.Block.Txs
		parsedBlock.TxCount = len(block.Block.Txs)
		parsedBlock.Proposer = block.Block.ProposerAddress.String()
		validators := make([]string, 0)
		for _, val := range block.Block.LastCommit.Signatures {
			validators = append(validators, val.ValidatorAddress.String())
		}
		parsedBlock.Validators = validators

		peers := make([]models.PeerScore, 0)

		v, err := rpcClient.Validators(ctx, &i, nil, nil)
		if err != nil {
			return i, err
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

		err = models.SaveBlockData(parsedBlock)
		if err != nil {
			return i, err
		}
	}
	return latestBlockHeight, nil
}

func main() {
	models.NewDB()
	go func() {
		fmt.Println("Running")
		r := router.Router()
		err := http.ListenAndServe(":80", r)
		if err != nil {
			fmt.Println("Error establishing connection:", err)
		}
	}()

	rpcClient, err := tmClient.New("https://rpc.osmosis.zone/")

	if err != nil {
		fmt.Println("Error establishing client:", err)
	}
	ctx := context.Background()
	latestHeight, err := models.GetLatestHeight()
	if err != nil {
		latestHeight = 9739047
	}
	for {
		latestHeight, err = collectData(ctx, rpcClient, latestHeight)
		if err != nil {
			fmt.Println("Error collecting data:", err)
		}
		time.Sleep(30 * time.Second)
	}

}
