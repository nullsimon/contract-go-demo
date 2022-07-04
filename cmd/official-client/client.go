package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Block struct {
	Number string
}

func main() {
	client, err := rpc.Dial("https://mainnet.infura.io/v3/<Your-Project-ID>")
	if err != nil {
		log.Fatalf("Could not connect to Infura: %v", err)
	}

	var lastBlock Block
	err = client.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
	if err != nil {
		fmt.Println("Cannot get the latest block:", err)
		return
	}

	fmt.Printf("Latest block: %v\n", lastBlock.Number)
	value, err := strconv.ParseInt(lastBlock.Number[2:], 16, 64)
	if err != nil {
		fmt.Println("format string to int err:", err)
	}
	fmt.Printf("Latest block: %v\n", value)

	ws, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/<Your-Project-ID>")

	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := ws.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())

			block, err := ws.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())
			fmt.Println(block.Number().Uint64())
			fmt.Println(block.Time())
			fmt.Println(block.Nonce())
			fmt.Println(len(block.Transactions()))
		}
	}

}
