package utils

import (
	"strconv"

	"github.com/pkg/errors"
)

func ParseInt(s string, defaultValue int) (int, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultValue, errors.Wrapf(err, "parse %s to int failed", s)
	}

	return int(n), nil
}
