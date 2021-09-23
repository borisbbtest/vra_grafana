package vra

// Toke Struct
type Token struct {
	Token_Access  string `json:"access_token"`
	Refresh_Token string `json:"refresh_token"`
	Id_Token      string `json:"id_token"`
	Token_Type    string `josn:Bearer`
	Expires_in    string `json:"expires_in"`
}

type Content struct {
	Data          []Blueprint `json:"content"`
	TotalElements int64       `json:"totalElements"`
	TotalPages    int64       `json:"totalPages"`
}

type Blueprint struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	CreatedBy     string    `json:"createdBy"`
	CreatedAt     string    `json:"createdAt"`
	OwnedBy       string    `json:"ownedBy"`
	LastUpdatedBy string    `json:"lastUpdatedBy"`
	LastUpdatedAt string    `json:"lastUpdatedAt"`
	Status        string    `json:"status"`
	Data          []Catalog `json:"catalog"`
}

type Catalog struct {
	Description string `json:"description"`
	id          string `json:"id"`
	version     string `json:"version"`
	Name        string `json:"name"`
}
