package template

import (
	"fmt"
	"testing"
)

func TestTemplates(t *testing.T) {
	NewTea().Make()
	fmt.Println("****************************")
	NewCoffee().Make()
}
