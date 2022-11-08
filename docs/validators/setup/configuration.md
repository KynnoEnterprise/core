<!--
order: 2
-->

# Configuration

## Block Time

The timeout-commit value in the node config defines how long we wait after committing a block, before starting on the new height (this gives us a chance to receive some more pre-commits, even though we already have +2/3). The current default value is `"1s"`.

::: tip
**Note**: by default this is handled automatically by the server when initializing the node.
Validators will need to ensure their local node configurations in order to speed up the network to ~2s block times.
:::

```toml
# In kynnod/config/config.toml

#######################################################
###         Consensus Configuration Options         ###
#######################################################
[consensus]

### ... 

# How long we wait after committing a block, before starting on the new
# height (this gives us a chance to receive some more precommits, even
# though we already have +2/3).
timeout_commit = "1s"
```

## Peers

In `kynnod/config/config.toml` you can set your peers.

See the [Add persistent peers section](../testnet.md#add-persistent-peers) in our docs for an automated method, but field should look something like a comma separated string of peers (do not copy this, just an example):

```bash
persistent_peers = "5576d0d50761fe81ccdf88e06031a01bc8643d51@199.205.108.13:24656,13e850d14610f966cdf4fc2fd54f6dc35c7f4bf4@186.6.51.23:26656,38eb4984f89899a5cf54f04a79b356f15681bb78@52.142.132.562:26656,59c4351009215ef5s8ri4bd5ee4324926a5a11aa@62.13.117.32:26656"
```

### Sharing your Peer

You can see and share your peer with the `tendermint show-node-id` command

```bash
kynnod tendermint show-node-id
ac29d2d50a4514545048a4481d16c12f59b2e58b
```

- **Peer Format**: `node-id@ip:port`
- **Example**: `ad54d21f516885424548a4481d16c12f59b2e58b@133.171.245.175:26656`

### Healthy peers

If you are relying on just seed node and no persistent peers or a low amount of them, please increase the following params in the `config.toml`:

```bash
# Maximum number of inbound peers
max_num_inbound_peers = 120

# Maximum number of outbound peers to connect to, excluding persistent peers
max_num_outbound_peers = 60
```
