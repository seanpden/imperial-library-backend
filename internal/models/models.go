package models

type BookContent struct {
	Link_To_Content string `json:"link_to_content"`
	Author          string `json:"author"`
	Comment         string `json:"comment"`
	Text            string `json:"text"`
}

func (b *BookContent) TableName() string {
	return "book_content"
}

type BookInfo struct {
	Title           string `json:"title"`
	Link_To_Content string `json:"link_to_content"`
	Author          string `json:"author"`
	Summary         string `json:"summary"`
}

func (b *BookInfo) TableName() string {
	return "book_info"
}

type PGDBConfig struct {
	DATABASE_URL string
	PGDATABASE   string
	PGHOST       string
	PGPASSWORD   string
	PGPORT       string
	PGUSER       string
}

type ConnectionString struct {
	CONNECTIONSTRING string
}
