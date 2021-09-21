package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}

	iter := l.Head
	for iter.Next != nil {
		iter = iter.Next
	}

	return iter.Data
}
