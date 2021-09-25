package vra

import . "time"

type ClinetObject struct {
	Login    string `admin`
	Password string `P@ss0wrd`
	SSL      bool
	FQDN     string `none`
	Result   []Content
}

// Toke Struct
type Token struct {
	Token_Access  string `json:"access_token"`
	Refresh_Token string `json:"refresh_token"`
	Id_Token      string `json:"id_token"`
	Token_Type    string `josn:Bearer`
	Expires_in    string `json:"expires_in"`
}

type Content struct {
	Data             []Blueprint `json:"content"`
	TotalElements    int64       `json:"totalElements"`
	TotalPages       int64       `json:"totalPages"`
	NumberOfElements int64       `json:"numberOfElements"`
	Empty            bool        `json:"empty"`
	First            bool        `json:"first"`
	Last             bool        `json:"last"`
}

type Blueprint struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	CreatedBy     string `json:"createdBy"`
	CreatedAt     Time   `json:"createdAt"`
	OwnedBy       string `json:"ownedBy"`
	LastUpdatedBy string `json:"lastUpdatedBy"`
	LastUpdatedAt Time   `json:"lastUpdatedAt"`
	// Статусы в системе
	// CREATE_SUCCESSFUL,
	// CREATE_INPROGRESS,
	// CREATE_FAILED,
	// UPDATE_SUCCESSFUL,
	// UPDATE_INPROGRESS,
	// UPDATE_FAILED,
	// DELETE_SUCCESSFUL,
	// DELETE_INPROGRESS,
	// DELETE_FAILED
	CatalogData Catalog `json:"catalog"`
	Status      string  `json:"status"`
}

// /
type Catalog struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Version     string `json:"version"`
	Name        string `json:"name"`
}
