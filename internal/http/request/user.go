package request

import (
	"fmt"

	"github.com/ptaas-tool/base-api/pkg/enum"
	"github.com/ptaas-tool/base-api/pkg/models/user"
)

type (
	UserProfileRequest struct {
		Name string `json:"username"`
		Pass string `json:"password"`
	}

	UserRegisterRequest struct {
		Name string    `json:"username"`
		Pass string    `json:"password"`
		Role enum.Role `json:"role"`
	}
)

func (u UserProfileRequest) Validate() error {
	if len(u.Name) == 0 {
		return fmt.Errorf("username cannot be empty")
	}

	if len(u.Pass) == 0 {
		return fmt.Errorf("password cannot be empty")
	}

	return nil
}

func (u UserProfileRequest) ToModel() *user.User {
	return &user.User{
		Username: u.Name,
		Password: u.Pass,
	}
}

func (u UserRegisterRequest) Validate() error {
	if len(u.Name) == 0 {
		return fmt.Errorf("username cannot be empty")
	}

	if len(u.Pass) == 0 {
		return fmt.Errorf("password cannot be empty")
	}

	return nil
}

func (u UserRegisterRequest) ToModel() *user.User {
	return &user.User{
		Username: u.Name,
		Password: u.Pass,
		Role:     u.Role,
	}
}
