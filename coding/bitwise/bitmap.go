package bitwise

type BitMap struct {
	bitMap []byte
	size   int
}

func NewBitMap(maxNum int64) *BitMap {
	bitMap := make([]byte, (maxNum>>3)+1)
	return &BitMap{bitMap: bitMap, size: 0}
}

func (receiver *BitMap) IsExist(num int64) bool {
	arrayIndex := num >> 3
	bitIndex := num % 8
	return receiver.isExist(arrayIndex, bitIndex)
}

func (receiver *BitMap) isExist(arrayIndex int64, bitIndex int64) bool {
	return receiver.bitMap[arrayIndex]&(1<<bitIndex) != 0
}

func (receiver *BitMap) Set(num int64) {
	arrayIndex := num >> 3
	bitIndex := num % 8
	if !receiver.isExist(arrayIndex, bitIndex) {
		receiver.size++
		receiver.bitMap[arrayIndex] |= 1 << bitIndex
	}
}

func (receiver *BitMap) Remove(num int64) {
	arrayIndex := num >> 3
	bitIndex := num % 8
	if receiver.isExist(arrayIndex, bitIndex) {
		receiver.size--
		receiver.bitMap[arrayIndex] &= ^(1 << bitIndex)
	}
}

func (receiver *BitMap) Reverse() []int {
	reverse := make([]int, receiver.size)
	reverseIndex := 0
	for arrayIndex := 0; arrayIndex < len(receiver.bitMap); arrayIndex++ {
		value := receiver.bitMap[arrayIndex]

		for bitIndex := 0; bitIndex < 8; bitIndex++ {
			if value > 0 {
				if value&1 == 1 {
					reverse[reverseIndex] = arrayIndex*8 + bitIndex
					reverseIndex++
				}
				value >>= 1
			}
		}
	}

	return reverse
}
