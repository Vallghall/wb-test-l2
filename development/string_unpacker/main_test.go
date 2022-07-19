package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnpack(t *testing.T) {
	initial := "a4bc2d5e"
	expected := "aaaabccddddde"
	res, err := unpack(initial)

	require.NoError(t, err)
	require.Equal(t, expected, res)
}

func TestUnpack_EscapeSuccess(t *testing.T) {
	initial := `qwe\45\\`
	expected := `qwe44444\`
	res, err := unpack(initial)

	require.NoError(t, err)
	require.Equal(t, expected, res)
}

func TestUnpack_InvalidString_FirstEqNumber(t *testing.T) {
	initial := "1bc"
	_, err := unpack(initial)
	require.Error(t, err)
}

func TestUnpack_InvalidString_TwoConsequentNumbers(t *testing.T) {
	initial := "a3b45"
	_, err := unpack(initial)
	require.Error(t, err)
}

func TestUnpack_InvalidString_BadEscape(t *testing.T) {
	initial := `a4c\t`
	_, err := unpack(initial)
	require.Error(t, err)
}
