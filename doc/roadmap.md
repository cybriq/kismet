# Roadmap

In order to accelerate time to proof of concept, the designs described in the [White Paper](./whitepaper.md) will be implemented one by one in their order.

1.   Implement proof of work chain for validator tokens, integrate as a module into a Cosmos SDK chain creating a validator schedule, account model address and transfers and simple multisig transfers. Proof of work derived from the [ParallelCoin](https://github.com/cybriq/p9) work.
2.   Create a multi platform wallet and mining control interface based on the prior work done for ParallelCoin.
3.   Launch mainnet.
4.   Add proposals and congress tokens for creating governance proposal but without a notion of treasury or any kind of DAO. DAO could be an subchain, but should not be part of the core ledger for simplicity.