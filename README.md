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
