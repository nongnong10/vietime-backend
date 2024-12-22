package dto

// Create Fact
type CreateFactRequest struct {
	Content string `json:"content" binding:"required"`
}

type CreateFactResponse struct {
	Success bool `json:"success"`
}

// Get Fact
type GetFactResponse struct {
	Fact string `json:"fact"`
}
