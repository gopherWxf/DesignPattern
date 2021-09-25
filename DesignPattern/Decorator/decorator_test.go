package Decorator

import (
	"log"
	"os"
	"testing"
)

func TestWrapLogger(t *testing.T) {
	f := WrapLogger(pi, log.New(os.Stdout, "test ", 1))
	f(100000)
}
