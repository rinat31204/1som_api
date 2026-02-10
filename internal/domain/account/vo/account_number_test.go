package vo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateAccountNumber_Format(t *testing.T) {
	for i := 0; i < 1000; i++ {
		ls, err := GenerateAccountNumber()
		require.NoError(t, err)
		require.Regexp(t, `^[A-Z][0-9]{2}-[A-Z][0-9]{2}$`, ls)
	}
}

func TestNormalizeAccountNumber(t *testing.T) {
	tests := map[string]string{
		"v84-h83":   "V84-H83",
		" V84–H83 ": "V84-H83",
		"v84—h83":   "V84-H83",
	}

	for input, expected := range tests {
		ls, err := NormalizeAccountNumber(input)
		require.NoError(t, err)
		require.Equal(t, expected, ls)
	}
}

func TestNormalizeAccountNumber_Invalid(t *testing.T) {
	invalid := []string{
		"",
		"V8-H83",
		"V84H83",
		"123-456",
		"VVV-123",
	}

	for _, v := range invalid {
		_, err := NormalizeAccountNumber(v)
		require.Error(t, err)
	}
}
