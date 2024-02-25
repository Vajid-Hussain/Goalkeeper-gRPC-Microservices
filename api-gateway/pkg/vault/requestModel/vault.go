package requestmodel

import "time"

type CollecionDetails struct {
	UserID         string `json:"-"`
	CollectionName string `json:"collectionname"`
}

type Data struct {
	UserID     string `json:"-"`
	CategoryID string `json:"categoryID"`
	Data       string `json:"data"`
	Reminder   string `json:"reminder"`
}

type GetDataRequest struct {
	UserID     string `json:"-"`
	CategoryID string `josn:"categoryid"`
}

type GetDataResponse struct {
	Data      string
	Remainder time.Time
}
