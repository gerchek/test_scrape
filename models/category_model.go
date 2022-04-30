package models

type Category struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Order    string `json:"order"`
	ParentID string `json:"parent_id"`
	SargaID  string `json:"sarga_id"`
	Slug     string `json:"slug"`
	Weight   string `json:"weight"`
}
