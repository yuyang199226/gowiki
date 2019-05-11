package model

type Book struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Publisher string `json:"publisher"`
}


func (*Book) TableName() string {
	return "book"
}
