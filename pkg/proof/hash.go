package proof

import (
	"lukechampine.com/blake3"
	"math/big"
)

func reverse(b []byte) []byte {

	bytesLen := len(b)
	halfBytesLen := bytesLen / 2
	for i := 0; i < halfBytesLen; i++ {

		// Reversing items in a slice without explicitly defining the intermediary
		// (should be implemented via register)
		b[i], b[bytesLen-i] = b[bytesLen-i], b[i]
	}

	return b
}

// Blake3 takes bytes and returns a Blake3 256 bit hash
func Blake3(bytes []byte) []byte {

	b := blake3.Sum256(bytes)
	return b[:]
}

// DivHash is a hash function that combines the use of very large integer
// multiplication and division in addition to Blake3 hashes to create extremely
// large integers that cannot be produced without performing these very time
// expensive iterative long division steps.
//
// The function has a parameter to repeat the operation. It can completely blow
// out the stack and heap if it is repeated enough times, so this number is
// usually only somewhere between 2 and 5 steps at most.
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

	// Convert the reverse of original block, and the two above values to big
	// integers
	reversedBlockInt := big.NewInt(0).SetBytes(reverse(blockBytes))
	firstHalfInt := big.NewInt(0).SetBytes(firstHalf)
	secondHalfInt := big.NewInt(0).SetBytes(secondHalf)

	// square each half, then multiply the two products together, and divide by the
	// reverse of the original block
	squareFirstHalf := firstHalfInt.Mul(firstHalfInt, firstHalfInt)
	squareSecondHalf := secondHalfInt.Mul(secondHalfInt, secondHalfInt)
	productOfSquares := firstHalfInt.Mul(squareFirstHalf, squareSecondHalf)
	productDividedByBlockInt := productOfSquares.Div(productOfSquares, reversedBlockInt)
	ddd := productDividedByBlockInt.Bytes()

	// Scramble the product by hashing progressively shorter segments to produce a
	// scrambled version
	dddLen, dddMod := len(ddd)/32, len(ddd)%32
	if dddMod > 0 {
		dddLen++
	}
	output := make([]byte, dddLen*32)
	for i := 0; i < dddLen; i++ {

		// we are hashing the next 32 bytes each time
		segment := Blake3(ddd[32*i : 32*(i+1)])
		copy(output[32*i:32*(i+1)], segment)
	}

	// trim the result back to the original length
	output = output[:len(ddd)]

	// By repeating this process several times we end up with an extremely long
	// value that doesn't have a shortcut to creating it, and requiring very common
	// but expensive long division units to produce.
	if howmany > 0 {

		return DivHash(output, howmany-1)
	}

	// After all repetitions are done, the very large bytes produced at the end are
	// hashed and reversed.
	return reverse(Blake3(output))
}
