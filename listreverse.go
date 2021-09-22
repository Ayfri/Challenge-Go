package piscine

func ListReverse(l *List) {
	result := &List{}

	for iter := l.Head; iter != nil; iter = iter.Next {
		ListPushFront(result, iter.Data)
	}

	*l = *result
}
