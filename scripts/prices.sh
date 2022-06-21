#!/bin/bash
curl -X 'GET' 'https://api.coingecko.com/api/v3/simple/price?ids=bitcoin%2Cethereum%2Cavalanche-2%2Ccardano%2Csolana%2Cpolkadot&vs_currencies=eur%2Cbtc' -H 'accept: application/json' -o ~/.config/conky/Graffias/prices.json
