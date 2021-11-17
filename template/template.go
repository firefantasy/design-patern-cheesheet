package template

import "fmt"


type Make interface {
	Make()
}

type Impl interface {
	Brew() // 抽象出来的具体步骤
	Name() string
	IsAddSomething() bool // 钩子方法
	Add() 
}

func NewBrew(i Impl) *Beverage{
	return &Beverage{i}
}

type Beverage struct {
	Impl
}

func (i *Beverage) Boil() {
	fmt.Println("prepare to boil water")
}

func (i *Beverage) Pour() {
	fmt.Println("pour the drink to cup")
}

func (i *Beverage) IsAddSomething() bool {
	return true
}

func (i *Beverage) Add() {
	fmt.Println("add nothing")
}

// 模板方法
func (i *Beverage) Make() {
	fmt.Println("ok, now prepare to make the beverage")
	// 注意必须使用 i.Impl.xx 调用
	i.Boil() 
	i.Impl.Brew()
	// 预埋钩子，具体实现可以通过重新定义该方法有机会影响模板的执行
	if i.Impl.IsAddSomething() {
		i.Impl.Add()
	} else {
		i.Add()
	}
	i.Pour()
}


// golang 没有继承，具体实现上需要内嵌Beverage实体

func NewTea() Make {
	i := &tea{}
	i.Beverage = NewBrew(i)
	return i
}

type tea struct {
	*Beverage
}

func (i *tea) Name() string {
	return "tea"
}

func (tea *tea) Brew() {
	fmt.Println("make tea")
}
func (i *tea) Add() {
	fmt.Println("add lemon to Tea !")
}

func NewCoffee() Make {
	i := &coffee{}
	i.Beverage = NewBrew(i)
	return i
}

type coffee struct {
	*Beverage
}

func (c *coffee) Name() string {
	return "coffee"
}

func (c *coffee) Brew() {
	fmt.Println("make coffee")
}

func (i *coffee) IsAddSomething() bool {
	return true
}



