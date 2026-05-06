package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	if s.Length() == 0 {
		return initial
	}
	initial = fn(initial, s[0])
	if s.Length() == 1 {
		return initial
	}
	return s[1:].Foldl(fn, initial)
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	if s.Length() == 0 {
		return initial
	}
	initial = fn(s[s.Length()-1], initial)
	if s.Length() == 1 {
		return initial
	}
	return s[:s.Length()-1].Foldr(fn, initial)
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
	for i := 0; i < s.Length(); i++ {
		reversed[(s.Length() - (1 + i))] = s[i]
	}
	return reversed
}

func (s IntList) Append(lst IntList) IntList {
	added := make(IntList, s.Length()+lst.Length())
	for i, v := range s {
		added[i] = v
	}
	for i, v := range lst {
		added[s.Length()+i] = v
	}
	return added
}

func (s IntList) Concat(lists []IntList) IntList {
	intCount := 0
	for _, l := range lists {
		intCount += l.Length()
	}

	concatenated := make(IntList, intCount)
	offset := 0
	for _, l := range lists {
		for i, v := range l {
			concatenated[offset+i] = v
		}
		offset += l.Length()
	}

	return concatenated
}
