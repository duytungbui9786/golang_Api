package model

type Product struct {
	Id         int64
	CategoryId int64   `json:"CategoryId"`
	Name       string  `json:"Name"`
	Price      float32 `json:"Price"`
	Image      string  `json:"Image"`
	IsSale     bool    `json:"IsSale"`
	CreatedAt  int64   `json:"CreatedAt"`
	ModifiedAt int64   `json:"ModifiedAt"`
	Rating     float32 `json:"rating"`
}
