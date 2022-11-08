<!--
order: 3
-->

# Accounts

This document describes the in-built accounts system of Kynno. {synopsis}

## Pre-requisite Readings

- [Cosmos SDK Accounts](https://docs.cosmos.network/main/basics/accounts.html) {prereq}
- [Ethereum Accounts](https://ethereum.org/en/whitepaper/#ethereum-accounts) {prereq}

## Kynno Accounts

Kynno defines its own custom `Account` type that uses Ethereum's ECDSA secp256k1 curve for keys. This
satisfies the [EIP84](https://github.com/ethereum/EIPs/issues/84) for full [BIP44](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki) paths.
The root HD path for Kynno-based accounts is `m/44'/60'/0'/0`.

+++ https://github.com/ahmedoubadi/ethermint/blob/main/types/account.pb.go#L28-L33

## Addresses and Public Keys

[BIP-0173](https://github.com/satoshilabs/slips/blob/master/slip-0173.md) defines a new format for segregated witness output addresses that contains a human-readable part that identifies the Bech32 usage. Kynno uses the following HRP (human readable prefix) as the base HRP:

| Network   | Mainnet | Testnet |
|-----------|---------|---------|
| Kynno     | `kynno` | `kynno` |

There are 3 main types of HRP for the `Addresses`/`PubKeys` available by default on Kynno:

- Addresses and Keys for **accounts**, which identify users (e.g. the sender of a `message`). They are derived using the **`eth_secp256k1`** curve.
- Addresses and Keys for **validator operators**, which identify the operators of validators. They are derived using the **`eth_secp256k1`** curve.
- Addresses and Keys for **consensus nodes**, which identify the validator nodes participating in consensus. They are derived using the **`ed25519`** curve.

|                    | Address bech32 Prefix | Pubkey bech32 Prefix | Curve           | Address byte length | Pubkey byte length |
|--------------------|-----------------------|----------------------|-----------------|---------------------|--------------------|
| Accounts           | `kynno`               | `kynnopub`           | `eth_secp256k1` | `20`                | `33` (compressed)  |
| Validator Operator | `kynnovaloper`        | `kynnovaloperpub`    | `eth_secp256k1` | `20`                | `33` (compressed)  |
| Consensus Nodes    | `kynnovalcons`        | `kynnovalconspub`    | `ed25519`       | `20`                | `32`               |

## Address formats for clients

`EthAccount` can be represented in both [Bech32](https://en.bitcoin.it/wiki/Bech32) (`kynno1...`) and hex (`0x...`) formats for Ethereum's Web3 tooling compatibility.

The Bech32 format is the default format for Cosmos-SDK queries and transactions through CLI and REST
clients. The hex format on the other hand, is the Ethereum `common.Address` representation of a
Cosmos `sdk.AccAddress`.

- **Address (Bech32)**: `kynno18g9p03n9ckasqgmwwkp6v56a6l9fhykqgufak3`
- **Address ([EIP55](https://eips.ethereum.org/EIPS/eip-55) Hex)**: `0x91defC7fE5603DFA8CC9B655cF5772459BF10c6f`
- **Compressed Public Key**: `{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"AsV5oddeB+hkByIJo/4lZiVUgXTzNfBPKC73cZ4K1YD2"}`

### Address conversion

The `kynnod debug addr <address>` can be used to convert an address between hex and bech32 formats. For example:

:::: tabs
::: tab Bech32

```bash
kynnod debug addr kynno18g9p03n9ckasqgmwwkp6v56a6l9fhykqgufak3
  Address: [58 10 23 198 101 197 187 0 35 110 117 131 166 83 93 215 202 155 146 192]
  Address (hex): 3A0A17C665C5BB00236E7583A6535DD7CA9B92C0
  Bech32 Acc: kynno18g9p03n9ckasqgmwwkp6v56a6l9fhykqgufak3
  Bech32 Val: kynnovaloper18g9p03n9ckasqgmwwkp6v56a6l9fhykqx86x9e
```

:::
::: tab Hex

```bash
kynnod debug addr 3A0A17C665C5BB00236E7583A6535DD7CA9B92C0
  Address: [58 10 23 198 101 197 187 0 35 110 117 131 166 83 93 215 202 155 146 192]
  Address (hex): 3A0A17C665C5BB00236E7583A6535DD7CA9B92C0
  Bech32 Acc: kynno18g9p03n9ckasqgmwwkp6v56a6l9fhykqgufak3
  Bech32 Val: kynnovaloper18g9p03n9ckasqgmwwkp6v56a6l9fhykqx86x9e
```

:::
::::

### Key output

::: tip
The Cosmos SDK Keyring output (i.e `kynnod keys`) only supports addresses and public keys in Bech32 format.
:::

We can use the `keys show` command of `kynnod` with the flag `--bech <type> (acc|val|cons)` to
obtain the addresses and keys as mentioned above,

:::: tabs
::: tab Account

```bash
kynnod keys show mykey --bech acc
- name: madou
  type: local
  address: kynno18g9p03n9ckasqgmwwkp6v56a6l9fhykqgufak3
  pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"A+tidH9dZ0hOtvQquBcpQjpQCYMAc3dgIs3sLh2tIteZ"}'
  mnemonic: ""
```

:::
::: tab Validator

```bash
kynnod keys show mykey --bech val
- name: madou
  type: local
  address: kynnovaloper18g9p03n9ckasqgmwwkp6v56a6l9fhykqx86x9e
  pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"A+tidH9dZ0hOtvQquBcpQjpQCYMAc3dgIs3sLh2tIteZ"}'
  mnemonic: ""
```

:::
::: tab Consensus

```bash
kynnod keys show mykey --bech cons
- name: madou
  type: local
  address: kynnovalcons18g9p03n9ckasqgmwwkp6v56a6l9fhykqj5f6fc
  pubkey: '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"A+tidH9dZ0hOtvQquBcpQjpQCYMAc3dgIs3sLh2tIteZ"}'
  mnemonic: ""
```

:::
::::

## Querying an Account

You can query an account address using the CLI, gRPC or

### Command Line Interface

```bash
# NOTE: the --output (-o) flag will define the output format in JSON or YAML (text)
kynnod q auth account $(kynnod keys show mykey -a) -o text
|
  '@type': /ethermint.types.v1.EthAccount
  base_account:
    account_number: "18"
    address: kynno18g9p03n9ckasqgmwwkp6v56a6l9fhykqgufak3
    pub_key:
      '@type': /ethermint.crypto.v1.ethsecp256k1.PubKey
      key: A+tidH9dZ0hOtvQquBcpQjpQCYMAc3dgIs3sLh2tIteZ
    sequence: "6"
  code_hash: 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470
```

### Cosmos gRPC and REST

``` bash
# GET /cosmos/auth/v1beta1/accounts/{address}
curl -X GET "http://localhost:10337/cosmos/auth/v1beta1/accounts/kynno18g9p03n9ckasqgmwwkp6v56a6l9fhykqgufak3" -H "accept: application/json"
```

### JSON-RPC

To retrieve the Ethereum hex address using Web3, use the JSON-RPC [`eth_accounts`](./../../developers/json-rpc/endpoints.md#eth-accounts) or [`personal_listAccounts`](./../../developers/json-rpc/endpoints.md#personal-listAccounts) endpoints:

```bash
# query against a local node
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545

curl -X POST --data '{"jsonrpc":"2.0","method":"personal_listAccounts","params":[],"id":1}' -H "Content-Type: application/json" http://localhost:8545
```
