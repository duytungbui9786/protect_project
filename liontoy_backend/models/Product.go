package models

type Product struct {
	Id           int64
	idbrand      int64   `json:"idbrand"`
	Name         string  `json:"Name"`
	AgesRank     int     `json:"AgesRank`
	Newprice     float32 `json:"Newprice"`
	Image        string  `json:"Image"`
	Inforgeneral string  `json:"Inforgeneral`
	Infor        string  `json:"Infor`
	Featured     int     `json:"Featured`
	New          int     `json:"New`
	Comming      int     `json:"Comming`
	CreatedAt    string  `json:"CreatedAt"`
	UpdatedAt    string  `json:"ModifiedAt"`
}
