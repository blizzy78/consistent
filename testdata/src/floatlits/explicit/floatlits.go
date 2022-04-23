package floatlits

func floatLits() {
	_ = 0.5
	_ = -0.5
	_ = 000.5

	_ = .5  // want "add zero before decimal point in floating-point literal"
	_ = -.5 // want "add zero before decimal point in floating-point literal"
}
