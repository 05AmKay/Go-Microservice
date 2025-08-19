package helpers

import (
	"fmt"
	"strconv"
)

func ConvertStringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("error converting string to int: %w", err)
	}
	return i, nil
}
