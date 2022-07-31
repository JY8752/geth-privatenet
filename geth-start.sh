#!bin/bash

geth --datadir . --networkid 15 \
--nodiscover --maxpeers 0 --mine --miner.threads 1 \
--http --http.addr "0.0.0.0" --http.corsdomain "*" \
--http.vhosts "*" --http.api "eth,web3,personal,net" \
--ipcpath /tmp/blockchain/geth-privatenet/geth.ipc --ws --ws.addr "0.0.0.0" \
--ws.api "eth,web3,personal,net" --ws.origins "*" \
--unlock 0,1 --password ./password --allow-insecure-unlock
