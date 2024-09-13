package shapes

import strategy "github.com/backend-school/go-patterns/internal/patterns/behavioral/strategy/example2"

type TextSquare struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
