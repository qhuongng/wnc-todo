package model

type TodoRequest struct {
	Content   string `json:"content" binding:"required"`
	Completed bool   `json:"completed" binding:"required"`
	UserId    int64  `json:"user_id" binding:"required"`
}
