package Factory

import "testing"

func TestNewRestaurant(t *testing.T) {
	NewRestaurant("A").GetFood()
	NewRestaurant("B").GetFood()
}
