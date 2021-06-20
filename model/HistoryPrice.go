package model

type HistoryPrice struct {
	Id        int64   `json:"Id"`
	IdProduct int64   `json:"IdProduct"`
	OldPrice  float32 `json:"OldPrice"`
	Time      string  `json:"Time"`
}
