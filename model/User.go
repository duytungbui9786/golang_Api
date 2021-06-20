package model

type User struct {
	Id         int64  `json:Id `
	FirstName  string `json:FirstName `
	LastName   string `json:LastName `
	Username   string `json:Username `
	Email      string `json:Email `
	Password   string `json:Password `
	Avatar     string `json:Avatar `
	Gender     string `json:Gender `
	Phone      string `json:Phone `
	Birthday   string `json:Birthday `
	Status     bool   `json:Status`
	CreatedAt  int64  `json:CreatedAt `
	ModifiedAt int64  `json:ModifiedAt `
}
