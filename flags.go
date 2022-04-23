package consistent

import (
	"fmt"
	"regexp"
	"strings"
)

type enumValue struct {
	allowed []string
	value   string
}

type regexpValue struct {
	r *regexp.Regexp
	s string
}

type invalidValueError string

func (e *enumValue) Set(value string) error {
	for _, a := range e.allowed {
		if a == value {
			e.value = value
			return nil
		}
	}

	return invalidValueError(value)
}

func (e *enumValue) String() string {
	return e.value
}

func (e enumValue) description(desc string) string {
	return desc + " (" + strings.Join(e.allowed, "/") + ")"
}

func newRegexpValue(s string) regexpValue {
	return regexpValue{
		r: regexp.MustCompile(s),
		s: s,
	}
}

func (r *regexpValue) Set(value string) error {
	if value == "" {
		r.r = nil
		r.s = ""

		return nil
	}

	var err error
	r.r, err = regexp.Compile(value)
	r.s = value

	if err != nil {
		return fmt.Errorf("%s: parse regexp: %w", value, err)
	}

	return nil
}

func (r *regexpValue) String() string {
	if r.r == nil {
		return ""
	}

	return r.s
}

func (e invalidValueError) Error() string {
	return string(e) + ": invalid value"
}
