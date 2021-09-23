package piscine

func ListMerge(l1 *List, l2 *List) {
	for iter := l1.Head; iter != nil; iter = iter.Next {
		ListPushFront(l2, iter.Data)
	}
}
