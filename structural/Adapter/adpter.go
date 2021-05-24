package Adapter

import "fmt"

// 适配器模式

type Target interface {
	Execute()
}

type Adaptee struct {
}

func (a *Adaptee) SpecificExecute() {
	fmt.Println("国士无双……")
}

type Adapter struct {
	*Adaptee
}

func (a *Adapter) Execute() {
	a.SpecificExecute()
}
