package quotes

// Quote describes a quote.
type Quote struct {
	ID      int    `json:"id_quote"`
	Content string `json:"content"`
	Score   int    `json:"score"`
	UUID    string `json:"uuid"`
}
