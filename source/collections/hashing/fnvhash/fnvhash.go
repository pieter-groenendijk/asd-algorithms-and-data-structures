package fnvhash

const offset uint32 = 2166136261
const prime uint32 = 16777619

func Hash(value []byte) uint32 {
	hash := offset
	for _, valueByte := range value {
		hash ^= uint32(valueByte)
		hash *= prime
	}

	return hash
}
