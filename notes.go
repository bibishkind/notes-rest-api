package bi_notes_rest_api

type Note struct {
	Id      int
	ListId  int
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
}

type List struct {
	Id          int
	UserId      int
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
