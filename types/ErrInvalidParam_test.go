package types

import (
	"testing"
)

func TestErrInvalidParam_Error(t *testing.T) {
	err := ErrInvalidParam{Param: []string{"OrganizationId", "storeId"}}
	expected := "Incorrect value for parameter: [OrganizationId storeId]"
	if err.Error() != expected {
		t.Errorf("Failed Expected %v Got %v", expected, err)
	}
}
