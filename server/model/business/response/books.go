package response

type Books struct {
	BookId          string `json:"bookId"`
	BookName        string `json:"bookName"`
	Author          string `json:"author"`
	Price           string `json:"price"`
	Publisher       string `json:"publisher"`
	PublicationDate string `json:"publicationDate"`
	Isbn            string `json:"isbn"`
	CoverImage      string `json:"coverImage"`
	Summary         string `json:"summary"`
	Category        int8   `json:"category"`
	Remark          string `json:"remark"`
	Status          int8   `json:"status"`
}
