package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
/* -------------------------------- Add Tests ------------------------------- */
func TestAddCorrect (t *testing.T){
	//arrange - valores de entrada y resultado esperado
	num1 := 10
	num2 := 5
	expected := 15

	//act - llamar a la funcion
	result, err := Add(num1, num2)

	//assert - validar resultado esperado con resultado actual
	//el mensaje es opcional
	assert.Nil(t, err)
	assert.Equal(t, expected, result, "the result %d is different than expected %d", result, expected)
	/* if result != expected {
		t.Errorf("the result %d is different than expected %d", result, expected)
	} */

}
func TestAddIncorrect (t *testing.T){
	//Arrange
	num1 := 0
	num2 := 3
	expectedError := "numbers can't be 0"
	//act
	_, err := Add(num1, num2)
	//assert
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, expectedError)
}

/* -------------------------------- Sub Tests ------------------------------- */
func TestSubCorrect (t *testing.T){
	//arrange
	num1 := 9
	num2 := 2
	expected := 7

	//act 
	result := Sub(num1, num2)

	//assert
	assert.Equal(t, expected, result, "the result %d is different than expected %d", result, expected)

}

