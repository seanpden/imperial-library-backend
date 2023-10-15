package models

type BookContent struct {
	Content string
}

func ReadBookContent() BookContent {
	bookcontent := BookContent{"content"}
	return bookcontent
}
