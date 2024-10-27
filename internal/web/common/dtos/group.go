package dtos

// CreateGroupDTO is a struct used for group input/output in API
type CreateGroupDTO struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description"`
}

// UpdateGroupDTO is a struct used for group input/output in API
type UpdateGroupDTO struct {
	Name        string `form:"name"`
	Description string `form:"description"`
}
