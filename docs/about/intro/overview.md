<!--
order: 1
-->

# High-level Overview

Learn about Kynno and its primary features. {synopsis}

## What is Kynno

Kynno is a scalable, high-throughput Proof-of-Stake blockchain that is fully compatible and
interoperable with Ethereum. It's built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/) which runs on top of [Tendermint Core](https://github.com/tendermint/tendermint) consensus engine.

Kynno allows for running vanilla Ethereum as a [Cosmos](https://cosmos.network/)
application-specific blockchain. This allows developers to have all the desired features of
Ethereum, while at the same time, benefit from Tendermint’s PoS implementation. Also, because it is
built on top of the Cosmos SDK, it will be able to exchange value with the rest of the Cosmos
Ecosystem through the Inter Blockchain Communication Protocol (IBC).

The name Kynno is extracted from the Old French and Greek word ‘Kynosoura’ meaning ‘the center of attraction. We aim for Kynno to become the center of attraction of everyone’s cryptoverse.

Through its revolutionary AI-powered, permissionless blockchain platform designed specifically to function as a secure digital infrastructure, Kynno strives at becoming a cornerstone of the blockchain ecosystem and aims at breaking the data silo paradigm. How ? By bridging with other ecosystems and blockchains to centralize your Web 3 identity, Assets, Data and Documents.

### Features

Here’s a glance at some of the key features of Kynno:

* Web3 and EVM compatibility
* High throughput via [Tendermint Core](https://github.com/tendermint/tendermint)
* Horizontal scalability via [IBC](https://cosmos.network/ibc)
* Ebded Nft minting service
* Fast transaction finality

Kynno enables these key features by:

* Implementing Tendermint Core's Application Blockchain Interface ([ABCI](https://docs.tendermint.com/master/spec/abci/)) to manage the blockchain
* Leveraging [modules](https://docs.cosmos.network/main/building-modules/intro.html) and other mechanisms implemented by the [Cosmos SDK](https://docs.cosmos.network/).
* Utilizing [`geth`](https://github.com/ethereum/go-ethereum) as a library to promote code reuse and improve maintainability.
* Exposing a fully compatible Web3 [JSON-RPC](./../../developers/json-rpc/server.md) layer for interacting with existing Ethereum clients and tooling ([Metamask](./../../users/wallets/metamask.md), [Remix](./../../developers/tools/remix.md), [Truffle](./../../developers/tools/truffle.md), etc).

The sum of these features allows developers to leverage existing Ethereum ecosystem tooling and
software to seamlessly deploy smart contracts which interact with the rest of the Cosmos
[ecosystem](https://cosmos.network/ecosystem)!
