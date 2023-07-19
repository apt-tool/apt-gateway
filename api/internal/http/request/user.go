package request

import "fmt"

type Login struct {
	Email string `json:"email"`
}

func (l *Login) Validate() error {
	if len(l.Email) == 0 {
		return fmt.Errorf("email cannot be empty")
	}

	return nil
}
