The repository hosts the node software for ZetaChain. 

There are two important binaries produced: `zetacored` and `zetaclientd`.
`zetacored` is the ZetaChain node daemon (full node), and
`zetaclientd` is special companion daemon for interacting
with external/foreign blockchains such as Ethereum, Binance
Smart Chain, Polygon, or Bitcoin (unless you are special
validators, you don't need to run `zetaclientd`).

To compile and install both binaries under your `$GOPATH/bin`:

```bash
$ make install
```

To setup a local development/testing environment, see
[ZetaChain private net setup](./contrib/localnet/README.md)