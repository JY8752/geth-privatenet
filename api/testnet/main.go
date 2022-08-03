package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	cl, err := ethclient.Dial("https://eth-rinkeby.alchemyapi.io/v2/4MlUB9M-jIOmtxfOqtyFxtNDibyWG_cM")
	if err != nil {
		log.Fatalf("faile to initialize ethclient. err: %v", err)
	}

	ctx := context.Background()
	chaineId, err := cl.ChainID(ctx)
	if err != nil {
		log.Fatalf("faile to get chaineId. err: %v", err)
	}

	//rinkeby -> 4
	fmt.Printf("chaineId: %v\n", chaineId)

	//blockNumber
	blockNumber, _ := cl.BlockNumber(ctx)
	fmt.Printf("blockNumber: %v", blockNumber)

	address := common.HexToAddress("0x590C0655135261c97cDC4d0B91caa6fe55AEf43a")

	nft, err := NewNft(address, cl)
	if err != nil {
		log.Fatalf("faile to initialize NFT contract instance. err: %v", err)
	}

	baseURI, _ := nft.BaseURI(&bind.CallOpts{})

	fmt.Println("--------------- NFT contract --------------------")
	fmt.Printf("baseURI: %v\n", baseURI)

}
