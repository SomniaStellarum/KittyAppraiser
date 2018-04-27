package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/SomniaStellarum/KittyAppraiser/ckabi"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"github.com/SomniaStellarum/KittyAppraiser/cksaleabi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Panic("Error with dialing: ", err)
	}

	ck, err := ckabi.NewCK(common.HexToAddress("0x06012c8cf97BEaD5deAe237070F9587f8E7A266d"), conn)
	if err != nil {
		log.Panic("Error Instatiating CK: ", err)
	}

	ck_sale, err := cksaleabi.NewCKSale(common.HexToAddress("0xb1690C08E213a35Ed9bAb7B318DE14420FB57d8C"), conn)
	if err != nil {
		log.Panic("Error Instatiating CK: ", err)
	}

	//addr, err := ck.CKCaller.SaleAuction(nil)
	//if err != nil {
	//	log.Panic("Error Fetching Sale Address: ", err)
	//}
	//fmt.Print("Sale Address: ", addr.Hex(), "\n")

	cksaleI, err := ck_sale.CKSaleFilterer.FilterAuctionSuccessful(&bind.FilterOpts{Start: 5640000, End: nil, Context: nil})
	// 4608528
	if err != nil {
		log.Panic("Error fetching sales: ", err)
	}
	fmt.Print(cksaleI.Event)
	for i := 0; i < 10; i++ {
		if cksaleI.Next() {
			id := cksaleI.Event.TokenId
			kitty, err := ck.CKCaller.GetKitty(nil, id)
			if err != nil {
				log.Panic("Error getting kitty: ", err)
			}
			fmt.Print("Token ID: ", id,
				"\nToken Price: ", cksaleI.Event.TotalPrice,
				"\nKittyGenes: ", kitty.Genes, "\n")
		}
	}
}