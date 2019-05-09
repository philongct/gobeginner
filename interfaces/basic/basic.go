package basic

type Rect struct {
	Color string
}

type Circle struct {
	Color string
}

func (r *Rect) Render(out Output) {
	out.WriteString((*r).Color + " Rect\n")
}

func (r *Rect) SetColor(color string) {
	// same
	//(*r).Color = color
	r.Color = color
}

func (c Circle) Render(out Output) {
	out.WriteString(c.Color + " Circle\n")
}

func (c Circle) SetColor(color string) {
	c.Color = color
}

// This doesn't compiles
//func (s *Shape) Clear() {
//}

type Shape interface {
	Render(out Output)
	SetColor(color string)
}

type Output interface {
	WriteString(text string) (int, error)
}
