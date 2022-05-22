package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	check := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("test for square", func(t *testing.T) {
		got := Perimeter("square", 10.0)
		want := 40.0
		check(t, got, want)
	})
	t.Run("test for rectangle", func(t *testing.T) {
		got := Perimeter("rectangle", 10.0, 20.0)
		want := 60.0
		check(t, got, want)
	})
	t.Run("test for right triangle", func(t *testing.T) {
		got := Perimeter("right triangle", 3.0, 4.0)
		want := 12.0
		check(t, got, want)
	})
	t.Run("test for another cases", func(t *testing.T) {
		got := Perimeter("else", 3.0, 4.0, 10.0)
		want := 0.0
		check(t, got, want)
	})
}

func TestArea(t *testing.T) {
	check := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("test for square", func(t *testing.T) {
		got := Area("square", 10.0)
		want := 100.0
		check(t, got, want)
	})
	t.Run("test for rectangle", func(t *testing.T) {
		got := Area("rectangle", 10.0, 20.0)
		want := 200.0
		check(t, got, want)
	})
	t.Run("test for right triangle", func(t *testing.T) {
		got := Area("right triangle", 3.0, 4.0)
		want := 6.0
		check(t, got, want)
	})
	t.Run("test for another cases", func(t *testing.T) {
		got := Area("else", 3.0, 4.0, 10.0)
		want := 0.0
		check(t, got, want)
	})
}
