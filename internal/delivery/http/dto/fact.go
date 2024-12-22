package dto

// Create Fact
type CreateFactRequest struct {
	Content string `json:"content" binding:"required"`
}

type CreateFactResponse struct {
	Success bool `json:"success"`
}
