package render

import shapes "gocligui/shapes/rectangle"

type Renderer struct {
	Children []shapes.Rectangle
}

func New() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Render() {
	for _, c := range r.Children {
		c.Render()
	}
}

func (r *Renderer) AddChild(child shapes.Rectangle) {
	r.Children = append(r.Children, child)
}

func (r *Renderer) RemoveChild(index int) {
	r.Children = append(r.Children[:index], r.Children[index+1:]...)
}
