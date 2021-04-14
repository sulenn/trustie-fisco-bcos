package main

import (
	"fmt"
	"log"

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
	contractAddress := common.HexToAddress("0x904ed579402eD8BBb80ee7F0eb02e8226d78a70f") // 0x481D3A1dcD72cD618Ea768b3FbF69D78B46995b0
	instance, err := table.NewTableTest(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tabletestSession := &table.TableTestSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	name := "Bob"

	names, item_ids, item_names, err := tabletestSession.Select(name) // call select API
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(names); i++ {
		fmt.Printf("name: %v, item_id: %v, item_name: %v \n", names[i], item_ids[i], item_names[i])
	}

}
