package models

type Category struct {
	Rev       string `json:"_rev"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Order     string `json:"order"`
	ParentID  string `json:"parent_id"`
	SargaID   string `json:"sarga_id"`
	Slug      string `json:"slug"`
	Weight    string `json:"weight"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
