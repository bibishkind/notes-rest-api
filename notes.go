package bi_notes_rest_api

type NoteItem struct {
	Id      int
	ListId  int
	Title   string
	Content string
}

type NoteList struct {
	Id          int
	Title       string
	Description string
}
