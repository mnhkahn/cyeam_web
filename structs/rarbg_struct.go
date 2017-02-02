package structs

type RarbgView struct {
	Items []*RarbgResult
	Pages []*Nav
	Pre   *Nav
	Next  *Nav
}

type Nav struct {
	Page     int
	Url      string
	Selected bool
}

type RarbgResult struct {
	ID            string
	Name          string
	OriginalTitle string
	Figure        string
	Format        string
	Rate          float64
	Date          string
	Size          string
}

type DoubanResult struct {
	Count    int64 `json:"count"`
	Start    int64 `json:"start"`
	Subjects []struct {
		Alt string `json:"alt"`
		// Casts []struct {
		// 	Alt     string `json:"alt"`
		// 	Avatars struct {
		// 		Large  string `json:"large"`
		// 		Medium string `json:"medium"`
		// 		Small  string `json:"small"`
		// 	} `json:"avatars"`
		// 	ID   string `json:"id"`
		// 	Name string `json:"name"`
		// } `json:"casts"`
		// CollectCount int64 `json:"collect_count"`
		// Directors    []struct {
		// 	Alt     string `json:"alt"`
		// 	Avatars struct {
		// 		Large  string `json:"large"`
		// 		Medium string `json:"medium"`
		// 		Small  string `json:"small"`
		// 	} `json:"avatars"`
		// 	ID   string `json:"id"`
		// 	Name string `json:"name"`
		// } `json:"directors"`
		// Genres []string `json:"genres"`
		ID     string `json:"id"`
		Images struct {
			Large  string `json:"large"`
			Medium string `json:"medium"`
			Small  string `json:"small"`
		} `json:"images"`
		// OriginalTitle string `json:"original_title"`
		Rating struct {
			Average float64 `json:"average"`
			Max     int64   `json:"max"`
			Min     int64   `json:"min"`
			Stars   string  `json:"stars"`
		} `json:"rating"`
		// Subtype string `json:"subtype"`
		Title string `json:"title"`
		// Year  string `json:"year"`
	} `json:"subjects"`
	// Title string `json:"title"`
	Total int64 `json:"total"`
}
