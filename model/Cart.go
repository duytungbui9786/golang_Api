package model

type Cart struct {
	Id         int64  `json:"Id"`
	CategoryId int64  `json:"CategoryId"`
	Image      string `json:"Image"`
	Name       string `json:"Name"`
	Price      string `json:"Price"`
	Quantity   int64  `json:"Quantity"`
	CreatedAt  string `json:"CreatedAt"`
	ModifiedAt string `json:"ModifiedAt"`
}
