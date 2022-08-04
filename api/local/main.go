package main

import (
	"context"
	"fmt"
	"local/nft"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	//anvilにデプロイしたコントラクトアドレス
	CONTRACT_ADDRESS string = "0x5fbdb2315678afecb367f032d93f642f64180aa3"
	//anvilで作成されたアカウント情報
	PUBLIC_KEY  string = "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"
	PRIVATE_KEY string = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
)

func main() {
	cl, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("faile to initialize ethclient. err: %v", err)
	}

	ctx := context.Background()
	chaineId, err := cl.ChainID(ctx)
	if err != nil {
		log.Fatalf("faile to get chaineId. err: %v", err)
	}

	fmt.Printf("chaineId: %v\n", chaineId)

	//blockNumber
	blockNumber, _ := cl.BlockNumber(ctx)
	fmt.Printf("blockNumber: %v\n", blockNumber)

	address := common.HexToAddress(CONTRACT_ADDRESS)

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

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		log.Fatalf("faile to initialize PrivateKey. err: %v", err)
	}

	gasPrice, err := cl.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("faile to get gasPrice. err: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chaineId)
	if err != nil {
		log.Fatalf("faile to initialize transactor. err: %v", err)
	}
	auth.Value = mintPrice //wei
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000)
	recipient := common.HexToAddress(PUBLIC_KEY)

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
