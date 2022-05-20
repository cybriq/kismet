# consensus

Just as in any blockchain that is grown through a distributed lottery called "proof of work", there must be rules to arbitrate whether a new block is considered to be valid or not.

The proof of work lottery is for leadership, that is, the right to publish a block. The consensus is which chain is considered to be valid.

One of the difficult problems of the proof of work lottery comes from the relatively frequent occurrence of block solutions that are closer together than the network can disambiguate one as being before the other.

In the bitcoin consensus, this decision is made based on the numerical value of the hash of the block - the lowest hash value, being more rare than higher, is selected when there is a conflict of two blocks to succeed a previous block.

Because Kismet's chain does not certify transactions or directly grant fungible tokens, it does not have the same strict need to form a singular chain. The chain can have short branches, and multiple parallel threads. 

The resolution of this is necessary to establish a difficulty adjustment to maintain a prescribed frequency of issuance of the various governance tokens that Kismet uses, validator, proposer and congress.

## Important differences from Bitcoin Consensus

    1. No fungible tokens are issued

The blocks created for Kismet chain are not for registering transactions, but rather, recording the results of an ongoing lottery for time and use limited tokens that grant the right to perform governance tasks on a separate but connected distributed ledger (or ledgers).

Thus, it doesn't matter if two solutions come at the same time, they are not considered to be in conflict. They just have the effect of doubling the difficulty reduction of building a token on either side.

When a fork is found, miners are required to link to both known fork points. All nodes will reject any block that does not close a known fork.

    2. Merging forks beats extending them, merged blocks are double weight

It can happen that a node does not see the branch of a fork before another miner has issued a single link forward from one side, but the heaviest chain consensus will put non-merging second and later blocks on a fork on side branches and due to metastability, the network will prefer the heavier merged forks over the continuing forks

Nodes that decide not to do this, for whatever reason, will not have their blocks propagated by the mempools of other nodes that see the coincident, forking blocks.

Because the tokens do not have immediate utility, but rather simply register a claim that can be used to gain the right to perform an action on the main ledger, it does not matter if they fork, only that forks are merged back.

In Kismet, the head of the chain is expected to reorganise regularly.

It can even happen that two fork merges occur at the same time. The same rule still applies, and so sometimes a chain will have two parallel paths, and then merge back together, and split again, and merge again. Most of the time the split will last only one 'height'.

Merge blocks have double difficulty, but are preferred over blocks that extend them, thus, nodes will choose to mine forks over extensions, when there is a fork.

    3. Each token type is a parallel chain

There is three token types in the Kismet chain, Validator, Proposer and Congress. The last two are in a single chain, and expire when a new block is found that makes the current active proposals exceed a given count, the Congress blocks must refer to the proposal hash found on the pBFT ledger.

Blocks of each of the three types can only refer to previous blocks of the same type. There is no relationship between the parallel chains, they simply record the issuance of new privilege tokens and control the rate with a difficulty adjustment consensus.

    4. Monotonic timestamps, half time minimum difference, time tolerance rejection

Timestamps on blocks must always be after the blocks they refer to.

Timestamps cannot be less than half of the block time target after.

Blocks with timestamps more than the block time target different from the current time subjective to a node are not accepted. This forces nodes to keep their time accurate, without creating a vulnerable time averaging consensus, it is in the interests of a miner to consume their mining capabilities and the energy required to run it, to not produce blocks with timestamps that will be rejected.