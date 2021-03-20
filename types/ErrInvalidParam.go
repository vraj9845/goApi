package types

import "fmt"

type ErrInvalidParam struct {
	Param []string
}

func (e ErrInvalidParam) Error() string {
	return fmt.Sprintf("Incorrect value for parameter: %v", e.Param)
}
