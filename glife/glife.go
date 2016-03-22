package glife

type Board [][]bool

type Pos [2]int

func (b *Board) Initialize(x, y int) {
	*b = make([][]bool, x)
	for i := 0; i < y; i++ {
		(*b)[i] = make([]bool, y)
	}
}

func (b *Board) LeftAdd() {
	t := make([][]bool, 1)
	t[0] = make([]bool, len((*b)[0]))
	t = append(t, (*b)...)
	*b = t
}

func (b *Board) AboveAdd() {
	for i := 0; i < len(*b); i++ {
		(*b)[i] = append(make([]bool, 1), ((*b)[i])...)
	}
}

func (b *Board) RightAdd() {
	*b = append(*b, make([]bool, len((*b)[0])))
}

func (b *Board) BelowAdd() {
	for i := 0; i < len(*b); i++ {
		(*b)[i] = append((*b)[i], false)
	}
}

type OAC Pos

func (b *Board) Eval(p Pos) {}
