# tokens

## How tokens work in kismet

Tokens are issued at a fixed frequency for which the difficulty adjustment algorithm adjusts the maximum hash value target in order to maintain the frequency given the indefinite number of miners attempting to mint them.

The total amount of computation dedicated to mining kismet tokens will be directly proportional to the number of competitors vying for governance roles.

### Token chain block format

1. previous block hash
2. token type identifier
3. public key for token

### Token Chain Consensus

Because the data volume of the chain is quite small, instead of a consensus, the blocks are permitted to form a branched tangle format instead. There is no "canonical chain" but rather all new blocks on the chain must be linked to a previous block. This allows forks when the same token is issued to two miners, and not cause a conflict, just counts as two tokens that will affect the issuance rate control difficulty adjustment accordingly for the token type.

The only rule of the consensus is that previous blocks may not be more than an hour old compared to the new token time and nodes can refuse to attach blocks if their clock says the old stamp is too old, and that new tokens cannot be issued before half of their prescribed duration according to stamps.

### Token types

#### Validator

Validator tokens grant the a spot in the validator queue for a target validator set size of 60 with 60 usages per token, with expiry at 86400 seconds after issuance. 

This permits an effective maximum of 120 running validators at any given time, but the real effective maximum number of validators is around 90.

Issuance rate aims for 2 tokens per minute to effectively provide 2 minutes of coverage for the network

The validator tokens must be used to produce an announcement of service to become active and the existing validators will record the announcement in the next pBFT chain block, after which the new validators are appended to the validator queue.

The validator queue is maintained at a maximum of 90 members of which 61 votes are required to ratify a pBFT block.

#### Proposal

Proposal tokens allow the submission of a chain proposal, which lives on the token chain, and is a link to an IPFS filesystem containing the text of a proposal and ties together the proposal's necessary materials.

Proposal tokens are issued at a rate of 1 every 24 hours, and expire in 28 days

#### Congress

Congress tokens grant a position in a limited size congress to make votes on proposals

Congress tokens are issued when a Proposal token is used to publish a proposal

#### Veto

Veto token is a special type of token that cancels a proposal. One is issued at genesis by the Benevolent Dictator For Life, and can be used to issue delegate tokens by a special transaction on chain that a delegate provides a public key to authenticate.
