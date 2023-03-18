package core

import (
	"fmt"
)

type Blockchain struct {
	Chain []Block
}

func (bc Blockchain) AddBlock(block Block) (Blockchain, error) {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	if !ValidateBlock(lastBlock, block) {
		return bc, fmt.Errorf("adding a invalid block %v", block)
	}
	return Blockchain{Chain: append(bc.Chain, block)}, nil
}
