package entities

type BookCore struct {
	ID          uint
	Title       string
	Publisher   string
	Author      string
	PublishYear string
	UserID      uint
	User        UserCore
}

type BookRequest struct {
	Title       string `json:"title" form:"title"`
	Publisher   string `json:"publisher" form:"publisher"`
	Author      string `json:"author" form:"author"`
	PublishYear string `json:"publish_year" form:"publish_year"`
	UserID      uint   `json:"user_id" form:"user_id"`
}

type BookResponse struct {
	ID          uint   `json:"title"`
	Title       string `json:"title"`
	Publisher   string `json:"publisher"`
	Author      string `json:"author"`
	PublishYear string `json:"publish_year"`
	UserID      uint   `json:"user_id"`
}

// mapping dari struct Request ke struct Core
func BookRequestToCore(dataReq BookRequest) BookCore {
	return BookCore{
		Title:       dataReq.Title,
		Publisher:   dataReq.Publisher,
		Author:      dataReq.Author,
		PublishYear: dataReq.PublishYear,
		UserID:      dataReq.UserID,
	}
}
