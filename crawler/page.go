package crawler

type Page struct {
	ID    int64  `db:"id"`
	Title string `db:"page_title"`
	Url   string `db:"page_url"`
	Html  string `db:"page_html"`
}
