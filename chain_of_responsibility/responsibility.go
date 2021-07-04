package chain_of_responsibility

import (
	"fmt"
	"strings"
)

// 实现职责的接口
type MailFilter interface{
	Filter(string) bool
	SetSuccessor(MailFilter)
}

// 链表组织职责链
// 各个职责实体需要包含后继处理


func NewJunkFilter() *junkFilter {
	return &junkFilter{}
}

type junkFilter struct {
	successor MailFilter
}

func (j *junkFilter) Filter(content string) bool {
	if strings.Contains(content, "junk") {
		fmt.Println("this mail filter out by junkFilter")
		return true
	} 
	if j.successor != nil {
		return j.successor.Filter(content)
	}
	return false
}

func (j *junkFilter) SetSuccessor(mailFilter MailFilter) {
	j.successor = mailFilter
}

func NewHelpFilter() *helpFilter {
	return &helpFilter{}
}

type helpFilter struct {
	successor MailFilter
}

func (h *helpFilter) Filter(content string) bool {
	if strings.Contains(content, "help") {
		fmt.Println("this mail filter out by helpFilter")
		return true
	}
	if h.successor != nil {
		return h.successor.Filter(content) 
	}
	return false
}

func (h *helpFilter) SetSuccessor(filter MailFilter) {
	h.successor = filter
}

func NewHelloFilter() *helloFilter{
	return &helloFilter{}
}

type helloFilter struct {
	successor MailFilter
}

func (g *helloFilter) Filter(content string) bool {
	if strings.Contains(content, "hello") {
		fmt.Println("this mail filter out by helloFilter")
		return true
	}
	if g.successor != nil {
		return g.successor.Filter(content) 
	}
	return false
}


func (g *helloFilter) SetSuccessor(filter MailFilter) {
	g.successor = filter
}

func NewChain() *chain {
	return &chain{}
}

type chain struct {
	head, tail MailFilter
}

func (c *chain) AddFilter(f MailFilter) {
	f.SetSuccessor(nil)
	if c.head == nil {
		c.head, c.tail = f, f
		return 
	}
	c.tail.SetSuccessor(f)
	c.tail = f
}

func (c *chain) Filter(s string) {
	if c.head != nil {
		ok := c.head.Filter(s)
		if !ok {
			fmt.Printf("have no filter to do with content:%s", s)
		}
	}
}
