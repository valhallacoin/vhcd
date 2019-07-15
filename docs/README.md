### Table of Contents
1. [About](#About)
2. [Getting Started](#GettingStarted)
    1. [Installation](#Installation)
    2. [Configuration](#Configuration)
    3. [Controlling and Querying vhcd via vhcctl](#VhcctlConfig)
    4. [Mining](#Mining)
3. [Help](#Help)
    1. [Network Configuration](#NetworkConfig)
    2. [Wallet](#Wallet)
4. [Contact](#Contact)
    1. [Community](#ContactCommunity)
5. [Developer Resources](#DeveloperResources)
    1. [Code Contribution Guidelines](#ContributionGuidelines)
    2. [JSON-RPC Reference](#JSONRPCReference)
    3. [Go Modules](#GoModules)
    4. [Module Hierarchy](#ModuleHierarchy)

<a name="About" />

### 1. About

vhcd is a full node Valhalla implementation written in [Go](http://golang.org),
and is licensed under the [copyfree](http://www.copyfree.org) ISC License.

This software is currently under active development.  It is extremely stable and
has been in production use since February 2016.

It also properly relays newly mined blocks, maintains a transaction pool, and
relays individual transactions that have not yet made it into a block.  It
ensures all individual transactions admitted to the pool follow the rules
required into the block chain and also includes the vast majority of the more
strict checks which filter transactions based on miner requirements ("standard"
transactions).

<a name="GettingStarted" />

### 2. Getting Started

<a name="Installation" />

**2.1 Installation**<br />

The first step is to install vhcd.  The installation instructions can be found
[here](https://github.com/valhallacoin/vhcd/tree/master/README.md#Installation).

<a name="Configuration" />

**2.2 Configuration**<br />

vhcd has a number of [configuration](http://godoc.org/github.com/valhallacoin/vhcd)
options, which can be viewed by running: `$ vhcd --help`.

<a name="VhcctlConfig" />

**2.3 Controlling and Querying vhcd via vhcctl**<br />

vhcctl is a command line utility that can be used to both control and query vhcd
via [RPC](http://www.wikipedia.org/wiki/Remote_procedure_call).  vhcd does
**not** enable its RPC server by default;  You must configure at minimum both an
RPC username and password or both an RPC limited username and password:

* vhcd.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
* vhcctl.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
```
OR
```
[Application Options]
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
For a list of available options, run: `$ vhcctl --help`

<a name="Mining" />

**2.4 Mining**<br />
vhcd supports the [getwork](https://github.com/valhallacoin/vhcd/tree/master/docs/json_rpc_api.md#getwork)
RPC.  The limited user cannot access this RPC.<br />

**1. Add the payment addresses with the `miningaddr` option.**<br />

```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
miningaddr=DsExampleAddress1
miningaddr=DsExampleAddress2
```

**2. Add vhcd's RPC TLS certificate to system Certificate Authority list.**<br />

`cgminer` uses [curl](http://curl.haxx.se/) to fetch data from the RPC server.
Since curl validates the certificate by default, we must install the `vhcd` RPC
certificate into the default system Certificate Authority list.

**Ubuntu**<br />

1. Copy rpc.cert to /usr/share/ca-certificates: `# cp /home/user/.vhcd/rpc.cert /usr/share/ca-certificates/vhcd.crt`<br />
2. Add vhcd.crt to /etc/ca-certificates.conf: `# echo vhcd.crt >> /etc/ca-certificates.conf`<br />
3. Update the CA certificate list: `# update-ca-certificates`<br />

**3. Set your mining software url to use https.**<br />

`$ cgminer -o https://127.0.0.1:9209 -u rpcuser -p rpcpassword`

<a name="Help" />

### 3. Help

<a name="NetworkConfig" />

**3.1 Network Configuration**<br />
* [What Ports Are Used by Default?](https://github.com/valhallacoin/vhcd/tree/master/docs/default_ports.md)
* [How To Listen on Specific Interfaces](https://github.com/valhallacoin/vhcd/tree/master/docs/configure_peer_server_listen_interfaces.md)
* [How To Configure RPC Server to Listen on Specific Interfaces](https://github.com/valhallacoin/vhcd/tree/master/docs/configure_rpc_server_listen_interfaces.md)
* [Configuring vhcd with Tor](https://github.com/valhallacoin/vhcd/tree/master/docs/configuring_tor.md)

<a name="Wallet" />

**3.2 Wallet**<br />

vhcd was intentionally developed without an integrated wallet for security
reasons.  Please see [vhcwallet](https://github.com/valhallacoin/vhcwallet) for more
information.

<a name="Contact" />

### 4. Contact

<a name="ContactCommunity" />

**4.1 Community**<br />

If you have any further questions you can find us at:

https://valhallacoin.org/community

<a name="DeveloperResources" />

### 5. Developer Resources

<a name="ContributionGuidelines" />

**5.1 Code Contribution Guidelines**

* [Code Contribution Guidelines](https://github.com/valhallacoin/vhcd/tree/master/docs/code_contribution_guidelines.md)

<a name="JSONRPCReference" />

**5.2 JSON-RPC Reference**

* [JSON-RPC Reference](https://github.com/valhallacoin/vhcd/tree/master/docs/json_rpc_api.md)
    * [RPC Examples](https://github.com/valhallacoin/vhcd/tree/master/docs/json_rpc_api.md#ExampleCode)

<a name="GoModules" />

**5.3 Go Modules**

The following versioned modules are provided by vhcd repository:

* [rpcclient](https://github.com/valhallacoin/vhcd/tree/master/rpcclient) - Implements
  a robust and easy to use Websocket-enabled Valhalla JSON-RPC client
* [vhcjson](https://github.com/valhallacoin/vhcd/tree/master/vhcjson) - Provides an
  extensive API for the underlying JSON-RPC command and return values
* [wire](https://github.com/valhallacoin/vhcd/tree/master/wire) - Implements the
  Valhalla wire protocol
* [peer](https://github.com/valhallacoin/vhcd/tree/master/peer) - Provides a common
  base for creating and managing Valhalla network peers
* [blockchain](https://github.com/valhallacoin/vhcd/tree/master/blockchain) -
  Implements Valhalla block handling and chain selection rules
  * [stake](https://github.com/valhallacoin/vhcd/tree/master/blockchain/stake) -
    Provides an API for working with stake transactions and other portions
    related to the Proof-of-Stake (PoS) system
* [txscript](https://github.com/valhallacoin/vhcd/tree/master/txscript) -
  Implements the Valhalla transaction scripting language
* [vhcec](https://github.com/valhallacoin/vhcd/tree/master/vhcec) - Provides constants
  for the supported cryptographic signatures supported by Valhalla scripts
  * [secp256k1](https://github.com/valhallacoin/vhcd/tree/master/vhcec/secp256k1) -
    Implements the secp256k1 elliptic curve
  * [edwards](https://github.com/valhallacoin/vhcd/tree/master/vhcec/edwards) -
    Implements the edwards25519 twisted Edwards curve
* [database](https://github.com/valhallacoin/vhcd/tree/master/database) -
  Provides a database interface for the Valhalla block chain
* [mempool](https://github.com/valhallacoin/vhcd/tree/master/mempool) - Provides a
  policy-enforced pool of unmined Valhalla transactions
* [vhcutil](https://github.com/valhallacoin/vhcd/tree/master/vhcutil) - Provides
  Valhalla-specific convenience functions and types
* [chaincfg](https://github.com/valhallacoin/vhcd/tree/master/chaincfg) - Defines
  chain configuration parameters for the standard Valhalla networks and allows
  callers to define their own custom Valhalla networks for testing puproses
  * [chainhash](https://github.com/valhallacoin/vhcd/tree/master/chaincfg/chainhash) -
    Provides a generic hash type and associated functions that allows the
    specific hash algorithm to be abstracted
* [certgen](https://github.com/valhallacoin/vhcd/tree/master/certgen) - Provides a
  function for creating a new TLS certificate key pair, typically used for
  encrypting RPC and websocket communications
* [addrmgr](https://github.com/valhallacoin/vhcd/tree/master/addrmgr) - Provides a
  concurrency safe Valhalla network address manager
* [connmgr](https://github.com/valhallacoin/vhcd/tree/master/connmgr) - Implements a
  generic Valhalla network connection manager
* [hdkeychain](https://github.com/valhallacoin/vhcd/tree/master/hdkeychain) - Provides
  an API for working with  Valhalla hierarchical deterministic extended keys
* [gcs](https://github.com/valhallacoin/vhcd/tree/master/gcs) - Provides an API for
  building and using Golomb-coded set filters useful for light clients such as
  SPV wallets
* [fees](https://github.com/valhallacoin/vhcd/tree/master/fees) - Provides methods for
  tracking and estimating fee rates for new transactions to be mined into the
  network

<a name="ModuleHierarchy" />

**5.4 Module Hierarchy**

The following diagram shows an overview of the hierarchy for the modules
provided by the vhcd repository.

![Module Hierarchy](./assets/module_hierarchy.svg)