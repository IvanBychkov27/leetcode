package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isNumber_OK(t *testing.T) {
	data := []string{"46.e3", ".1", "0", "2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}

	var res bool
	for _, s := range data {
		res = isNumber(s)
		assert.Equal(t, true, res)
		if !res {
			fmt.Printf("error: %q\n", s)
		}
	}
}

func Test_isNumber_NOT(t *testing.T) {
	data := []string{"-E3", "+e3", "-.", "+.", "6+1", ".e1", ".", "-", "+", "e", "abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", " 95a54e53"}

	var res bool
	for _, s := range data {
		res = isNumber(s)
		assert.Equal(t, false, res)
		if res {
			fmt.Printf("error: %q\n", s)
		}
	}

}
