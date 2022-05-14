# tokens

## How tokens work in proto

Tokens are issued at a fixed frequency for which the difficulty adjustment algorithm adjusts the maximum hash value target in order to maintain the frequency given the indefinite number of miners attempting to mint them.

### Token chain block format

1. previous block hash
2. token type identifier
3. public key for token

### Token types

#### Validator

Validator tokens grant the a spot in the validator queue for a target validator set size of 60 with 60 usages per token, with expiry at 86400 seconds after issuance

Issuance rate aims for 2 tokens per minute to effectively provide 2 minutes of coverage for the network

#### Proposal

Proposal tokens allow the submission of a chain proposal, which lives on the token chain, and is a link to an IPFS filesystem containing the text of a proposal and ties together the proposal's necessary materials.

Proposal tokens are issued at a rate of 1 every 24 hours, and expire in 28 days

#### Congress

Congress tokens grant a position in a limited size congress to make votes on proposals

Congress tokens are issued when a Proposal token is used to publish a proposal

#### Veto

Veto token is a special type of token that cancels a proposal. One is issued at genesis by the Benevolent Dictator For Life, and can be used to issue delegate tokens by a special transaction on chain that a delegate provides a public key to authenticate.
