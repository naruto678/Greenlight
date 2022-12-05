package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	json_value := strconv.Quote(fmt.Sprintf("%d mins", r))
	return []byte(json_value), nil
}
