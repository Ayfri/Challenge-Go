package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	*l.Tail, *l.Head = *l.Head, NodeL{
		data,
		l.Head,
	}
}
