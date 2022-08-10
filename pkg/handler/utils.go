package handler

import (
	"errors"
	"strconv"
)

func parseLimitAndOffset(limitString, offsetString string) (int, int, error) {

	var limitInt int

	if limitString == "" {
		limitInt = -1
	} else {
		var err error
		limitInt, err = strconv.Atoi(limitString)
		if err != nil {
			return 0, 0, errors.New("bad query")
		}
	}

	var offsetInt int

	if offsetString == "" {
		offsetInt = 0
	} else {
		var err error
		offsetInt, err = strconv.Atoi(offsetString)
		if err != nil {
			return 0, 0, errors.New("bad query")
		}
	}

	return limitInt, offsetInt, nil
}
