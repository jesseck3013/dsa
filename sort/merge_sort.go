package sort

func MergeSort(s []int) {
	if len(s) <= 1 {
		return
	}

	if len(s) == 2 {
		if s[0] > s[1] {
			s[0], s[1] = s[1], s[0]
		}
		return
	}

	pivot := len(s) / 2

	MergeSort(s[:pivot])
	MergeSort(s[pivot:])
}
