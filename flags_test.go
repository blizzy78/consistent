package consistent

import (
	"testing"

	"github.com/matryer/is"
)

func TestEnumValue_Set(t *testing.T) {
	is := is.New(t)

	val := enumValue{
		allowed: []string{"foo"},
	}

	_ = val.Set("foo")

	is.Equal(val.value, "foo")
}

func TestEnumValue_Set_Invalid(t *testing.T) {
	is := is.New(t)

	val := enumValue{
		allowed: []string{"foo"},
	}

	err := val.Set("bar")

	is.True(err != nil)
}

func TestEnumValue_String(t *testing.T) {
	is := is.New(t)

	val := enumValue{
		allowed: []string{"foo"},
	}

	_ = val.Set("foo")

	is.Equal(val.String(), "foo")

	is.Equal((&enumValue{}).String(), "")
}

func TestRegexpValue_Set(t *testing.T) {
	is := is.New(t)

	val := regexpValue{}
	_ = val.Set("^[a-z]+$")

	is.True(val.r.MatchString("foo"))
	is.Equal(val.s, "^[a-z]+$")
}

func TestRegexpValue_Invalid(t *testing.T) {
	is := is.New(t)

	val := regexpValue{}

	err := val.Set("[")

	is.True(err != nil)
}

func TestRegexpValue_String(t *testing.T) {
	is := is.New(t)

	val := regexpValue{}
	_ = val.Set("^[a-z]+$")

	is.Equal(val.String(), "^[a-z]+$")

	is.Equal((&regexpValue{}).String(), "")
}
