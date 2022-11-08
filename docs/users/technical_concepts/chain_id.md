<!--
order: 3
-->

# Chain ID

Learn about the Kynno chain-id format {synopsis}

## Official Chain IDs

::: tip
**NOTE**: The latest Chain ID (i.e highest Version Number) is the latest version of the software and mainnet.
:::

:::: tabs
::: tab Mainnet

| Name                                            | Chain ID                                      | Identifier | EIP155 Number                         | Version Number                              |
| ----------------------------------------------- | --------------------------------------------- | ---------- | ------------------------------------- | ------------------------------------------- |
| Kynno 1                                         | `kynno_9701-1` | `kynno`    | `9701` | `1`                                         |
:::
::: tab Testnets

| Name                              | Chain ID                                              | Identifier | EIP155 Number                                 | Version Number                                      |
| --------------------------------- | ----------------------------------------------------- | ---------- | --------------------------------------------- | --------------------------------------------------- |
| Kynno Public Testnet              | `kynno_9700-1` | `kynno`    | `9700` | `1` |

:::
::::


## The Chain Identifier

Every chain must have a unique identifier or `chain-id`. Tendermint requires each application to
define its own `chain-id` in the [genesis.json fields](https://docs.tendermint.com/master/spec/core/genesis.html#genesis-fields). However, in order to comply with both EIP155 and Cosmos standard for chain upgrades, Kynno-compatible chains must implement a special structure for their chain identifiers.

## Structure

The Kynno Chain ID contains 3 main components

- **Identifier**: Unstructured string that defines the name of the application.
- **EIP155 Number**: Immutable [EIP155](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md) `CHAIN_ID` that defines the replay attack protection number.
- **Version Number**: Is the version number (always positive) that the chain is currently running.
This number **MUST** be incremented every time the chain is upgraded or forked in order to avoid network or consensus errors.

### Format

The format for specifying and Kynno compatible chain-id in genesis is the following:

```bash
{identifier}_{EIP155}-{version}
```

The following table provides an example where the second row corresponds to an upgrade from the first one:

| ChainID        | Identifier | EIP155 Number | Version Number |
| -------------- | ---------- | ------------- | -------------- |
| `kynno_9700-1` | kynno      | 9700          | 1              |
| `kynno_9700-2` | kynno      | 9700          | 2              |
| `...`          | ...        | ...           | ...            |
| `kynno_9700-N` | kynno      | 9700          | N              |
