package types

type SharedTodo struct {
	ID               int `json:"id"`
	TodoID           int `json:"todoID"`
	AuthorizedUserID int `json:"authorizedUserID"`
}
