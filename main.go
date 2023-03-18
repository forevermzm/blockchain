package main

import (
	"forevermzm/blockchain/pkg/core"
	"log"
)

func main() {
	log.Println("Hello World")
	genesisBlock := core.CreateBlock(core.Block{
		Index: 0,
		Hash:  "",
	}, core.Person{
		FirstName: "Hello",
		LastName:  "World",
		Age:       0,
	})
	log.Println(genesisBlock)

	currBlock := core.CreateBlock(genesisBlock, core.Person{"Good", "Boy", 2})
	valid := core.ValidateBlock(genesisBlock, currBlock)
	log.Println(currBlock)
	log.Printf("It is a valid block: %t", valid)
}
