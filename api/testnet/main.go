package main

import (
	"context"
	"fmt"
	"log"

	"testnet/nft"

	env "env"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	cl, err := ethclient.Dial(env.RinkebyRpcUrl())
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
	fmt.Printf("blockNumber: %v\n", blockNumber)

	address := common.HexToAddress(env.CONTRACT_ADDRESS())

	nft, err := nft.NewNft(address, cl)
	if err != nil {
		log.Fatalf("faile to initialize NFT contract instance. err: %v", err)
	}

	baseURI, _ := nft.BaseURI(&bind.CallOpts{})
	currentTokenId, _ := nft.CurrentTokenId(&bind.CallOpts{})
	totalSupply, _ := nft.TOTALSUPPLY(&bind.CallOpts{})
	mintPrice, _ := nft.MINTPRICE(&bind.CallOpts{})
	name, _ := nft.Name(&bind.CallOpts{})
	symbol, _ := nft.Symbol(&bind.CallOpts{})

	fmt.Println("--------------- NFT contract --------------------")
	fmt.Printf("baseURI: %v\n", baseURI)
	fmt.Printf("currentTokenId: %v\n", currentTokenId)
	fmt.Printf("totalSupply: %v\n", totalSupply)
	fmt.Printf("mintPrice: %v\n", mintPrice)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("symbol: %v\n", symbol)

	fmt.Println("----------- mintTo ------------------------------")

	privateKey, err := crypto.HexToECDSA(env.PrivateKey())
	if err != nil {
		log.Fatalf("faile to initialize PrivateKey. err: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chaineId)
	if err != nil {
		log.Fatalf("faile to initialize transactor. err: %v", err)
	}
	auth.Value = mintPrice //wei
	recipient := common.HexToAddress(env.PUBLIC_KEY())

	tx, err := nft.MintTo(auth, recipient)
	if err != nil {
		log.Fatalf("failed mitTo. err: %v", err)
	}

	tokenURI, err := nft.TokenURI(&bind.CallOpts{}, currentTokenId)
	if err != nil {
		log.Fatalf("failed get tokenURI. err: %v", err)
	}

	fmt.Printf("tx: %v\n", tx)
	fmt.Printf("tokenUri: %v\n", tokenURI)

}
