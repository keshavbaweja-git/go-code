package simplestringindicies

import (
	"errors"
)

// Solution returns the index of corresponding closing brace given
// index of an opening index
func Solution(str string, idx uint) (uint, error) {
	if str[idx] != '(' {
		return uint(len(str)), errors.New("Not an opening brace")
	}

	counter := 0
	for i := idx + 1; i < uint(len(str)); i++ {
		if str[i] == '(' {
			counter++
		}

		if str[i] == ')' {
			if counter == 0 {
				return i, nil
			} else {
				counter--
			}
		}
	}
	return uint(len(str)), errors.New("No matching closing brace")
}
