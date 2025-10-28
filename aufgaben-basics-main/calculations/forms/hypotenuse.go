package forms

import "math"

// Erwartet die Längen der Katheten eines rechtwinkligen Dreiecks.
// Liefert die Länge der Hypotenuse.
func Hypotenuse(a, b float64) float64 {
	var c = ((a * a) + (b * b))
	var d = math.Sqrt(c)

	return d
}
