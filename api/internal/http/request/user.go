package request

import "fmt"

type UserRegister struct {
	Name string `json:"username"`
	Pass string `json:"password"`
}

func (u UserRegister) Validate() error {
	if len(u.Name) == 0 {
		return fmt.Errorf("username cannot be empty")
	}

	if len(u.Pass) == 0 {
		return fmt.Errorf("password cannot be empty")
	}

	return nil
}
