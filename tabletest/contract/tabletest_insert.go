package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	table "github.com/sulenn/trustie-fisco-bcos/tabletest"
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
	contractAddress := common.HexToAddress("0x904ed579402eD8BBb80ee7F0eb02e8226d78a70f") // 0x9526BDd51d7F346ec2B48192f25a800825A8dBF3
	instance, err := table.NewTableTest(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tabletestSession := &table.TableTestSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	name := "Bob"
	item_id := big.NewInt(100010001001)
	item_name := "Laptop"
	tx, receipt, err := tabletestSession.Insert(name, item_id, item_name) // call Insert API
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
