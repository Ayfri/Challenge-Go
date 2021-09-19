package piscine

func ForEach(f func(int), a []int) {
	for _, integer := range a {
		f(integer)
	}
}
