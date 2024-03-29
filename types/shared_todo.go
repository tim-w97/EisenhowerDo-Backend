package types

type SharedTodo struct {
	ID          int `json:"id"`
	TodoID      int `json:"todoID"`
	OtherUserID int `json:"otherUserID"`
}
