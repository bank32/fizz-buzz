package main

import (
	"testing"

	r "github.com/stretchr/testify/require"
)

func TestTell(t *testing.T) {
	// test with not matching number
	r.Equal(t, "1", tell(1))
	// test with divinded 3 number
	r.Equal(t, "Fizz", tell(3))
	// test with divinded 5 number
	r.Equal(t, "Buzz", tell(5))
	// test with divinded 3 and 5 number
	r.Equal(t, "FizzBuzz", tell(15))
}
