package Singleton

func IncrementAge1(){
	p:=GetInstance()
	p.IncrementAge()
}