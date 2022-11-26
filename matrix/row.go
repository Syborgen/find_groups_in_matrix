package matrix

import (
	"strconv"
	"strings"
)

type Row []int

func (r Row) Stirng() string {
	rowAsStringsSlice := make([]string, 0, len(r))
	for _, element := range r {
		rowAsStringsSlice = append(rowAsStringsSlice, strconv.Itoa(element))
	}

	return strings.Join(rowAsStringsSlice, " ")
}
