package ginhandlers

import "leisure_time/internal/domain/dto"

// User responses
type userListResponse struct {
	Users []dto.UserDTO
}
