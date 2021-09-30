package ResponsibilityChain

import (
	"fmt"
	"testing"
)

func TestHandler_Handler(t *testing.T) {
	wang := NewHandler("wang", nil, 1)
	zhang := NewHandler("zhang", wang, 2)
	wu := NewHandler("wu", zhang, 3)
	r := wu.Handler(2)
	fmt.Println(r)
}
