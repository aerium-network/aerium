package tests

import (
	"fmt"
	"time"

	"github.com/aerium-network/aerium/util/logger"
	aerium "github.com/aerium-network/aerium/www/grpc/gen/go"
)

func lastHeight() uint32 {
	res, err := tBlockchainClient.GetBlockchainInfo(tCtx,
		&aerium.GetBlockchainInfoRequest{})
	if err != nil {
		panic(err)
	}

	return res.LastBlockHeight
}

func waitForNewBlocks(num uint32) {
	for i := uint32(0); i < num; i++ {
		height := lastHeight()
		if lastHeight() > height {
			break
		}
		time.Sleep(4 * time.Second)
	}
}

func lastBlock() *aerium.GetBlockResponse {
	return getBlockAt(lastHeight())
}

func getBlockAt(height uint32) *aerium.GetBlockResponse {
	for i := 0; i < 120; i++ {
		res, err := tBlockchainClient.GetBlock(tCtx,
			&aerium.GetBlockRequest{
				Height:    height,
				Verbosity: aerium.BlockVerbosity_BLOCK_VERBOSITY_INFO,
			},
		)
		if err != nil {
			fmt.Printf("getBlockAt err: %s\n", err.Error())
			time.Sleep(1 * time.Second)

			continue
		}

		return res
	}
	logger.Panic("getBlockAt timeout", "height", height)

	return nil
}
