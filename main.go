package main

import (
	"awesomeProject1/block"
	"encoding/binary"
	"fmt"
	"time"
)

func blockchainRun() {
	chain := block.InitBlockChain()
	time.Sleep(2 * time.Second)
	chain.AddBlock("First Block after Genesis")
	time.Sleep(2 * time.Second)
	chain.AddBlock("Second Block after Genesis")
	time.Sleep(2 * time.Second)
	chain.AddBlock("Third Block after Genesis")

	for _, b := range chain.Blocks {
		fmt.Printf("\n================================\n")
		fmt.Printf("Previous Hash: %x\n", b.PrevHash)
		fmt.Printf("Data in Block: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		now := binary.LittleEndian.Uint64(b.Timestamp)
		fmt.Printf("Timestamp: %v\n", time.UnixMicro(int64(now)))

	}
}
func main() {
	blockchainRun()
}
