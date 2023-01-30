<!--
order: 1
-->

# Wallet Integration

Learn how to properly integrate [Metamask](https://metamask.io/) with a dApp on Kynno. {synopsis}

:::tip
**Note**: want to learn more about wallet integration beyond what's covered here? Check out both the [MetaMask Wallet documentation](https://docs.metamask.io/guide/).
:::

## Pre-requisite Readings

- [MetaMask documentation](https://docs.metamask.io/guide/) {prereq}

## Implementation Checklist

The integration implementation checklist for dApp developers consists of three categories:

1. [Frontend features](#frontend)
2. [Transactions and wallet interactions](#transactions)
3. [Client-side provider](#connections)

### Frontend

Make sure to create a wallet-connection button for Metamask on the frontend of the application. 

### Transactions

Developers enabling transactions on their dApp have to [determine wallet type](#determining-wallet-type) of the user, [create the transaction](#create-the-transaction), [request signatures](#sign-and-broadcast-the-transaction) from the corresponding wallet, and finally [broadcast the transaction](#sign-and-broadcast-the-transaction) to the network.

#### Determining Wallet Type

Developers should determine whether users are using MetaMask or another wallet. Whether MetaMask or Another is installed on the user device can be determined by checking the corresponding `window.ethereum` value.

- **For MetaMask**: `await window.ethereum.enable(chainId);`

If `window.ethereum` returns `undefined` after `document.load`, then MetaMask (or another wallet) is not installed. There are several ways to wait for the load event to check the status: for instance, developers can register functions to `window.onload`, or they can track the document's ready state through the document event listener.

After the user's wallet type has been determined, developers can proceed with creating, signing, and sending transactions.

#### Create the Transaction

:::tip
**Note**: The example below uses the Kynno Testnet `chainID`. For more info, check the Kynno Chain IDs reference document [here](../../users/technical_concepts/chain_id.md).
:::

Developers can create `MsgSend` transactions using the [kynno-sdk-js](../libraries/kynnojs.md) library.

```js
import {newClient as kynnoSdkClient} from 'kynno-sdk-js';
import { createMessageSend } from @tharsis/transactions
const ClientConfig = {
    api :'http://127.0.0.1:1317',
    node: 'http://127.0.0.1:26657',
    chainNetwork: 0,
    chainId: 'kynno_9700-1',
    gas: '20000000',
    fee: { denom: 'akynno', amount: '20000' },
}
const client = kynnoSdkClient(ClientConfig).withRpcConfig({timeout: 15000})

const sender = {
    accountAddress: 'kynno1m63tc4gflpgx904fu2vmekdex652tp4vvvhdqy',
    sequence: 1,
    accountNumber: 9,
    pubkey: 'AgTw+4v0daIrxsNSW4FcQ+IoingPseFwHO1DnssyoOqZ',
}

const memo = ''

const chain = {
    chainId: 9700,
    cosmosChainId: 'kynno_9700-1',
}
const fee = {
    amount: '20',
    denom: 'akynno',
    gas: '200000',
}

const params = {
    destinationAddress: 'kynno1mx9nqk5agvlsvt2yc8259nwztmxq7zjq50mxkp',
    amount: (100*(10**18)).toString(),
    denom: 'akynno',
}
const msg = client.transaction._createMessageSend(
    chain,sender,fee,memo,params
)


// msg.signDirect is the transaction in Keplr format
// msg.legacyAmino is the transaction with legacy amino
// msg.eipToSign is the EIP712 data to sign with metamask
```

#### Sign and Broadcast the Transaction

:::tip
**Note**: The example below uses an Kynno Testnet [RPC node](../connect.md#public-available-endpoints).
:::

<!-- textlint-disable -->
After creating the transaction, developers need to send the payload to the appropriate wallet to be signed ([`msg.signDirect`](https://docs.keplr.app/api/#sign-direct-protobuf) is the transaction in Keplr format, and `msg.eipToSign` is the [`EIP712`](https://eips.ethereum.org/EIPS/eip-712) data to sign with MetaMask).

With the signature, we add a Web3Extension to the transaction and broadcast it to the Kynno node.

<!-- textlint-enable -->
```js
// Note that this example is for MetaMask, using kynno-sdk-js

// Follow the previous code block to generate the msg object

// Init Metamask
await window.ethereum.enable();

// Request the signature
let signature = await window.ethereum.request({
    method: 'eth_signTypedData_v4',
    params: [client.utils.toEth(sender.accountAddress), JSON.stringify(msg.eipToSign)],
});

// The chain and sender objects are the same as the previous example
let extension = client.transaction._signatureToWeb3Extension(chain, sender, signature)

// Create the txRaw
let rawTx = client.transaction._createTxRawEIP712(msg.legacyAmino.body, msg.legacyAmino.authInfo, extension)

let response = await client.transaction._broadcastTx(rawTx)
```

### Connections

For Ethereum RPC, Kynno gRPC, and/or REST queries, dApp developers should implement providers client-side, and store RPC details in the environment variable as secrets.
