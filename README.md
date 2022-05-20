![logo](doc/logoshadow.svg)

# κισμέτ

kismet

Protocol for distributed systems governance

> Live by the dice, but never go all in - David Vennik

## Documentation

- [tokens](doc/tokens.md) - about kismet tokens
- [consensus](doc/consensus.md) - chain selection rules

## About

In distributed systems, there is a not well understood distinction between leader selection and consensus. 

The great innovation of Bitcoin was not its consensus. The consensus is simply that the heaviest chain (sum of hashes of the sequence of blocks being the smallest numerical value) is the consensus, canonical chain. 

The great innovation was using proof of work to decide a leader from an indefinite number of network participants of unlimited size, thus making it completely public and open, unlike the classical distributed systems designs, which require foreknowledge of the leader for a task before it is done.

Kismet is a distributed system design that takes this idea of a distributed lottery for issuing a limited supply of tokens, and instead of a direct monetary reward, the tokens are a limited right to perform leadership tasks for the network.

This produces a very small, simple blockchain which only concerns itself with authenticating the winners of these tokens.

The first application of these tokens will be in granting a limited number of spots in a validator set queue to run a pBFT based Tendermint type chain, where validator membership is governed by these limited issue tokens. 

These tokens could be exchanged, but to do so, buyers would have to trust the seller to not use the tokens in parallel. Thus, they will effectively only be exchangeable with trustworthy individuals and thus will fail to achieve a market price. There will never be a key exchange process.

The core governance token chain will have multiple token types, that will be added to limit and distribute membership in other ways. 

- validator tokens grant a limited number of blocks to be minted on the main chain. These tokens expire in a fixed amount of time after issuance in order to guarantee progress of the main chain, and are supplied at double the needed rate in order to ensure chain progress against a (remotely) possible denial of service attack.
- proposal tokens, which grant the right to make a proposal to the congress
- congress tokens grant a limited number of votes on a specific proposal, can only be mined in a limited number after a proposal is used on the pBFT ledger

Standard fungible tokens are issued to validators on a fixed supply expansion percentage, which is locked and cannot be changed, ever, at a rate of 5% per year, compounded on a block by block basis, which is one second per block. In theory this could be changed, but long centuries of experience proves that any power to change the money supply is always abused. 

As such, the leader of the development team for Kismet will be a benevolent dictatorship for life (BDFL) position, and any proposal involving changing issuance rate will be vetoed by the team as not ever going to do, and this one power to say no to this exact type of proposal is the only privilege. To enable this, the BDFL role includes a key pair that can be used to veto any proposal. It does not grant any other right than stopping proposals.

This privilege will also cover the prohibition of the addition of a key exchange process to the token chain. People can sell the tokens but cannot be sure the seller will not keep the secret key and use it. These are position statements of the BDFL. 

In addition, the veto token can issue delegates, in order to distribute the veto power to account for a growth in the amount of governance needed.

The Kismet chain runs independently from the pBFT chain, and refers to it. Proposals are tied to IPFS file repositories that contain the proposal text as well as all associated code that implements it.

The block hash target is not the blake3 hash used for indexing, but uses a Big integer multiplication/division cycle data expansion to make ASIC optimisation impossible.

