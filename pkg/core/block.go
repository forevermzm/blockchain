package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) text() string {
	return fmt.Sprintf("First:%s, Last:%s, Age:%d", p.FirstName, p.LastName, p.Age)
}

type Block struct {
	Index     int
	Timestamp string
	Data      Person
	Hash      string
	PrevHash  string
}

func CreateBlock(oldBlock Block, data Person) Block {
	var block Block
	block.Index = oldBlock.Index + 1
	block.Timestamp = time.Now().Format(time.RFC3339)
	block.Data = data
	block.PrevHash = oldBlock.Hash
	block.Hash = calculateHash(block)
	return block
}

func ValidateBlock(prev Block, curr Block) bool {
	if prev.Index+1 != curr.Index {
		return false
	}

	if prev.Hash != curr.PrevHash {
		return false
	}

	if calculateHash(curr) != curr.Hash {
		return false
	}

	return true
}

func calculateHash(block Block) string {
	record := fmt.Sprintf("%d", block.Index) + block.Timestamp + block.Data.text() + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
