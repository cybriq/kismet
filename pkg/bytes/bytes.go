package bytes

func ToUint64(bytes []byte) (out uint64) {
	out = uint64(bytes[0]) +
		uint64(bytes[1])<<8 +
		uint64(bytes[2])<<16 +
		uint64(bytes[3])<<24 +
		uint64(bytes[4])<<32 +
		uint64(bytes[5])<<40 +
		uint64(bytes[6])<<48 +
		uint64(bytes[7])<<56
	return
}

func FromUint64(i uint64) (bytes []byte) {
	bytes[0] = byte(i)
	bytes[1] = byte(i >> 8)
	bytes[2] = byte(i >> 16)
	bytes[3] = byte(i >> 24)
	bytes[4] = byte(i >> 32)
	bytes[5] = byte(i >> 40)
	bytes[6] = byte(i >> 48)
	bytes[7] = byte(i >> 56)
	return
}

func ToInt64(bytes []byte) (out int64) {
	out = int64(bytes[0]) +
		int64(bytes[1])<<8 +
		int64(bytes[2])<<16 +
		int64(bytes[3])<<24 +
		int64(bytes[4])<<32 +
		int64(bytes[5])<<40 +
		int64(bytes[6])<<48 +
		int64(bytes[7])<<56
	return
}

func FromInt64(i int64) (bytes []byte) {
	bytes[0] = byte(i)
	bytes[1] = byte(i >> 8)
	bytes[2] = byte(i >> 16)
	bytes[3] = byte(i >> 24)
	bytes[4] = byte(i >> 32)
	bytes[5] = byte(i >> 40)
	bytes[6] = byte(i >> 48)
	bytes[7] = byte(i >> 56)
	return
}
