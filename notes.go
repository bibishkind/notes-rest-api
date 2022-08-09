package bi_notes_rest_api

type Note struct {
	Id      int
	ListId  int
	Title   string
	Content string
}

type List struct {
	Id          int
	Title       string
	Description string
}
