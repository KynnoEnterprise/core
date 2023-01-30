<!--
order: 1
-->

# Faucet

Check how to obtain testnet tokens from the Kynno faucet website {synopsis}

The Kynno Testnet Faucet distributes small amounts of {{ $themeConfig.project.testnet_denom }} to anyone who can provide a valid testnet address for free. To Request funds follow the instructions on this page.

::: tip
Follow the [Metamask](../../users/wallets/metamask.md) or [Keyring](../../users/keys/keyring.md) guides for more info on how to setup your wallet account.
:::

## Request Testnet tokens

<!-- markdown-link-check-disable-next-line -->
Once you are signed in to the Metamask extension, visit the [Faucet](https://faucet.kynno.dev/) to request tokens for the testnet.


After approval, you can see a transaction confirmation informing you that {{ $themeConfig.project.testnet_denom }} have been successfully transferred to your [kynno address](../../users/technical_concepts/accounts.md#address-formats-for-clients) on the testnet.

::: warning
**Note**: only Ethereum compatible addresses (i.e `eth_secp256k1` keys) are supported on Kynno.
:::

Alternatively you can also fill in your address on the input field in Bech32 (`kynno1...`) or Hex (`0x...`) format.

::: warning
If you use your Bech32 address, make sure you input the [account address](../../users/technical_concepts/accounts.md#addresses-and-public-keys) (`kynno1...`) and **NOT** the validator operator address (`kynnovaloper1...`)
:::

![faucet site](../../img/faucet_web_page.png)

View your account balance either by clicking on the Metamask extension or by using the [Testnet Explorer](https://evm.kynno.dev).

## Rate limits

::: tip
All addresses **Can** request testnet token once per day.
:::

To prevent the faucet account from draining the available funds, the Kynno testnet faucet imposes a maximum number of requests for a period of time. By default, the faucet service accepts 1 request per day per address. You can request {{ $themeConfig.project.testnet_denom }} from the faucet for each address only once every 24h. If you try to request multiple times within the 24h cooldown phase, no transaction will be initiated. Please try again in 24 hours.

## Amount

For each request, the faucet transfers 1 {{ $themeConfig.project.testnet_denom }} to the given address.

## Faucet Addresses

The public faucet addresses for the testnet are:

- **Hex**: [`0x5e33d4a30cad41937b0f7e893ca730c7730a78c5`](https://evm.kynno.dev/address/0x5e33d4a30cad41937b0f7e893ca730c7730a78c5)
- **Bech32**: [`kynno1tceafgcv44qex7c006ynefescaes57x9cc7z48`](https://evm.kynno.dev/address/0x5e33d4a30cad41937b0f7e893ca730c7730a78c5)
