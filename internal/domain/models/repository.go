package models

type Repository struct {
	Id       int    `json:"id"`
	NodeId   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login string `json:"login"`
		Id    int    `json:"id"`
	} `json:"owner"`
	HtmlUrl     string      `json:"html_url"`
	Description interface{} `json:"description"`
}
