package piscine

func ListSize(l *List) int {
	result := 0
	iter := l.Head

	for iter != nil {
		iter = iter.Next
		result++
	}

	return result
}
