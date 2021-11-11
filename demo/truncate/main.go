package main

type items []int

var nilItems = make(items, 16)

func (s *items) truncate1(i int) {
	*s = (*s)[:i]
}

func (s *items) truncate2(i int) {
	var toClear items
	*s, toClear = (*s)[:i], (*s)[i:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilItems):]
	}
}

func main() {
	s := items{10, 20, 30}
	s.truncate2(1)
}
