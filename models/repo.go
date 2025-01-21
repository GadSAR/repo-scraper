package models

type CheckRepoRequest struct {
	CloneURL string `json:"clone_url" binding:"required"`
	Size     int    `json:"size" binding:"gte=0"`
}

type CheckRepoResponse struct {
	Total int    `json:"total"`
	Files []File `json:"files"`
}
