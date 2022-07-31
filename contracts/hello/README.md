# スマートコントラクト
Foundryを使ってみる。実装はFoundryのチュートリアル

## install

```terminal
curl -L https://foundry.paradigm.xyz | bash

//.zshrcの更新
source ~/.zshrc 

foundryup
```

## setup

```terminal
//vscodeで実装する場合、remappings.txtが必要なため、オプションをつけて初期化
forge init hello --vscode

//ビルド
forge build

//テスト
forge test
```

## NFT with Solmate

### deploy

```terminal
forge install transmissions11/solmate Openzeppelin/openzeppelin-contracts

//ライブラリを追加したら下記コマンドでremappingを表示してremappins.txtを更新する必要がある
forge remappings 


//ガス代を払う必要があるため自分のウォレットに送金する必要がある
export RPC_URL=http://localhost:8545
export PRIVATE_KEY=<my private key>

forge create NFT --rpc-url=$RPC_URL --private-key=$PRIVATE_KEY --constructor-args MyNFT ETH　--legacy

Deployer: 0x4c8169984b34f667aFD655cB3BFFBf5CAf844997
Deployed to: 0x1997076e35B519f5001022a904F7a06B57A4eD73
Transaction hash: 0x50780b53e2dc59868120eb59181bdc3465d8e29b7a9fdeef0b9d996e1b405f64

```

### mint

```terminal
cast send --rpc-url=$RPC_URL 0x1997076e35B519f5001022a904F7a06B57A4eD73  "mintTo(address)" 0x4c8169984b34f667aFD655cB3BFFBf5CAf844997 --private-key=$PRIVATE_KEY --legacy

blockHash               0x25d919b5c9c22e60531aee69e0f83df372cf8783a802632994cf77a658121d8b
blockNumber             3518
contractAddress         
cumulativeGasUsed       90805
effectiveGasPrice       1000000000
gasUsed                 90805
logs                    [{"address":"0x1997076e35b519f5001022a904f7a06b57a4ed73","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x0000000000000000000000004c8169984b34f667afd655cb3bffbf5caf844997","0x0000000000000000000000000000000000000000000000000000000000000001"],"data":"0x","blockHash":"0x25d919b5c9c22e60531aee69e0f83df372cf8783a802632994cf77a658121d8b","blockNumber":"0xdbe","transactionHash":"0x29c49ba51e6b83c53c4a90c855a23e13c8737ab5ec64db87b960d515a46bad22","transactionIndex":"0x0","logIndex":"0x0","removed":false}]
logsBloom               0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000040000000001000000000000000008000000000000000000040000000000000000000000000000020000000000000000000800000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000002000000000000100000000000000000000000000000000000000060000000000000000000000080000000000000000000000000000000000000000000
root                    
status                  1
transactionHash         0x29c49ba51e6b83c53c4a90c855a23e13c8737ab5ec64db87b960d515a46bad22
transactionIndex        0
type                    0

//mintできたか確認
cast call --rpc-url=$RPC_URL --private-key=$PRIVATE_KEY  "ownerOf(uint256)" 1 --legacy

> 0x0000000000000000000000004c8169984b34f667afd655cb3bffbf5caf844997
```

