package taskmanager

type TodoList struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	ID     int64 `json:"id"`
	ListID int64 `json:"list_id"`
	UserID int64 `json:"user_id"`
}

type TodoItem struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	ID     int64 `json:"id"`
	ListID int64 `json:"list_id"`
	ItemID int64 `json:"item_id"`
}
