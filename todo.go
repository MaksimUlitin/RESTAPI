package restApi

type TodoList struct {
	id          int    `json:"id" :"id"`
	title       string `json:"title" :"title"`
	Description string `json:"description" :"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
