package playground

type A struct {
	ID int
}

func (a *A) GetID() int {
	return a.ID
}

type B struct {
	ID int
	A  // embedding
}
