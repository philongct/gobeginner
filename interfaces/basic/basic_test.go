package basic

import (
	"os"
	"testing"
)

// go test -run TestBasic
func TestBasic(t *testing.T) {
	// With interface, Pointer receiver merely accepts pointer, none-pointer receiver accepts both
	shapes := []Shape{&Rect{}, Circle{}, &Circle{}}
	// Error: Rect does not implement Shape (Show method has pointer receiver)
	//shapes := []Shape{Rect{}, Circle{}}
	for _, s := range shapes {
		// Only pointer receiver retain changes
		s.SetColor("Blue")
		s.Render(os.Stdout)
	}
}

// go test -run TestBasic
func TestInterfacePointer(t *testing.T) {
	// We can use pointer interface
	render := func(s *Shape, out Output) {
		(*s).SetColor("Blue")
		(*s).Render(out)
	}
	shapes := []Shape{&Rect{}, Circle{}}
	for _, s := range shapes {
		render(&s, os.Stdout)
	}
}
