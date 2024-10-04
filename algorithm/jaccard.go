package algorithm

import "math"

// jaccard 算法旨在寻找两个集合相似度，两个集合需要是已排序
func JaccardForSortedCollection[T Comparable](collection1, collection2 []T) float64 {
	if len(collection1) == 0 || len(collection2) == 0 {
		return 0
	}
	i, j := 0, 0
	intersection := 0
	for i < len(collection1) && j < len(collection2) {
		if collection1[i] == collection2[j] {
			intersection++
			i++
			j++
		} else if collection1[i] < collection2[j] {
			i++
		} else {
			j++
		}
	}

	return math.Round(float64(intersection)/float64(len(collection1)+len(collection2)-intersection)*100) / 100
}
