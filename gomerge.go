package gomerge

import (
	"fmt"
)

type Gomerge struct {
}

func (g *Gomerge) Println(a ...any) (n int, err error) {
	str := a
	return fmt.Println(str)
}
