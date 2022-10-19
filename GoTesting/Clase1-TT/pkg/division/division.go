package division

import "fmt"

func Division(num, denominator int) (int, error){
	if denominator == 0 {
		return 0, fmt.Errorf("denominator can't be 0")
	}
	return num / denominator, nil
}