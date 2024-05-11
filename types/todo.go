package types

type Todo struct {
	ID          string `json:"id"`
	UserID      int    `json:"userID"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	IsImportant bool   `json:"isImportant"`
	IsUrgent    bool   `json:"isUrgent"`
	CategoryID  int    `json:"categoryID"`
}
