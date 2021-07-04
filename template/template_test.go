package template

import (
	"testing"
)

func TestTemplates(t *testing.T) {
	NewTea().Make()
	NewCoffee().Make()
}