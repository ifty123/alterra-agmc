package model

type Books struct {
	Id        int    `json:"id_book"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}
