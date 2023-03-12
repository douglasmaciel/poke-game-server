package types

import (
	"fmt"
)

type string50 string

func NewString50(s string) (string50, error) {
	const l = 50
	if len(s) > l {
		return "", fmt.Errorf("max length is %d", l)
	}

	return string50(s), nil
}
