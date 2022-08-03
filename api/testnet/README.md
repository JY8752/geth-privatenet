# Go packageを使用してスマートコントラクトにアクセスする

## setup

```
go get -d github.com/ethereum/go-ethereum
go get github.com/ethereum/go-ethereum/rpc@v1.10.21
go get github.com/ethereum/go-ethereum/accounts/keystore@v1.10.21
```

## abigen

```
//abiファイルがなければ下記コマンドでabiファイルをビルド時に生成しておく
forge clean
forge build --extra-output-files abi

abigen --abi NFT.abi.json --pkg main --type nft --out nft.go
```
