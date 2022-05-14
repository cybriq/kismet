# proto
Protocol for distributed systems governance

## About

In distributed systems, there is a not well understood distinction between leader selection and consensus. 

The great innovation of Bitcoin was not its consensus, as blockchains as authenticity certification are quite an old technology, but rather, the use of proof of work to randomly select a leader to propose a new block.

Proto is a distributed system design that takes this idea of a distributed lottery for issuing a limited supply of tokens, and instead of a direct monetary reward, the tokens are time limited right to perform leadership tasks for the network.

This produces a very small, simple blockchain which only concerns itself with authenticating the winners of these tokens, who can then trade them for a position in a distributed system.

The first application of these tokens will be in granting a limited number of spots in a validator set queue to run a pBFT based tendermint type chain, where validator membership is governed by these limited issue tokens. 

The tokens obviously will be exchangeable, so mining the tokens can be separated from operating a consensus node on this initial blockchain system.

Further, the core governance token chain will have multiple token types, that will be added to limit and distribute membership in other ways. 

- validator tokens grant a limited number of blocks to be minted on the main chain
- congress tokens grant a limited number of votes on proposals for chain governance to form definite quora
- proposal tokens, which grant the right to make a proposal to the congress

Standard fungible tokens are issued to validators on a fixed supply expansion percentage, which is locked and cannot be changed, ever, at a rate of 5% per year, compounded on a block by block basis, which is one second per block. In theory this could be changed, but long centuries of experience proves that any power to change the money supply is always abused. 

As such, the leader of the development team for Proto will be a benevolent dictatorship for life position, and any proposal involving changing issuance rate will be vetoed by the team as not ever going to do, and this one power to say no to this exact type of proposal is the only privilege.

The proto chain runs independently from the pBFT chain, and refers to it. Proposals are tied to IPFS file repositories that contain the proposal text as well as all associated code that implements it.
