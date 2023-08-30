package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Contract represents the smart contract for adding two numbers
type Contract struct {
	Address common.Address
	ABI     *abi.ABI
}

// NewContract creates a new Contract instance
func NewContract(address common.Address, abi *abi.ABI) *Contract {
	return &Contract{
		Address: address,
		ABI:     abi,
	}
}

// Add adds two numbers together
func (c *Contract) Add(a, b *big.Int) (*big.Int, error) {
	// Create a new transaction
	tx := types.NewTransaction(
		types.BlockNumber(block.Number()),
		c.Address,
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		types.TxData{
			Data: c.ABI.Pack("add", a, b),
		},
	)

	// Sign the transaction
	signedTx, err := types.SignTx(tx, c.PrivateKey)
	if err != nil {
		return nil, err
	}

	// Send the transaction
	err = client.SendTransaction(signedTx)
	if err != nil {
		return nil, err
	}

	// Wait for the transaction to be mined
	err = client.WaitForTransaction(signedTx.Hash(), nil)
	if err != nil {
		return nil, err
	}

	// Get the transaction receipt
	receipt, err := client.GetTransactionReceipt(signedTx.Hash())
	if err != nil {
		return nil, err
	}

	// Get the result of the addition
	result := new(big.Int)
	err = rlp.DecodeBytes(receipt.Logs[0].Data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}