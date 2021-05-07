package Sigleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	for i:=0; i<100; i++ {
		go func() {
			defer wg.Done()
			IncrementAge()
		}()

		go func() {
			defer wg.Done()
			IncrementAge2()
		}()
	}
	wg.Wait()
	p := GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}

func TestGetInstance1(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(12)
	for i:=0; i < 5; i++ {
		go func() {
			defer wg.Done()
			GetInstance1()
		}()
	}
	wg.Wait()
	GetInstance1()
}