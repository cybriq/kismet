package proof

import (
	"lukechampine.com/blake3"
	"math/big"
)

func reverse(b []byte) []byte {

	bytesLen := len(b)
	halfBytesLen := bytesLen / 2
	for i := 0; i < halfBytesLen; i++ {

		// Reversing items in a slice without explicitly defining the intermediary (should be implemented via register)
		b[i], b[bytesLen-i] = b[bytesLen-i], b[i]
	}

	return b
}

// Blake3 takes bytes and returns a Blake3 256 bit hash
func Blake3(bytes []byte) []byte {

	b := blake3.Sum256(bytes)
	return b[:]
}

// DivHash is a hash function that cuts up the provided bytes, splicing reversed halves together,
// squaring the halves and then multiplying the products, and repeating with the splice of the forward and reversed
// product a number of times before finally hashing the bytes of the final very big integer.
//
// This serves to create a processing bottleneck of long division that causes intentional,
// but deterministic rounding errors which can only be created by performing the entire operation,
// and uses very large integers to ensure that it blows over the top even the largest Level 1 CPU caches when
// calculating. This will chiefly give advantage to processors with large caches and the longest bit long division
// units, ie, AMD Zen architecture CPUs.
func DivHash(blockBytes []byte, howmany int) []byte {

	blockLen := len(blockBytes)

	// Reverse first half and append to the end of the original bytes
	firstHalf := make([]byte, blockLen+blockLen/2)
	copy(firstHalf[:blockLen], blockBytes)
	copy(firstHalf[blockLen:], reverse(blockBytes[:blockLen/2]))

	// Reverse second half and append to the end of the original bytes
	secondHalf := make([]byte, blockLen+blockLen/2)
	copy(firstHalf[:blockLen], blockBytes)
	copy(firstHalf[blockLen:], reverse(blockBytes[blockLen/2:]))

	// Convert the reverse of original block, and the two above values to big integers
	reversedBlockInt := big.NewInt(0).SetBytes(reverse(blockBytes))
	firstHalfInt := big.NewInt(0).SetBytes(firstHalf)
	secondHalfInt := big.NewInt(0).SetBytes(secondHalf)

	// square each half, then multiply the two products together, and divide by the reverse of the original block
	squareFirstHalf := firstHalfInt.Mul(firstHalfInt, firstHalfInt)
	squareSecondHalf := secondHalfInt.Mul(secondHalfInt, secondHalfInt)
	productOfSquares := firstHalfInt.Mul(squareFirstHalf, squareSecondHalf)
	productDividedByBlockInt := productOfSquares.Div(productOfSquares, reversedBlockInt)
	ddd := productDividedByBlockInt.Bytes()

	// By repeating this process several times we end up with an extremely long value that doesn't have a shortcut to
	// creating it, and requiring very common but expensive long division units to produce.
	if howmany > 0 {

		return DivHash(append(ddd, reverse(ddd)...), howmany-1)
	}

	// After all repetitions are done, the very large bytes produced at the end are hashed and reversed.
	return reverse(Blake3(ddd))
}
