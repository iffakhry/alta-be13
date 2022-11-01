package library

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	t.Run("Test Fungsi Penambahan dengan 1 bilangan negatif dan 1 bilangan positif", func(t *testing.T) {
		a := -15
		b := 20
		// expected := 5
		actual := Penambahan(a, b)
		assert.Equal(t, 5, actual, "hasil penambahan tidak sesuai")

	})

}

func TestKurang(t *testing.T) {
	t.Run("Test Fungsi Pengurangan dengan 2 bilangan positif", func(t *testing.T) {
		a := 30
		b := 20
		expected := 10
		actual := Pengurangan(a, b)
		assert.Equal(t, expected, actual, "hasil pengurangan tidak sesuai")
	})

	t.Run("Test Fungsi Pengurangan dengan 2 bilangan negatif", func(t *testing.T) {
		a := -30
		b := -20
		expected := -10
		actual := Pengurangan(a, b)
		assert.Equal(t, expected, actual, "hasil pengurangan tidak sesuai")
	})
}

func TestPembagian(t *testing.T) {
	t.Run("Test Fungsi Pembagian dengan 2 bilangan positif", func(t *testing.T) {
		a := 10
		b := 2
		expected := 5
		actual := Pembagian(a, b)
		assert.Equal(t, expected, actual, "hasil pembagian tidak sesuai")
	})
}

func TestKonversiNilai(t *testing.T) {
	t.Run("Test Fungsi Konversi Nilai dengan nilai = 90", func(t *testing.T) {
		data := 90
		expected := "A"
		actual := KonversiNilai(data)
		assert.Equal(t, expected, actual, "hasil konversi tidak sesuai")
	})

	t.Run("Test Fungsi Konversi Nilai dengan nilai = 80", func(t *testing.T) {
		data := 80
		expected := "B+"
		actual := KonversiNilai(data)
		assert.Equal(t, expected, actual, "hasil konversi tidak sesuai")
	})
}
