package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("環境変数の読み込みに失敗しました. err: %v", err)
	}
}

func RinkebyRpcUrl() string {
	return os.Getenv("RINKEBY_RPC_URL")
}

func PrivateKey() string {
	return os.Getenv("PRIVATE_KEY")
}

func ETHERSCAN_KEY() string {
	return os.Getenv("ETHERSCAN_KEY")
}

func PUBLIC_KEY() string {
	return os.Getenv("PUBLIC_KEY")
}

func CONTRACT_ADDRESS() string {
	return os.Getenv("CONTRACT_ADDRESS")
}
