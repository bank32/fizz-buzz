package main

import (
	"testing"

	r "github.com/stretchr/testify/require"
)

func TestTell(t *testing.T) {
	var flagtests = []struct {
		actual int
		expect string
		desc   string
	}{
		{1, "1", "not match"},
		{3, "Fizz", "divided by 3"},
		{5, "Buzz", "divdied by 5"},
		{15, "FizzBuzz", "divided by 3 and 5"},
	}

	for _, i := range flagtests {
		r.Equal(t, i.expect, tell(i.actual), i.desc)
	}
}
