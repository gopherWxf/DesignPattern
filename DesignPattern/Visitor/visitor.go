package Visitor

import "fmt"

type IVisitor interface {
	Visit()
}
type WeiBoVisitor struct {
}

func (w WeiBoVisitor) Visit() {
	fmt.Println("visit weibo")
}

type IQIYIVisitor struct {
}

func (I IQIYIVisitor) Visit() {
	fmt.Println("visit iqiyi")
}

type IElement interface {
	Accept(visitor IVisitor)
}
type Element struct {
}

func (e Element) Accept(v IVisitor) {
	v.Visit()
}
