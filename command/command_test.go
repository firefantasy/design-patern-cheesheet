package command

import "testing"

func TestCommand(t *testing.T) {
	player := NewPlayer()
	step := []Command{NewPlayCommnd(player), NewStopCommnd(player)}
	step1 := []Command{NewNextCommnd(player), NewPreviousCommnd(player)}

	for _, s := range step {
		s.Execute()
	}
	
	for _, s := range step1 {
		s.Execute()
	}
}

