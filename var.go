package main

const (
	VarInt = 1
)

type Var struct {
	Type int
	Int  int
}

func NewIntVar(val int) *Var {
	return &Var{Type: VarInt, Int: val}
}
