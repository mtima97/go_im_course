package playground

import "testing"

func TestEmbedding(t *testing.T) {
	m := B{
		A:  A{ID: 1},
		ID: 2,
	}

	id := m.GetID()

	if id != 1 {
		t.Fatal("failed")
	}
}
