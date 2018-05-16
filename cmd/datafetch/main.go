package main

import (
	"flag"
	"github.com/SomniaStellarum/KittyAppraiser/ckabi"
	"github.com/SomniaStellarum/KittyAppraiser/cksaleabi"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"fmt"
)

var server = flag.String("host", "http://localhost:8545", "Set server host.")

func main() {
	flag.Parse()
	conn, err := ethclient.Dial(*server)
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
	endInt   := uint64(4690000)
	startInt := uint64(4680001)
	num := uint64(10000)

	for j := 0; j < 10; j++ {
		if j != 0 {
			endInt = endInt + num
			startInt = startInt + num
		}
		fmt.Print("Fetching sales from block ", startInt, " to block ", endInt, "\n")
		cksaleI, err := ck_sale.CKSaleFilterer.FilterAuctionSuccessful(&bind.FilterOpts{Start: startInt, End: &endInt, Context: nil})
		// 4608528

		fmt.Print("Sales Returned. Adding to Database.\n")

		if err != nil {
			log.Panic("Error fetching sales: ", err)
		}
		salesToStore := make([]salesData, 0)
		i := 0
		for {
			if cksaleI.Next() {
				var sale salesData
				sale.BlockNumber = cksaleI.Event.Raw.BlockNumber
				sale.TokenID = cksaleI.Event.TokenId
				sale.TotalPrice = cksaleI.Event.TotalPrice

				kitty, err := ck.CKCaller.GetKitty(nil, sale.TokenID)
				if err != nil {
					log.Panic("Error getting kitty: ", err)
				}
				sale.KittyGenes = kitty.Genes
				sale.KittyGeneration = kitty.Generation

				salesToStore = append(salesToStore, sale)
				if i == 1000 {
					i = 0
					fmt.Printf("Preparing %v sales to add to database\n", len(salesToStore))
				} else {
					i++
				}
				//b, err := json.Marshal(saleToStore)
				//if err != nil {
				//	log.Panic("JSON Marshal Error: ", err)
				//}
				//_, err = os.Stdout.Write(b)
				//if err != nil {
				//	log.Panic("Stdout Write Error: ", err)
				//}
			} else {
				fmt.Printf("Adding %v sales to database\n", len(salesToStore))

				err = insertSalesData(salesToStore)
				if err != nil {
					log.Panic("Error writing sales data to DB: ", err)
				}
				printBucketStats()
				return
			}
		}
	}
}