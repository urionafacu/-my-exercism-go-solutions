package lasagna

import "testing"

func TestRemainingOvenTime(t *testing.T) {
	cases := []struct {
		param    int
		expected int
	}{
		{10, 30},
		{20, 20},
		{60, 0},
		{-10, 40},
		{1, 39},
	}

	for _, value := range cases {
		if observed := RemainingOvenTime(value.param); observed != value.expected {
			t.Fatalf("RamainingOvenTime() = %v, want %v", observed, value.expected)
		}
	}
}

func TestPreparationTime(t *testing.T) {
	cases := []struct {
		param    int
		expected int
	}{
		{2, 4},
		{1, 2},
		{0, 0},
		{10, 20},
		{-2, 0},
	}

	for _, value := range cases {
		if observed := PreparationTime(value.param); observed != value.expected {
			t.Fatalf("PreparationTime() = %v, want %v", observed, value.expected)
		}
	}
}

func TestElapsedTime(t *testing.T) {
	cases := []struct {
		numberOfLayers      int
		actualMinutesInOven int
		expected            int
	}{
		{4, 30, 38},
		{1, 8, 10},
	}

	for _, value := range cases {
		if observed := ElapsedTime(value.numberOfLayers, value.actualMinutesInOven); observed != value.expected {
			t.Fatalf("ElapsedTime(%v, %v) = %v, want %v", value.numberOfLayers, value.actualMinutesInOven, observed, value.expected)
		}
	}
}
