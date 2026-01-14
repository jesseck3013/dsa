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
	res := merge(s[:pivot], s[pivot:])
	copy(s, res)
}

func merge(s1, s2 []int) []int {
	res := make([]int, len(s1)+len(s2))
	l1 := 0
	l2 := 0
	for i := range res {
		if l1 < len(s1) && l2 < len(s2) {
			if s1[l1] < s2[l2] {
				res[i] = s1[l1]
				l1++
			} else {
				res[i] = s2[l2]
				l2++
			}
		} else if l1 < len(s1) {
			res[i] = s1[l1]
			l1++

		} else if l2 < len(s2) {
			res[i] = s2[l2]
			l2++
		} else {
		}
	}

	return res
}
