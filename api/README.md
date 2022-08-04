# Go packageを使用してスマートコントラクトにアクセスする

## setup

```
go get -d github.com/ethereum/go-ethereum
go get github.com/ethereum/go-ethereum/rpc@v1.10.21
go get github.com/ethereum/go-ethereum/accounts/keystore@v1.10.21

//.envファイル読み込むのに
go get github.com/joho/godotenv
```

## abigen

```
//abiファイルがなければ下記コマンドでabiファイルをビルド時に生成しておく
forge clean
forge build --extra-output-files abi

abigen --abi NFT.abi.json --pkg nft --type nft --out nft.go
```
