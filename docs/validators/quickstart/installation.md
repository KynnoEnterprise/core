<!--
order: 1
-->

# Installation

Build and install the Kynno binaries from source or using Docker. {synopsis}

## Pre-requisites

- [Install Go 1.19+](https://golang.org/dl/) {prereq}
- [Install jq](https://stedolan.github.io/jq/download/) {prereq}

## Install additional Linux dependencies
Use the following command to install additional Linux dependencies.
```bash
    apt-get update \
    && apt-get install -y --no-install-recommends \
    tzdata \
    ca-certificates \
    build-essential \
    pkg-config \
    cmake
```
## Install Go

::: warning
Kynno is built using [Go](https://golang.org/dl/) version `1.19+`
:::

```bash
go version
```

:::tip
If the `kynnod: command not found` error message is returned, confirm that your [`GOPATH`](https://golang.org/doc/gopath_code#GOPATH) is correctly configured by running the following command:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

:::

## Install Binaries

::: tip
The latest {{ $themeConfig.project.name }} [version](https://github.com/kynnoenterprise/core/releases) is `{{ $themeConfig.project.binary }} {{ $themeConfig.project.latest_version }}`
:::

### GitHub

Clone and build {{ $themeConfig.project.name }} using `git`:

```bash
git clone https://github.com/kynnoenterprise/core.git
cd core
make install
```

Check that the `{{ $themeConfig.project.binary }}` binaries have been successfully installed:

```bash
kynnod version
```


### Releases

You can also download a specific release available on the {{ $themeConfig.project.name }} [repository](https://github.com/kynnoenterprise/core/releases) or via command line:

```bash
go install github.com/kynnoenterprise/core@latest
```
