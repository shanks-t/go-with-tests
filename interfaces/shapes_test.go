package interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, hasArea: 7.0},
		{shape: Circle{Radius: 10}, hasArea: 31.1592653589793},
		{shape: Triangle{Base: 10, Height: 5}, hasArea: 2.0},
	}

	for _, tt := range areaTests {
		// using the tt.name from the case to use it as the 't.run' test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

}
