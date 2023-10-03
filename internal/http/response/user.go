package response

import (
	"time"

	"github.com/ptaas-tool/base-api/pkg/enum"
	"github.com/ptaas-tool/base-api/pkg/models/user"
)

type (
	Token struct {
		Token string `json:"token"`
	}

	UserResponse struct {
		ID        uint      `json:"id"`
		Username  string    `json:"username"`
		Role      enum.Role `json:"role"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func (u UserResponse) DTO(user *user.User) *UserResponse {
	u.ID = user.ID
	u.Username = user.Username
	u.Role = user.Role
	u.CreatedAt = user.CreatedAt

	return &u
}
