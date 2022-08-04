# スマートコントラクト

Foundry を使ってみる。実装は Foundry のチュートリアル

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

## anvil

```terminal
[~/work/myapp/study/blockchain/geth-privatenet/contracts/hello] % anvil


                             _   _
                            (_) | |
      __ _   _ __   __   __  _  | |
     / _` | | '_ \  \ \ / / | | | |
    | (_| | | | | |  \ V /  | | | |
     \__,_| |_| |_|   \_/   |_| |_|

    0.1.0 (af3c9d3 2022-07-31T00:09:44.89474Z)
    https://github.com/foundry-rs/foundry

Available Accounts
==================

(0) 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266 (10000 ETH)
(1) 0x70997970c51812dc3a010c7d01b50e0d17dc79c8 (10000 ETH)
(2) 0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc (10000 ETH)
(3) 0x90f79bf6eb2c4f870365e785982e1f101e93b906 (10000 ETH)
(4) 0x15d34aaf54267db7d7c367839aaf71a00a2c6a65 (10000 ETH)
(5) 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc (10000 ETH)
(6) 0x976ea74026e726554db657fa54763abd0c3a0aa9 (10000 ETH)
(7) 0x14dc79964da2c08b23698b3d3cc7ca32193d9955 (10000 ETH)
(8) 0x23618e81e3f5cdf7f54c3d65f7fbc0abf5b21e8f (10000 ETH)
(9) 0xa0ee7a142d267c1f36714e4a8f75612f20a79720 (10000 ETH)

Private Keys
==================

(0) 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
(1) 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
(2) 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a
(3) 0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6
(4) 0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a
(5) 0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba
(6) 0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e
(7) 0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356
(8) 0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97
(9) 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6

Wallet
==================
Mnemonic:          test test test test test test test test test test test junk
Derivation path:   m/44'/60'/0'/0/


Base Fee
==================

1000000000

Gas Limit
==================

30000000

Listening on 127.0.0.1:8545
```

## script でテストネットにデプロイ

```terminal
//gethでテストネット繋ぐと30Gくらい容量持ってかれそうなのでやめる

//alchemy
source ../.env
forge script script/NFT.s.sol:MyScript --rpc-url $RINKEBY_RPC_URL --private-key $PRIVATE_KEY --broadcast --verify --etherscan-api-key $ETHERSCAN_KEY -vvvv

Error:
Etherscan could not detect the deployment.

デプロイは成功するけど検証で失敗する(issueもある)
https://github.com/foundry-rs/foundry/issues/2435

forge script script/NFT.s.sol:MyScript --rpc-url $RINKEBY_RPC_URL --private-key $PRIVATE_KEY --verify --etherscan-api-key $ETHERSCAN_KEY -vvvv

```

## script で anvil で起動したローカルのプライベートネットにデプロイ

```terminal
forge script script/NFT.s.sol:MyScript --fork-url http://localhost:8545 \
 --private-key $PRIVATE_KEY --broadcast
```
