package client

import "fmt"

func wrapError(errorMsg string, originalErr error) error {
	return fmt.Errorf("%s : %v", errorMsg, originalErr)
}