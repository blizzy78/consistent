package hexlits

func hexLitLower() {
	_ = 0xff // want "use uppercase digits in hex literal"
}

func hexLitUpper() {
	_ = 0xFF
}
