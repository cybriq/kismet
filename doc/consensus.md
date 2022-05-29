# consensus

Just as in any blockchain that is grown through a distributed lottery called "proof of work", there must be rules to arbitrate whether a new block is considered to be canonical.

1. Kismet governance chain does not issue fungible tokens - there is not and never will be a way to novate (issue new key pair based on old, necessary for exchange) governance tokens, and they expire when the current supply rate passes the defined threshold.
2. Standard heaviest chain selection for consensus as Bitcoin.
3. Blocks cannot be timestamped less than half of their target time period in the past, nodes reject any new transaction first seen with older than the time period from node's subjective clock.
4. The governance chain has the potential to regulate issuance of 256 different token types, which potentially can be multiple other different chains. There would be no need to replicate the Proposal and Congress tokens, if needed they could be increased in frequency and active number to handle a larger number of governed chains.
5. The validator chain, and the proposal chain are separate chains. Likewise, any added chains created by different token types can be on their own chains. Proof of Work mining has structural limit of around 12 second average global propagation rate, so each chain can probably handle this rate of token issuance independently. If it ever became necessary to extend the number of tokens beyond 256, it is unlikely that 65536 would be exceeded.