package paste

type Paste struct {
	Code      string `json:"code"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}
