# Kismet Whitepaper

David Vennik <david@cybriq.systems>

June 2022

**Abstract.** Leadership and consensus are two distinct concepts in distributed systems. Proof of Work picks a leader, same as Proof of Stake, but in the Bitcoin design, it dictates the timing of blocks, and the emission of tokens all in one. While this has advantages of simplicity it makes the distributed system it creates somewhat inadequate for transaction processing. Proof of Stake is nothing more than distributed systems replication of plutocracy. Kismet brings the idea of Proof of Work for leadership while using a proven, fast Classical Consensus for the actual block scheduling and timing of transaction processing.



## Introduction

Currently in the field of public blockchain based financial ledgers there is the Proof of Work chains with random timing and probabalistic finality, and on the other hand, the Proof of Stake classical consensus chains with rapid or instant finality. There is a wide gulf between these two systems as regards to satisfying the requirements for a financial ledger for processing payments. 

"Proof of Work" chains have the problem that you can't be sure about transactions staying final until several hours later, but protect against monopolisation, which has happened many times to "Proof of Stake" blockchains.

Underpinning both models is the idea of distributing risk amongst a group of entities who get short term leadership of the system. This is a little similar to the idea of term limits in Democracy.

Proof of Stake chains sacrifice the breadth of this hedge against corruption for fast processing. Proof of Work sacrifices regular processing times for a broader hedge against corruption.

## Proof of Work for Leadership with Classical Consensus

What Kismet proposes is creating a chain purely for recording the winners of a Proof of Work lottery for various leadership roles, and accumulating these ahead of time to produce a future queue of validators for a classical consensus based on Tendermint's variant of the Practical Byzantine Fault Tolerance protocol.

Winning a validator slot enables the winner to join the pBFT chain validator queue, by applying to the current validator set for membership, who then agree to place the winner's right to membership into schedule for new blocks.

Standard pBFT chains use a new block generated seeding for random selection, combined with a weighting algorithm to reward bigger stakes with more leadership rewards.

The pBFT chains that use Proof of Stake to select leadership for rounds of the consensus have a problem that through mischief an outsized quantity of tokens can be acquired, and so they have to add further complexity with a rule called "slashing" that entails peers being able to unlock and divert stake to a governance treasury or burning the stake in order to disincentivise such mischief.

Proof of Work doesn't have this weakness because acquiring mining power cannot be cheated. Thus there is no need for this complication.

The problem gets even more acute in the context of Decentralised Finance, where one need not merely rob someone to gain a leadership position, but to defraud a credit market, or in other words, rob many someones. 

Leadership in distributed systems quite simply becomes more and more vulnerable to gaming the less randomly it is given out.

## Block Rewards

Block rewards will be issued on an exponential decay basis, and will continue until the block reward rounds down to zero, at which point it will just be zero. This will aim at around 25 years until zero. After this transaction fees will be the sole means of rewarding miner/validators.

## Lightning Network Interoperability

Blocks will be limited to a smaller size than normally used with Tendermint Consensus as we intend to implement Lightning style offchain transactions and keep the chain small enough that the pool of miner/validators can remain large.

In order to eliminate the need to alter this, the block size limit will slowly grow with time in concert with the reduction in block reward, and continue to grow at this rate. This rate will be chosen based on the historical trend of the equal cost for a given amount of storage. Something like doubling every 4 years.

Bitcoin compatible UXTO transaction scripting will be used in order to have direct interoperability with Lightning protocol. Lightning functionality should be implemented as soon as possible after the core system is up and running.

Block timing will be slower, at 15 seconds, or 4 per minute. This is sufficiently fast for use with a forum system that stores content on IPFS networks, funded by transaction fees for posting these links.

## Atomic Swaps

Once Lightning functionality is implemented, the chain should have atomic swaps implemented with selected counterpart blockchains. This is essential to eliminate the problem of centralised exchanges and their relationship with governments.

## Governance

Changing the rules of the consensus requires upgrading by miners. We are adopting the same conservative change averse policy as bitcoin. We will periodically make a major version upgrade, and if miners don't install it before the fork deadline, it doesn't happen. Simple as that. Then we have to go back to the drawing board, after more consultation. A successful hard fork will come from consultation with our miner-stakeholders.

## Conclusion

Like all good games, the best rules are the most concise. As such, this is the entire Whitepaper. 

We are just bolting on Tendermint consensus to a leadership only chain based on Bitcoin's Nakamoto Consensus. The idea is to create another option for Cosmos based chains to not use proven vulnerable Proof of Stake leadership selection.