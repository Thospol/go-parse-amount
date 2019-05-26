package amount

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCASE_001AMT_SUCCESS(t *testing.T) {
	amts := []string{"1", "10", "100", "1000", "1000.00", "10,000", "10,000.00", "100,000", "100,000.00", "1,000,000", "1,000,000.00", "10,000,000", "10,000,000.00", "100,000,000", "100,000,000.00"}
	for _, amt := range amts {
		_, err := ParseReqAmount(amt)
		assert.Equal(t, nil, err)
	}
}

func TestCASE_002AMT_FAIL(t *testing.T) {
	amts := []string{"01", "001", "1000ABC", "1000.50", "1000.05", "1000.00.00", "1000ABC.00", "0,1000", ",1000", "10,00", "100,0", "1000,", "1000,0", "100,00", "1,0000", "1,000ACB", "1,000.00ABC", "1,000.00.ABC"}
	errExpect := errors.New("9901")
	for _, amt := range amts {
		_, err := ParseReqAmount(amt)
		assert.Equal(t, errExpect, err)
	}
}
