package requestmodel

type CollecionDetails struct {
	UserID         string `json:"userid"`
	CollectionName string `json:"collectionname"`
}

type Data struct {
	UserID     string `json:"userid"`
	CategoryID string `json:"categoryID"`
	Data       string `json:"data"`
	Reminder   string `json:"reminder"`
}
