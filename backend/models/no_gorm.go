package models

type GetMostVisitedRequest struct {
	PageIndex   int `json:"page_index"`
	ItemPerPage int `json:"item_per_page"`
}

type GetMostVisitedResponse struct {
	NextPage  int            `json:"next_page"`
	PageIndex int            `json:"page_index"`
	Users     []*UserVisited `json:"users"`
}
