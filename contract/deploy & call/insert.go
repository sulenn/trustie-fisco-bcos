package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	table "github.com/sulenn/trustie-fisco-bcos/contract/opensource"
)

func main() {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]

	client, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}

	// load the contract
	contractAddress := common.HexToAddress("0xcb465127073489235bBB8cF6995172a7c7420b3F")
	instance, err := table.NewOpenSource(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tabletestSession := &table.OpenSourceSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	token_name := "repo1"
	owner := "sulenn"
	total_supply := big.NewInt(1000000)
	username := []string{"qiubing"}
	balance := []*big.Int{big.NewInt(100000)}
	tx, receipt, err := tabletestSession.Insert(token_name, owner, total_supply, username, balance) // call Insert API
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	insertedLines, err := strconv.Atoi(receipt.Output[2:])
	if err != nil {
		log.Fatalf("error when transfer string to int: %v\n", err)
	}
	fmt.Printf("inserted lines: %v\n", insertedLines)
}
