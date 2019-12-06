package core

// TODO add serialization tags
type Tweet struct {
	ID        int    `json:"id"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}
