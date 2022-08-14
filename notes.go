package notes_rest_api

type Note struct {
	Id      int    `json:"id"`
	ListId  int    `json:"listId"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

type List struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userId"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
