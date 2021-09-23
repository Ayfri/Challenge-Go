package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	result := &List{}
	iter := l.Head
	for iter != nil {
		if iter.Data == data_ref {
			iter = iter.Next
			continue
		}
		ListPushBack(result, iter.Data)
	}
	*l = *result
}
