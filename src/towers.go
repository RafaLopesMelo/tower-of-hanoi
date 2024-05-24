package main

import (
	"os"
)

type TowerID string

const (
	A TowerID = "A"
	B TowerID = "B"
	C TowerID = "C"
)

type Towers struct {
	A *Stack
	B *Stack
	C *Stack
	moves int
}

func (t *Towers) Move(from *Stack, to *Stack) {
	element, err := from.Pop()

	if err {
		os.Exit(1)
	}

	to.Push(element)
	t.moves = t.moves + 1
}

func (t *Towers) Moves() int {
	return t.moves
}

func (t *Towers) Tower(tower TowerID) *Stack {
	switch tower {
		case TowerID(A):
			return t.A;
		case TowerID(B):
			return t.B;
		case TowerID(C):
			return t.C;
	}

	os.Exit(1)
	return nil
}

func (t *Towers) Decrease()  {
	t.A.Decrease()
	t.B.Decrease()
	t.C.Decrease()
}

func (t *Towers) Increase()  {
	t.A.Increase()
	t.B.Increase()
	t.C.Increase()
}

func NewTowers(size int) *Towers {
	limit := size + 1

	a := NewStack(limit)
	for i := 0; i < size; i++ {
		a.Push(size - i)
	}

	towers := &Towers{
		A: a,
		B: NewStack(limit),
		C: NewStack(limit),
		moves: 0,
	}

	return towers
}