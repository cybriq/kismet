# tokens

## How tokens work in kismet

Tokens are issued at a fixed frequency for which the difficulty adjustment algorithm adjusts the maximum hash value target in order to maintain the frequency given the indefinite number of miners attempting to mint them.

The total amount of computation dedicated to mining kismet tokens will be directly proportional to the number of competitors vying for governance roles.

### Token chain block format

1. [32 bytes] previous block hash
2. [1 byte] token type identifier (most significant bit indicates a merge)
3. [8 bytes] 64 bit nanosecond precise Unix timestamp (nanoseconds since January 1, 1970 UTC, which can record time until 2262 AD)
4. [32 bytes] difficulty target (computed from parent block(s) targets and timestamps)
5. [32 bytes] ed25519 public key for token. Miner has a corresponding 64 byte private key that this key authenticates.
6. [32 bytes] other previous block hash (expected with merge bit set, otherwise not present)

The blake3 hash of this entire block of data is the block hash.

### Token types

#### Validator

Validator tokens grant the a spot in the validator queue for a target validator set size of 60 with 60 usages per token, with expiry at 86400 seconds after issuance. 

This permits an effective maximum of 120 running validators at any given time, but the real effective maximum number of validators is around 90.

Issuance rate aims for 2 tokens per minute to effectively provide 2 minutes of coverage for the network

The validator tokens must be used to produce an announcement of service to become active and the existing validators will record the announcement in the next pBFT chain block, after which the new validators are appended to the validator queue.

The validator queue is maintained at a maximum of 90 members of which 61 votes are required to ratify a pBFT block.

#### Proposal

Proposal tokens allow the submission of a chain proposal, which lives on the token chain, and is a link to an IPFS filesystem containing the text of a proposal and ties together the proposal's necessary materials.

Proposal tokens expire over time, at any given time the chain permits a number of proposals to exist, and when a new proposal block is found, the oldest is expired. Proposals are mined at a constant rate with a difficulty adjustment schedule that maintains an average time between token issuance, putting a cap on the rate of creating proposals. A proposal can be made to change this rate, if it is found that there is too many proposals being used or too few.

The veto power can be used to cancel proposals under the sole condition of changing the per block reward for validators, which is capped at 5% growth per year.

Proposal tokens can be used to create proposals on the pBFT ledger chain, so long as they are used before they are expired.

Proposals on the pBFT ledger are simply an IPFS hash to a filesystem tree maintained by the proposer.

#### Congress

Congress tokens grant a position in a limited size congress to make votes on proposals

Congress tokens are issued when a Proposal token is used to publish a proposal

Congress tokens are issued at schedule rate of issuance and once the quorum maximum size (give or take one fork at the end) is reached the tokens cannot be further mined, and then the miner can use them to issue a vote on the proposal the token is linked to.

This is an example of a quiescent subchain, it only can be mined if a proposal token is used to mint a new proposal. 

#### Veto

Veto token is a special type of token that cancels a proposal. One is issued at genesis by the Benevolent Dictator For Life, and can be used to issue delegate tokens by a special transaction on chain that a delegate provides a public key to authenticate, in order to revocably grant veto power to others to watch over proposals and reject proposals to change the token emission rate.
