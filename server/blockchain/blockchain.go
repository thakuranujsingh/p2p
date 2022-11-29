package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

type Blockchain struct {
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{data, prevBlockHash, ""}
	block.setHash()
	return block
}

func (bchain *Blockchain) AddBlock(data string) *Block {
	fmt.Printf("the print value %#v", bchain.Blocks)
	prevBlock := bchain.Blocks[len(bchain.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bchain.Blocks = append(bchain.Blocks, newBlock)

	return newBlock
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGensesBlock()}}
}

func NewGensesBlock() *Block {
	return NewBlock("Gensis block", "")
}
