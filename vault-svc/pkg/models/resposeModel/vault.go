package resposemodel

import "time"

type DataSet struct {
	ID         string    `bson:"_id"`
	CategoryID string    `bson:"categoryID"`
	Data       string    `bson:"data"`
	Expire     time.Time `bson:"expire"`
}

type Collections struct {
	CategoryName string `bson:"CategoryName"`
	ID           string `bson:"_id"`
}

type GetDataResponse struct{
	Data string	`bson:"data"`
	Reminder time.Time `bson:"expire"`
}
