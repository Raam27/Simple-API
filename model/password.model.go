package model

type Password struct {
	Password string `json:"password" form:"password" query:"password" xml:"password"`
}
