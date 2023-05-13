package main //nolint:revive // don't need docs here

import (
	"fmt"
	"os"
	"text/template"
)

const (
	rangeChecksLeft    = "left"
	rangeChecksCenter  = "center"
	rangeChecksRight   = "right"
	rangeChecksOutside = "outside"
)

const templateText = `package consistent

// code generated by "go generate" - DO NOT EDIT

var rangeExprStyles = map[uint16]string{
	{{- range $bits, $style := .Styles}}
		{{$bits}}: "{{$style}}",
	{{- end}}
}`

func main() {
	if len(os.Args) != 2 {
		panic("usage: rangeexprstyles OUTFILE")
	}

	if err := run(os.Args[1]); err != nil {
		panic(err)
	}
}

func run(path string) error {
	tmpl, err := template.New("").Parse(templateText)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	data := struct {
		Styles map[uint16]string
	}{
		Styles: buildStyles(),
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("%s: create: %w", path, err)
	}

	closed := false

	defer func() {
		if closed {
			return
		}

		_ = file.Close()
	}()

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("execute: %w", err)
	}

	closed = true

	if err := file.Close(); err != nil {
		return fmt.Errorf("%s: close: %w", path, err)
	}

	return nil
}

func buildStyles() map[uint16]string {
	// bits: AND OR LEFT_LESS LEFT_GREATER RIGHT_LESS RIGHT_GREATER X_LEFT X_CENTER X_RIGHT X_OUTSIDE
	var rangeExprStyles = map[uint16]string{}

	// x > low && x < high
	rangeExprStyles[0b10_01_10_1000] = rangeChecksLeft

	// x < low || x > high
	rangeExprStyles[0b01_10_01_1000] = rangeChecksLeft

	newStyles := map[uint16]string{}

	for oldBits := range rangeExprStyles {
		// build variants for center:
		// x > low && x < high  ->  low < x && x < high
		newBits := oldBits & 0b11_00_11_0000
		newBits |= ^oldBits & 0b00_11_00_0000
		newBits |= 0b00_00_00_0100
		newStyles[newBits] = rangeChecksCenter

		// build variants for right:
		// x > low && x < high  ->  low < x && high > x
		newBits = oldBits & 0b11_00_00_0000
		newBits |= ^oldBits & 0b00_11_11_0000
		newBits |= 0b00_00_00_0010
		newStyles[newBits] = rangeChecksRight

		// build variants for outside:
		// x > low && x < high  ->  x > low && high > x
		newBits = oldBits & 0b11_11_00_0000
		newBits |= ^oldBits & 0b00_00_11_0000
		newBits |= 0b00_00_00_0001
		newStyles[newBits] = rangeChecksOutside
	}

	for k, v := range newStyles {
		rangeExprStyles[k] = v
	}

	for k := range newStyles {
		delete(newStyles, k)
	}

	// build variants where left and right are swapped:
	// x > low && x < high  ->  x < high && x > low
	for oldBits, oldStyle := range rangeExprStyles {
		newBits := oldBits & 0b11_00_00_1111
		newBits |= oldBits & 0b00_11_00_0000 >> 2
		newBits |= oldBits & 0b00_00_11_0000 << 2
		newStyles[newBits] = oldStyle
	}

	for k, v := range newStyles {
		rangeExprStyles[k] = v
	}

	return rangeExprStyles
}
