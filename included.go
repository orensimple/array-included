package included

import (
	"math"
)

func IsInclude(array []int, subarray []int) bool {
	var countInc, firstInc int

	if len(subarray) == 0 {
		return true
	}

	if len(array) == 0 || array[len(array)-1] < subarray[len(subarray)-1] {
		return false
	}

	// binary search for most cases
	if math.Log2(float64(len(array)))+float64(len(subarray)) <= float64(len(array)) {
		firstInc = binarySearch(array, subarray[0])
	}

	for i := firstInc; i < len(array) && i != -1 && len(array)-i >= len(subarray)-countInc; i++ {
		if array[i] == subarray[countInc] {
			countInc++

			if len(subarray) == countInc {
				return true
			}

			continue
		}

		if countInc > 0 {
			return false
		}
	}

	return false
}

func binarySearch(a []int, search int) (result int) {
	mid := len(a) / 2

	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid] > search:
		result = binarySearch(a[:mid], search)
	case a[mid] < search:
		result = binarySearch(a[mid+1:], search)
		if result >= 0 {
			result += mid + 1
		}
	default:
		result = mid // found
	}

	return
}
