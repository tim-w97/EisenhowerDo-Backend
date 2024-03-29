package types

type Todo struct {
	ID          int    `json:"id"`
	UserID      int    `json:"userID"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	Position    int    `json:"position"`
	IsCompleted bool   `json:"isCompleted"`
}
