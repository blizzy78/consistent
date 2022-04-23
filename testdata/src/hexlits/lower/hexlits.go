package hexlits

func hexLitLower() {
	_ = 0xff
}

func hexLitUpper() {
	_ = 0xFF // want "use lowercase digits in hex literal"
}
