package calculator

import "fmt"

func Add(a int, b int) (int, error) {
	if a == 0 || b == 0{
		return 0, fmt.Errorf("numbers can't be 0")
	}
	return a + b, nil
}

func Sub(a int, b int) int{
	return a - b
}