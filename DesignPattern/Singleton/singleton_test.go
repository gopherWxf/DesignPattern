package Singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	wg:=sync.WaitGroup{}
	wg.Add(200)
	for i:=0;i<100;i++{
		go func(){
			defer wg.Done()
			IncrementAge1()
		}()
		go func(){
			defer wg.Done()
			IncrementAge2()
		}()
	}
	wg.Wait()
	p:=GetInstance()
	age:=p.GetAge()
	fmt.Println(age)
}
func TestGetInstance02(t *testing.T) {
	p1:=GetInstance()
	p2:=GetInstance()
	if p1==p2 {
		fmt.Println("单例")
	}
}