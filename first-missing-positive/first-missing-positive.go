package first_missing_positive

func firstMissingPositive(nums []int) int {
	bs := newBitSet(uint64(len(nums)))

	for _, num := range nums {
		if num > 0 && num <= len(nums) {
			bs.set(uint64(num))
		}
	}

	i := 1
	for ; i <= len(nums); i++ {
		if !bs.exist(uint64(i)) {
			return i
		}
	}
	return i
}

var wordSize uint64 = 64

type bitset struct {
	sets []uint64
}

func newBitSet(n uint64) *bitset {
	bs := new(bitset)
	bs.sets = make([]uint64, n/wordSize+1)
	return bs
}

func (bs *bitset) set(i uint64) {
	slot := (i - 1) / wordSize
	bs.sets[slot] |= 1 << (i - 1 - slot*wordSize)
}

func (bs *bitset) exist(i uint64) bool {
	slot := (i - 1) / wordSize
	exists := (bs.sets[slot])&(1<<(i-1-slot*wordSize)) > 0

	return exists
}
