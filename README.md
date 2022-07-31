# gethでプライベートネット作成

## setup

- Genesisファイルを作成する
- genesisiブロックの初期化

```terminal
geth --datadir . init ./myGenesis.json
```

- gethの起動

```terminal
//ipcpathでデータディレクトリ以外を指定しないとipcファイルが生成されなかったので
//フォワードで起動状態
//バックグラウンドで起動できないの謎
$ geth --networkid "15" --nodiscover --datadir . --ipcpath /tmp/blockchain/geth-privatenet/geth.ipc 

//別のterminalでattach
geth attach /tmp/blockchain/geth-privatenet/geth.ipc 
```

## アカウントの作成

- ethereumにはアカウントは2種類ある。
- EOAアカウントとコントラクトアカウント

```terminal
//geth コンソール

//アカウントの確認
> eth.accounts
[ ]

//アカウントの作成
> personal.newAccount("pass")
"0x78fe7b62a9b27b9fe38c8b8900024a4dd4ce301a"
> eth.accounts
["0x78fe7b62a9b27b9fe38c8b8900024a4dd4ce301a"]

//coinbase(マイニングしたときに報酬を紐づけるアカウントの表示)
> eth.coinbase
"0x78fe7b62a9b27b9fe38c8b8900024a4dd4ce301a"

//coinbaseの変更
> miner.setEtherbase(eth.accounts[1])
true
> eth.coinbase
"0x5520c468b2a2f8da55f91d8e99fbc941cca884fb"
```

## マイニング

```terminal
//geth コンソール

//マイニング開始
miner.start()

//マイニングの停止
miner.stop()

//ブロック高の確認
> eth.blockNumber
15

//マイニングしてるかの確認
> eth.mining
false

//マイニングしたブロックの確認
> eth.getBlock(15)
{
  difficulty: 131904,
  extraData: "0xd983010a14846765746888676f312e31382e338664617277696e",
  gasLimit: 132265053,
  gasUsed: 0,
  hash: "0x1014c84b0946ed269a16ec44e116bbde9eb29be16e2abc23cd9e4a5ef090b3a2",
  logsBloom: "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  miner: "0x5520c468b2a2f8da55f91d8e99fbc941cca884fb",
  mixHash: "0xb0de74a3e2270ef308907b856c6025615f68bab65484444e72efc6922b96a309",
  nonce: "0x61243a80122c8022",
  number: 15,
  parentHash: "0x9094f9eec2e21dec2299cba52a1eac3575e3440af1c2cff02326378968759310",
  receiptsRoot: "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
  sha3Uncles: "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
  size: 538,
  stateRoot: "0xf58e6a5dd845cb3599c0f94af046c2655c0ddd2400cffcd98d77be07dc0b28d8",
  timestamp: 1658885318,
  totalDifficulty: 1988288,
  transactions: [],
  transactionsRoot: "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
  uncles: []
}

//報酬の確認
> eth.getBalance(eth.accounts[0])
0
> eth.getBalance(eth.accounts[1])
30000000000000000000
> web3.fromWei(eth.getBalance(eth.accounts[1]), "ether")
30
```

## 送金

```terminal
//アカウントのロック解除
> personal.unlockAccount(eth.accounts[1])
Unlock account 0x5520c468b2a2f8da55f91d8e99fbc941cca884fb
Passphrase: 
true

> eth.sendTransaction({from: eth.accounts[1], to: eth.accounts[0], value: web3.toWei(0.01, "ether")})
"0xff96523023f346f89978e7ea96dc99591a5482650401d25f2458f11f61aa4735"

//送金はバックグラウンドでマイニングの処理を走らせておかないと成功しない

//トランザクションの確認
> eth.getTransaction('0xac7f81df9e54d203d8eb55d35f08a632e9888e71a31ff8538baacc0b61196dc3')
{
  blockHash: "0x879300adbf73b6922f088d659d00da3cb1a9240d43608a6da4845a8a0667623e",
  blockNumber: 224,
  from: "0x5520c468b2a2f8da55f91d8e99fbc941cca884fb",
  gas: 21000,
  gasPrice: 1000000000,
  hash: "0xac7f81df9e54d203d8eb55d35f08a632e9888e71a31ff8538baacc0b61196dc3",
  input: "0x",
  nonce: 4,
  r: "0xcbc248c64ca1efbf458075e88508af085de8872fd519ea9dda383008f1b20017",
  s: "0x7558c078324b09f791267d55785302e7fb50e936bb47a9b51d6c132cff7c9973",
  to: "0x78fe7b62a9b27b9fe38c8b8900024a4dd4ce301a",
  transactionIndex: 0,
  type: "0x0",
  v: "0x42",
  value: 10000000000000000
}
```

## スマートコントラクト

### コンパイラインストール

```terminal
brew update
brew upgrade
brew tap ethereum/ethereum
brew install solidity
//brew linkapps solidity これはエラーでこける
```

## gethの起動

```terminal
//起動スクリプト実行　(アカウントのアンロックの指定のところがよくわからない)
/bin/bash geth-start.sh

//起動確認
geth attach http://localhost:8545

>eth.maning 
true
```