package main

import (
	"errors"
	"fmt"
	"forevermzm/blockchain/pkg/core"
	"forevermzm/blockchain/pkg/handler"
	"log"
	"net/http"
	"os"
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
	handler.Blockchain = core.Blockchain{
		Chain: []core.Block{
			genesisBlock,
		},
	}
	log.Println(handler.Blockchain)

	http.HandleFunc("/blockchain", handler.HandleBlockchain)

	err := http.ListenAndServe(":8090", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
