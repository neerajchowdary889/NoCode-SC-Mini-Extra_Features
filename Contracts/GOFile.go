package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connect to an Ethereum node
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/<YOUR_PROJECT_ID>")
	if err != nil {
		panic(err)
	}

	// Get the current block number
	blockNumber := client.BlockNumber()

	// Get the list of all live transactions
	transactions, err := client.Transactions(blockNumber)
	if err != nil {
		panic(err)
	}

	// Print the list of transactions
	for _, transaction := range transactions {
		fmt.Println("Transaction:", transaction.Hash().Hex())
		fmt.Println("From:", transaction.From.Hex())
		fmt.Println("To:", transaction.To.Hex())
		fmt.Println("Value:", transaction.Value.String())
		fmt.Println("Gas:", transaction.Gas.String())
		fmt.Println("Gas Price:", transaction.GasPrice.String())
		fmt.Println("Nonce:", transaction.Nonce.String())
	}
}