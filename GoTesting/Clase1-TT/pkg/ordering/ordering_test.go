package ordering

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderIntSlice (t *testing.T){
	//arrange
	numberSlice := []int{34, 6, 7, 1, 67, 101, 2}
	expected := []int{1, 2, 6, 7, 34, 67, 101}
	//act
	result := OrderIntSlice(numberSlice...)
	//assert
	assert.Equal(t, result, expected)
}