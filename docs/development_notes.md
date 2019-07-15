Development notes
-----------------

### Runing a Simnet

See [Simnet](https://docs.decred.org/advanced/simnet/).

### Running a local Testnet

Daemon 1:

    vhcd --testnet --nodnsseed --generate \
         --miningaddr=Tsp3KKyVdCPfrqzs1TgNNvrNMcn9fFy63kA

Daemon 2:

    vhcd --testnet --nodnsseed --generate \
         --nolisten --norpc --appdata=/home/user/.vhcd2 --connect=localhost \
         --miningaddr=Tsoc17inurwPr3X1TUJmwQe2q6pP8PRmRtu

Wallet:

    vhcwallet --testnet

The two nodes have to talk to be connected to each other before they
start mining.
