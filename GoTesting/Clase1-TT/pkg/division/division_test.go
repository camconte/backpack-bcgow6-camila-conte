package division

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisionWithZero (t *testing.T){
	//arrange
	num := 10
	denominator := 0
	expectedError := "denominator can't be 0"
	//act
	_, err := Division(num, denominator)
	//assert
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, expectedError)
}