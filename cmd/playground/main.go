package main

import (
	"fmt"
	"gocourse/internal/service/playground"
)

func main() {
	m := playground.B{
		ID: 1,
		A:  playground.A{ID: 2},
	}

	fmt.Println(m.GetID())
}
