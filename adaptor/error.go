package adaptor

import (
	"fmt"

	"github.com/kumarabd/gokit/errors"
)

// ErrCreateInstance is the error for install mesh
func ErrCreateInstance(err error) error {
	return errors.New("1001", fmt.Sprintf("Error creating mesh intance: %s", err.Error()))
}
