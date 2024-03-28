package types

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	UserID      int    `json:"userID"`
	IsCompleted bool   `json:"isCompleted"`
}
