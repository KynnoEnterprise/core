<!--
order: 3
-->

# Disk Usage Optimization

Customize the configuration settings to lower the disk requirements for your validator node {synopsis}

Blockchain database tends to grow over time, depending e.g. on block
speed and transaction amount. For Kynno, we are talking about close to
20GB of disk usage in first two weeks.

There are few configurations that can be done to reduce the required
disk usage quite significantly. Some of these changes take full effect
only when you do the configuration and start syncing from start with
them in use.

## Indexing

If you do not need to query transactions from the specific node, you can
disable indexing. On `config.toml` set

```toml
indexer = "null"
```

If you do this on already synced node, the collected index is not purged
automatically, you need to delete it manually. The index is located
under the database directory with name `data/tx_index.db/`.

## Configure pruning

By default every 500th state, and the last 100 states are kept. This
consumes a lot of disk space on long run, and can be optimized with
following custom configuration:

```toml
pruning = "custom"
pruning-keep-recent = "100"
pruning-keep-every = "0"
pruning-interval = "10"
```

Configuring `pruning-keep-recent = "0"` might sound tempting, but this
will risk database corruption if the `kynnod` is killed for any reason.
Thus, it is recommended to keep the few latest states.

## Logging

By default the logging level is set to `info`, and this produces a lot of
logs. This log level might be good when starting up to see that the
node starts syncing properly. However, after you see the syncing is
going smoothly, you can lower the log level to `warn` (or `error`). On
`config.toml` set the following

```toml
log_level = "warn"
```

Also ensure your log rotation is configured properly.
