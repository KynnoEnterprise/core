<!--
order: 3
-->

# Deterministic Builds

Build the `kynnod` binary deterministically using Docker. {synopsis}

## Pre-requisites

- [Install Docker](https://docs.docker.com/get-docker/) {prereq}

## Introduction

The [Tendermint rbuilder Docker image](https://github.com/tendermint/images/tree/master/rbuilder) provides a deterministic build environment that is used to build Cosmos SDK applications. It provides a way to be reasonably sure that the executables are really built from the git source. It also makes sure that the same, tested dependencies are used and statically built into the executable.

::: tip
All the following instructions have been tested on *Ubuntu 18.04.2 LTS* with *Docker 20.10.2*.
:::

## Build with Docker

Clone `kynno`:

``` bash
git clone git@github.com:kynnoenterprise/core.git
```

Checkout the commit, branch, or release tag you want to build (eg `v1.0.0`):

```bash
cd code/
git checkout v1.0.0
```

The buildsystem supports and produces binaries for the following architectures:

* **linux/amd64**

Run the following command to launch a build for all supported architectures:

```bash
make distclean build-reproducible
```

The build system generates both the binaries and deterministic build report in the `artifacts` directory.
The `artifacts/build_report` file contains the list of the build artifacts and their respective checksums, and can be used to verify
build sanity. 