package functioned

import "math"

const INT_MAX = int(^uint(0) >> 1)

func getRes() (r1 int, r2 int) {

	r1, r2 = 1, 2

	return r1, r2
}

func getResNameless() (int, int) {
	return 1, 2
}

func Min(intSlice ...int) (minRes int) {

	minRes = math.MaxInt32

	for value := range intSlice {
		if value < minRes {
			value = minRes
		}
	}
	return minRes
}
