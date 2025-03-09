package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"time"
)

type Block struct {
	Hash      []byte
	Data      []byte
	PrevHash  []byte
	Timestamp []byte
}

func (b *Block) deriveHash() {

	info := bytes.Join([][]byte{b.Data, b.PrevHash, b.Timestamp}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	var timestamp []byte = make([]byte, 8)
	now := time.Now().UnixMicro()
	binary.LittleEndian.PutUint64(timestamp, uint64(now))

	block := &Block{[]byte{}, []byte(data), prevHash, timestamp}
	block.deriveHash()
	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
