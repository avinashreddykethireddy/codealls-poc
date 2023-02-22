package subber

import "testing"

func TestSub(t *testing.T) {
	sum := Sub(2, 2)
	expected := 0

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestCheckValues(t *testing.T) {
	cases := []struct {
		x        int
		y        int
		expected string
	}{
		{2, 5, "less"},
	}
	for _, c := range cases {
		t.Run("testing", func(t *testing.T) {
			got := CheckValues(c.x, c.y)
			if got != c.expected {
				t.Fatalf("Wanted %v , but got %v", c.expected, got)
			}
		})
	}
}
