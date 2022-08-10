package bi_notes_rest_api

type Note struct {
	Id      int
	ListId  int
	Title   string
	Content string
}

type List struct {
	Id          int
	UserId      int
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}
