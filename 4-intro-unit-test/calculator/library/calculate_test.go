package library

import "testing"

func TestPenambahan(t *testing.T) {
	t.Run("Test Fungsi Penambahan dengan 2 bilangan positif", func(t *testing.T) {
		a := 15
		b := 20
		expected := 35
		actual := Penambahan(a, b)
		if actual != expected {
			t.Error("hasil penambahan tdak sesuai. actual=", actual)
		}

	})

	t.Run("Test Fungsi Penambahan dengan 2 bilangan negatif", func(t *testing.T) {
		a := -15
		b := -20
		expected := -35
		actual := Penambahan(a, b)
		if actual != expected {
			t.Error("hasil penambahan tdak sesuai. actual=", actual)
		}

	})
}
