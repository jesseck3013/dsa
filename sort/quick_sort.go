package sort

import (
	"math/rand"
)

func QuickSort(s []int) {
	if len(s) < 2 {
		return
	}

	pivotIndex := rand.Intn(len(s))
	pivot := s[pivotIndex]
	leftPartiion := make([]int, 0)
	rightPartiion := make([]int, 0)

	for i, v := range s {
		if i != pivotIndex {
			if v < pivot {
				leftPartiion = append(leftPartiion, v)
			} else {
				rightPartiion = append(rightPartiion, v)
			}
		}
	}

	QuickSort(leftPartiion)
	QuickSort(rightPartiion)

	for i, v := range leftPartiion {
		s[i] = v
	}

	s[len(leftPartiion)] = pivot

	for i, v := range rightPartiion {
		s[len(leftPartiion)+1+i] = v
	}
}
