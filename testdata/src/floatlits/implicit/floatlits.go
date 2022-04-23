package floatlits

func floatLits() {
	_ = 0.5   // want "remove zero before decimal point in floating-point literal"
	_ = -0.5  // want "remove zero before decimal point in floating-point literal"
	_ = 000.5 // want "remove zero before decimal point in floating-point literal"

	_ = .5
	_ = -.5
}
