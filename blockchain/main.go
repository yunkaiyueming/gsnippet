package main

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

var BlockChain []Block

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func AddBlock(BPM int) {
	newBlock, err := generateBlock(BlockChain[len(BlockChain)-1], BPM)
	if err != nil {
		os.Exit(200)
	}

	isValid := isBlockValid(newBlock, BlockChain[len(BlockChain)-1])
	if !isValid {
		os.Exit(200)
	}

	BlockChain = append(BlockChain, newBlock)
}

func main() {
	t := time.Now()
	genesisBlock := Block{0, t.String(), 0, "", ""}
	BlockChain = append(BlockChain, genesisBlock)

	//TEST
	AddBlock(10)
	AddBlock(20)
	AddBlock(30)
	AddBlock(40)
	spew.Dump(BlockChain)
}
