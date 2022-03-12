package models

type GetVisitedUsersRequest struct {
	LastId      int `json:"last_id"`
	ItemPerPage int `json:"item_per_page"`
}

type GetVisitedUsersResponse struct {
	HasMoreData bool           `json:"has_more_data"`
	Users       []*UserVisited `json:"users"`
}

type CntVisitTimesByIp struct {
	Ip  string `json:"ip"`
	Num int    `json:"num"`
}

type MostVisitedUserResponse struct {
	Users []*CntVisitTimesByIp `json:"users"`
}
