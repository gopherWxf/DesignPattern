package Factory

import "fmt"

type Restaurant interface {
	GetFood()
}
type SpecificRestaurantA struct {
}

func (s *SpecificRestaurantA) GetFood() {
	fmt.Println("SpecificRestaurantA的饭菜已经生产完毕...")
}

type SpecificRestaurantB struct {
}

func (s *SpecificRestaurantB) GetFood() {
	fmt.Println("SpecificRestaurantB的饭菜已经生产完毕...")
}

func NewRestaurant(res string) Restaurant {
	switch res {
	case "A":
		return &SpecificRestaurantA{}
	case "B":
		return &SpecificRestaurantB{}
	}
	return nil
}
