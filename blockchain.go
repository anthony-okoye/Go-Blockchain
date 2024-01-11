package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index        int
	Timestamp    int64
	Data         string
	PreviousHash string
	Hash         string
}

// Blockchain represents the entire blockchain
type Blockchain struct {
	Chain []Block
}

// CalculateHash calculates the hash of a block
func CalculateHash(block Block) string {
	record := fmt.Sprintf("%d%d%s%s", block.Index, block.Timestamp, block.Data, block.PreviousHash)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// GenerateBlock creates a new block in the blockchain
func GenerateBlock(previousBlock Block, data string) Block {
	newBlock := Block{
		Index:        previousBlock.Index + 1,
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreviousHash: previousBlock.Hash,
	}
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	latestBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := GenerateBlock(latestBlock, data)
	bc.Chain = append(bc.Chain, newBlock)
}

// IsValid checks if a blockchain is valid
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		previousBlock := bc.Chain[i-1]

		if currentBlock.Hash != CalculateHash(currentBlock) {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func main() {
	// Create genesis block
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		Data:         "Genesis Block",
		PreviousHash: "",
	}
	genesisBlock.Hash = CalculateHash(genesisBlock)

	// Create blockchain and add the genesis block
	blockchain := Blockchain{}
	blockchain.Chain = append(blockchain.Chain, genesisBlock)

	// Add some blocks to the blockchain
	blockchain.AddBlock("Data for Block 1")
	blockchain.AddBlock("Data for Block 2")

	// Print the blockchain
	for _, block := range blockchain.Chain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("------------------------------")
	}

	// Check if the blockchain is valid
	if blockchain.IsValid() {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
}
