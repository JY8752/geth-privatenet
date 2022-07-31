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