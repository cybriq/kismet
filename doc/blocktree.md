# About Block Trees

Blockchains are not actually blockchains, but rather, slightly branchy trees that the consensus resolves one path often called the "best chain". I like to think of them as something more like a vine that has short, mostly abortive sidebranches. 

But vine doesn't sound as branchy and the chains do often branch whenever there is a coincidence of block solutions below the time precision of the network, mostly when a solution concurrently occurs in two separate locations under 12 seconds apart.

As such, to track this structure a blockchain node needs to have an index (see [index.go](../pkg/chain/index.go)) which stores every submitted block, which can contain several competing candidates for successors of a given predecessor block.

This process is subjective, as with all consensus processes, based on known data that (should be) is shared by all nodes on the network, namely the set of all known blocks. Because each node has to construct its own block tree, we can use a subjective index which does not need to be agreed by all nodes, namely, it can thus create an index with monotonic numbers of 64 bits and then create a tree with the following elements:

First, we need this index:

-   index -> block hash 

Next, each node of the blocktree has these elements:

-   block index
-   index of parent block
-   very large integer value that gives the current cumulative score of all prior blocks dating back to zero block

Note that Kismet does not have a genesis block. The first block(s) have zero as their parent block, which is the least possible hash value there is. Thus, there can be more than one branch from the very first block of Kismet. The "best" chain and its "head" block will be determined by the heaviest block that extends from the zero.

