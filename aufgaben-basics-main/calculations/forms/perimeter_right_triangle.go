package forms

import "math"

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func PerimeterRightTriangle(a, b float64) float64 {
	var c = (a*a + b*b)
	var d = math.Sqrt(c)
	var e = a + b + d
	return e
}
