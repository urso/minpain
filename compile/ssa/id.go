package ssa

type ID uint

type idAlloc struct {
	last ID
}

func (a *idAlloc) next() ID {
	x := a.last + 1
	if x == 0 {
		panic(errTooManyIDs)
	}

	a.last = x
	return x
}
