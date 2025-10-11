# Installing Aerium

## Requirements

You need to make sure you have installed [Git](https://git-scm.com/downloads)
and [Go 1.21 or higher](https://golang.org/) on your machine.
If you want to install a GUI application, make sure you have installed
[GTK+3](https://www.gtk.org/docs/getting-started/) as well.

## Compiling the code

Follow these steps to compile and build Aerium:

```bash
git clone https://github.com/aerium-network/aerium.git
cd aerium
make build
```

This will be compile `aerium-daemon` and `aerium-wallet` on your machine.
Make sure Aerium is properly compiled and installed on your machine:

```bash
cd build
./aerium-daemon version
```

If you want to compile the GUI application, run this command in the root folder:

```bash
make build_gui
```

To run the tests, use this command:

```bash
make test
```

This may take several minutes to finish.

## What is aerium-daemon?

`aerium-daemon` is a full node implementation of Aerium blockchain.
You can use `aerium-daemon` to run a full node:

```bash
./aerium-daemon init  -w=<working_dir>
./aerium-daemon start -w=<working_dir>
```

### Testnet

To join the TestNet, first you need to initialize your node
and then start the node:

```bash
./aerium-daemon init  -w=<working_dir> --testnet
./aerium-daemon start -w=<working_dir>
```

### Local net

You can create a local node to set up a local network for development purposes on your machine:

 ```bash
 ./aerium-daemon init  -w=<working_dir> --localnet
 ./aerium-daemon start -w=<working_dir>
 ```

## What is aerium-wallet?

Aerium wallet is a native wallet in the Aerium blockchain that lets users easily manage
their accounts on the Aerium blockchain.

### Getting started

To create a new wallet, run this command. The wallet will be encrypted by the
provided password.

```bash
./aerium-wallet --path ~/aerium/wallets/wallet_1 create
```

You can create a new address like this:

```bash
./aerium-wallet --path ~/aerium/wallets/wallet_1 address new
```

A list of addresses is available with this command:

```bash
./aerium-wallet --path ~/aerium/wallets/wallet_1 address all
```

To obtain the public key of an address, run this command:

```bash
./aerium-wallet --path ~/aerium/wallets/wallet_1 address pub <ADDRESS>
```

To publish a transaction, use the tx subcommand.
For example, to publish a bond transaction:

```bash
./aerium-wallet --path ~/aerium/wallets/wallet_1 tx bond <FROM> <TO> <AMOUNT>
```

You can recover a wallet if you have the seed phrase.

```bash
./aerium-wallet --path ~/aerium/wallets/wallet_2 recover
```

## Docker

You can run Aerium using a Docker file. Please make sure you have installed
[docker](https://docs.docker.com/engine/install/) on your machine.

Pull the Docker from Docker Hub:

```bash
docker pull aeriumdev/aerium:main
```

Let's create a working directory at `~/aerium/testnet` for the testnet:

```bash
docker run -it --rm -v ~/aerium/testnet:/root/aerium aeriumdev/aerium:main aerium-daemon init --testnet
```

Now we can run Aerium and join the testnet:

```bash
docker run -it -v ~/aerium/testnet:/root/aerium -p 8080:8080 -p 19933:19933 --name aerium-testnet aeriumdev/aerium:main aerium-daemon start
```

check "[http://localhost:8080](http://localhost:8080)" for the list of APIs.

Also you can stop/start docker:

```bash
docker start aerium-testnet
docker stop aerium-testnet
```

Or check the logs:

```bash
docker logs aerium-testnet --tail 1000 -f
```
