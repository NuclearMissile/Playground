package bitArray

const WORD_SIZE = 32 << (^uint(0) >> 63)

type BitArray struct {
	words []uint
}

func (s *BitArray) Has(x int) bool {
	index, offset := x/WORD_SIZE, uint(x%WORD_SIZE)
	return index < len(s.words) && s.words[index]&(1<<offset) != 0
}

func (s *BitArray) Add(x int) {
	index, offset := x/WORD_SIZE, uint(x%WORD_SIZE)
	for index >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[index] |= 1 << offset
}

func (s *BitArray) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

func (s *BitArray) UnionWith(ba *BitArray) {
	for i, baIndex := range ba.words {
		if i < len(s.words) {
			s.words[i] &= baIndex
		} else {
			s.words = append(s.words, baIndex)
		}
	}
}
