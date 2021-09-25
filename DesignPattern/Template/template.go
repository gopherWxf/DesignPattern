package Template

import "fmt"

type WorkInterface interface {
	GetUp()
	Work()
	Sleep()
}
type Worker struct {
	WorkInterface
}

func NewWorker(w WorkInterface) *Worker {
	return &Worker{WorkInterface: w}
}

func (w *Worker) Daily() {
	w.GetUp()
	w.Work()
	w.Sleep()
}

type Coder struct {
}

func (c *Coder) GetUp() {
	fmt.Println("Coder GetUp ")
}

func (c *Coder) Work() {
	fmt.Println("Coder Work ")
}

func (c *Coder) Sleep() {
	fmt.Println("Coder Sleep ")

}
