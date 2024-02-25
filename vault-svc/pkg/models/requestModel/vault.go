package requestmodel

type Colleciton struct {
	CategoryName string
	UserID       string
}

type Data struct {
	UserID     string
	CategoryID string
	Data       string
	Reminder   string
}

type GetDataRequest struct {
	UserID     string
	CategoryID string
}

