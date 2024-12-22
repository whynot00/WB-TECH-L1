package main

import "fmt"

// LegacyRectangle - старый класс, который нужно адаптировать
type LegacyRectangle struct {
	width, height float64
}

func NewLegacyRectangle(width, height float64) *LegacyRectangle {
	return &LegacyRectangle{
		width:  width,
		height: height,
	}
}

func (l *LegacyRectangle) GetWidth() float64 {
	return l.width
}

func (l *LegacyRectangle) GetHeight() float64 {
	return l.height
}

// Shape - целевой интерфейс, который ожидается клиентом
type Shape interface {
	Area() float64
	Perimeter() float64
}

// RectangleAdapter - адаптер, преобразующий LegacyRectangle в Shape
type RectangleAdapter struct {
	rectangle *LegacyRectangle
}

func NewRectangleAdapter(rectangle *LegacyRectangle) *RectangleAdapter {
	return &RectangleAdapter{
		rectangle: rectangle,
	}
}

func (ra *RectangleAdapter) Area() float64 {
	return ra.rectangle.GetWidth() * ra.rectangle.GetHeight()
}

func (ra *RectangleAdapter) Perimeter() float64 {
	return 2 * (ra.rectangle.GetWidth() + ra.rectangle.GetHeight())
}

func main() {
	legacyRect := NewLegacyRectangle(5, 10)
	adapter := NewRectangleAdapter(legacyRect)

	fmt.Printf("Площадь: %.2f\n", adapter.Area())
	fmt.Printf("Периметр: %.2f\n", adapter.Perimeter())
}
