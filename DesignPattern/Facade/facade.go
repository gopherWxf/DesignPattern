package Facade

import "fmt"

type CarModel struct {
}

func NewCarModel() *CarModel {
	return &CarModel{}
}

func (c *CarModel) SetModel() {
	fmt.Println("CarModel - SetModel")
}
func (c *CarModel) ModelBroke() {
	fmt.Println("Model was Broke")
}

type CarEngine struct {
}

func NewCarEngine() *CarEngine {
	return &CarEngine{}
}
func (c *CarEngine) SetEngine() {
	fmt.Println("CarEngine - SetEngine")
}
func (c *CarEngine) EngineBroke() {
	fmt.Println("Engine was Broke")

}

type CarBody struct {
}

func NewCarBody() *CarBody {
	return &CarBody{}
}
func (c *CarBody) SetCarBody() {
	fmt.Println("CarBody - SetBody")
}
func (c *CarBody) BodyBroke() {
	fmt.Println("Body was Broke")
}

type CarFacade struct {
	model  CarModel
	engine CarEngine
	body   CarBody
}

func NewCarFacade() *CarFacade {
	return &CarFacade{
		model:  CarModel{},
		engine: CarEngine{},
		body:   CarBody{},
	}
}
func (c *CarFacade) CreateCompleteCar() {
	c.model.SetModel()
	c.body.SetCarBody()
	c.engine.SetEngine()
}
func (c *CarFacade) MethodA() {
	c.model.SetModel()
	c.body.BodyBroke()
	c.engine.SetEngine()
	c.model.ModelBroke()
}
func (c *CarFacade) MethodB() {
	c.model.ModelBroke()
	c.body.BodyBroke()
	c.engine.EngineBroke()
}
