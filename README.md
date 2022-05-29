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

### CPU based Proof of Work using Long Division

It is of course important what kind of work is used to perform this leadership lottery. The original Bitcoin proof of work was taken directly from HashCash and also used in Bitmessage, of which neither two you ever heard of probably. These protocols use a widely used, performance optimized cryptographic hash collision discovery process that is inherently and by design simple.

Kismet uses a bit expansion process that leverages very large integer multiplication and division, division being the biggest bottleneck, as it is an iterative computation that simply cannot be accelerated beyond the bit width of the processor. As it is, most modern 64 bit CPUs use almost 25% of the chip for this one operation, and GPUs only have 32 bit long division, thus making them half as efficient per cycle, countering their advantage. This also means that widely deployed, supply constrained off the shelf CPUs are the best hardhware for this competition, which means all who compete to win slots in the leadership of Kismet are in competition with virtually the whole internet for controlling processing power.

In this way, the PoW of Kismet is inherently more "green" because miners are competing with a massive number of alternative uses for the same hardware, and no other hardware is as efficient per watt. It will also have the benefit of driving up the price of expensive to manufacture 64 bit CPUs, improving their profit margins, and giving more opportunity to clean up the waste products of this industry.

In general, 'cloud providers' will not allow people to run CPU mining nodes on their systems. The token chain does not need high timing precision to participate adequately in, but validator nodes do need to on a fast network to participate in the consensus. It has never proven to be any benefit for security to piggyback onto another blockchain's proof of work mining protocol, and the small risk of sudden rises in mining capacity should be mitigated by the high cost of running them, no different from the original intention of Bitcoin. Kismet will run DivHash, and never permit merge mining.
