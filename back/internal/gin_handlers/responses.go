package ginhandlers

import "leisure_time/internal/domain/models"

// User responses
type userListResponse struct {
	Users []models.User
}
