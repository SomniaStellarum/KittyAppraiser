package main

import "math/big"

type salesData struct {
	BlockNumber     uint64 `json:"blocknumber"`
	TokenID         *big.Int `json:"tokenid"`
	TotalPrice      *big.Int `json:"totalprice"`
	KittyGenes      *big.Int `json:"kittygenes"`
	KittyGeneration *big.Int `json:"kittygeneration"`
}
