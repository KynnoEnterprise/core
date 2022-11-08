<!--
order: 4
-->

# Backup

Learn how to backup your wallet's mnemonic and private key. {synopsis}

## Mnemonics

When you create a new key, you'll recieve a mnemonic phrase that can be used to restore that key. Backup the mnemonic phrase:

```bash
kynnod keys add mykey
{
  "name": "mykey"
  "type": "local"
  "address": "kynno1fegxqfl5l8ng9hmd8jz4v93qwc3q939dn9yu5n"
  "pubkey": '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"A+I6SbtSE0f7WhLPqF2UPniEHBiLlx+TxWHtHnr8/HuD"}'
  "mnemonic": ""
}

**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

# <24 word mnemonic phrase>
```

To restore the key:

```bash
$ kynnod keys add mykey-restored --recover
> Enter your bip39 mnemonic
banner genuine height east ghost oak toward reflect asset marble else explain foster car nest make van divide twice culture announce shuffle net peanut
{
  "name": "mykey-restored",
  "type": "local"
  "address": "kynno1fegxqfl5l8ng9hmd8jz4v93qwc3q939dn9yu5n"
  "pubkey": '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"A+I6SbtSE0f7WhLPqF2UPniEHBiLlx+TxWHtHnr8/HuD"}'
}
```

## Export Key

### Tendermint-Formatted Private Keys

To backup this type of key without the mnemonic phrase, do the following:

```bash
kynnod keys export mykey
Enter passphrase to decrypt your key:
Enter passphrase to encrypt the exported key:
-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 3015BD2DC9BC34D421A0FFDD38D3F4E5
type: eth_secp256k1

# <Tendermint private key>
-----END TENDERMINT PRIVATE KEY-----

$ echo "\
-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 3015BD2DC9BC34D421A0FFDD38D3F4E5
type: eth_secp256k1

# <Tendermint private key>
-----END TENDERMINT PRIVATE KEY-----" > mykey.export
```

### Ethereum-Formatted Private Keys

:::tip
**Note**: These types of keys are MetaMask-compatible.
:::

To backup this type of key without the mnemonic phrase, do the following:

```bash
kynnod keys unsafe-export-eth-key mykey > mykey.export
**WARNING** this is an unsafe way to export your unencrypted private key, are you sure? [y/N]: y
Enter keyring passphrase:
```

## Import Key

### Tendermint-Formatted Private Keys

```bash
$ kynnod keys import mykey-imported ./mykey.export
Enter passphrase to decrypt your key:
```

### Ethereum-Formatted Private Keys

```
$ kynnod keys unsafe-import-eth-key mykey-imported ./mykey.export
Enter passphrase to encrypt your key:
```

### Verification

Verify that your key has been restored using the following command:

```bash
$ kynnod keys list
[
  {
    "name": "mykey-restored",
    "type": "local"
    "address": "kynno1fegxqfl5l8ng9hmd8jz4v93qwc3q939dn9yu5n"
    "pubkey": '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"A+I6SbtSE0f7WhLPqF2UPniEHBiLlx+TxWHtHnr8/HuD"}'
  },
  {
    "name": "mykey",
    "type": "local"
    "address": "kynno1fegxqfl5l8ng9hmd8jz4v93qwc3q939dn9yu5n"
    "pubkey": '{"@type":"/ethermint.crypto.v1.ethsecp256k1.PubKey","key":"A+I6SbtSE0f7WhLPqF2UPniEHBiLlx+TxWHtHnr8/HuD"}'
  },
]
```
