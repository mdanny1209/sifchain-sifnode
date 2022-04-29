#!/usr/bin/env bash

set -x

sifnoded tx margin open \
  --from $SIF_ACT \
  --keyring-backend test \
  --borrow_asset cusdt \
  --collateral_asset rowan \
  --collateral_amount 1000 \
  --position long \
  --fees 100000000000000000rowan \
  --node ${SIFNODE_NODE} \
  --chain-id $SIFNODE_CHAIN_ID \
  --broadcast-mode block \
  -y