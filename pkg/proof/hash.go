package proof

import (
	"math/big"

	"lukechampine.com/blake3"
)

func reverse(b []byte) []byte {

	bytesLen := len(b)
	halfBytesLen := bytesLen / 2
	for i := 0; i < halfBytesLen; i++ {

		// Reversing items in a slice without explicitly defining the intermediary
		// (should be implemented via register)
		b[i], b[bytesLen-i-1] = b[bytesLen-i-1], b[i]
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
// This hash function has an operation time proportional to the size of the
// input. As such, applications using this function must use a repetition
// parameter fit for the duration delay sought in the application.
//
// This hash function will have a relatively flat difference in performance
// proportional to the integer long division unit in the processor. Because this
// unit is the most complex, and the operation is the most non-optimisable
// integer mathematics function this forms the basis of the fairest possible
// proof of work function that places miners on the flattest playing field
// possible and pits their use of this hardware against almost all computer
// systems applications in market competition, keeping the growth of hashrate
// constrained. Long division performance is almost linearly proportional to
// transistor count, which is almost linearly proportional to relative cost.
func DivHash(blockBytes []byte, repetitions int) []byte {

	if len(blockBytes) < 2 {
		panic("DivHash may not be computed with less than two bytes of input")
	}

	blockLen := len(blockBytes)

	// Reverse first half and append to the end of the original bytes
	firstHalf := make([]byte, blockLen+blockLen/2)
	copy(firstHalf[:blockLen], blockBytes)
	copy(firstHalf[blockLen:], reverse(blockBytes[:blockLen/2]))

	// Reverse second half and append to the end of the original bytes
	secondHalf := make([]byte, blockLen+blockLen/2)
	copy(secondHalf[:blockLen], blockBytes)
	copy(secondHalf[blockLen:], reverse(blockBytes[blockLen/2:]))

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
	productDividedByBlockInt := productOfSquares.Div(productOfSquares,
		reversedBlockInt)
	ddd := productDividedByBlockInt.Bytes()

	// Scramble the product by hashing progressively shorter segments to produce a
	// scrambled version
	dl := len(ddd)
	dddLen, dddMod := dl/32, dl%32
	if dddMod > 0 {
		dddLen++
	}
	output := make([]byte, dddLen*32)
	for i := 0; i < dddLen; i++ {

		// clamp the end to the end
		end := 32 * (i + 1)
		if end > dl {
			end = dl
		}

		// we are hashing the next 32 bytes each time
		segment := Blake3(ddd[32*i : end])
		copy(output[32*i:32*(i+1)], segment)
	}

	// trim the result back to the original length
	output = output[:dl]

	// By repeating this process several times we end up with an extremely long
	// value that doesn't have a shortcut to creating it, and requiring very common
	// but expensive long division units to produce.
	if repetitions > 0 {

		return DivHash(output, repetitions-1)
	}

	// After all repetitions are done, the very large bytes produced at the end are
	// hashed and reversed.
	return reverse(Blake3(output))
}

func DivHash4(input []byte) []byte {
	return DivHash(input, 4)
}
