package chain_of_responsibility

import (
	"testing"
)

func TestA(t *testing.T) {
	chain := NewChain()
	chain.AddFilter(NewJunkFilter())
	chain.AddFilter(NewHelpFilter())
	chain.AddFilter(NewHelloFilter())
	chain.Filter("junk")
	chain.Filter("help")
	chain.Filter("hello")
	chain.Filter("money")
}
