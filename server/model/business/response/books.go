package response

import "time"

type Books struct {
	BookId          int64     `json:"bookId"`
	BookName        string    `json:"bookName"`
	Author          string    `json:"author"`
	Price           string    `json:"price"`
	Publish         string    `json:"publish"`
	PublicationDate time.Time `json:"publicationDate"`
	Isbn            string    `json:"isbn"`
	CoverImage      string    `json:"coverImage"`
	Summary         string    `json:"summary"`
	Category        string    `json:"category"`
	Remark          string    `json:"remark"`
	Status          int       `json:"status"`
}
