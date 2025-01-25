package models

type CheckRepoRequest struct {
	CloneURL string  `json:"clone_url" binding:"required"`
	Size     float64 `json:"size" binding:"gte=0"`
}

type CheckRepoResponse struct {
	Total int    `json:"total"`
	Files []File `json:"files"`
}
