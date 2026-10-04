package shape

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Rectangle{10.0, 10.0}.Perimeter()
	want := 40.0

	cmpFloat(got, want, t)
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		cmpFloat(got, want, t)
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0

		checkArea(t, rectangle, want)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})
}

func TestArea2(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12.0, Height: 6.0}, want: 36.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			cmpFloat2(got, tt.want, t, tt.shape)
		})
	}
}

func cmpFloat(got, want float64, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func cmpFloat2(got, want float64, t testing.TB, shape Shape) {
	t.Helper()
	if got != want {
		t.Errorf("%#v: got %g want %g", shape, got, want)
	}
}
