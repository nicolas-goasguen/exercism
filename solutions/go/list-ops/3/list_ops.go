package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := s.Length() - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	var filtered IntList
	for _, v := range s {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func (s IntList) Length() int {
	n := 0
	for range s {
		n++
	}
	return n
}

func (s IntList) Map(fn func(int) int) IntList {
	mapped := make(IntList, s.Length())
	for i, v := range s {
		mapped[i] = fn(v)
	}
	return mapped
}

func (s IntList) Reverse() IntList {
	reversed := make(IntList, s.Length())
	length := s.Length()
	for i := 0; i < length; i++ {
		reversed[(length - (1 + i))] = s[i]
	}
	return reversed
}

func (s IntList) Append(lst IntList) IntList {
	length := s.Length()
	added := make(IntList, length+lst.Length())
	for i, v := range s {
		added[i] = v
	}
	for i, v := range lst {
		added[length+i] = v
	}
	return added
}

func (s IntList) Concat(lists []IntList) IntList {
	intCount := s.Length()

	for _, l := range lists {
		intCount += l.Length()
	}

	concatenated := make(IntList, intCount)
	for i, v := range s {
		concatenated[i] = v
	}

	offset := s.Length()
	for _, l := range lists {
		for i, v := range l {
			concatenated[offset+i] = v
		}
		offset += l.Length()
	}

	return concatenated
}
