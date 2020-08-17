package entity

type Pagination struct {
	CurrentPage     int64  `json:"current_page"`
	CurrentElements int64  `json:"current_elements"`
	TotalPages      int64  `json:"total_pages"`
	TotalElements   int64  `json:"total_elements"`
	SortBy          string `json:"sort_by"`
}

type Meta struct {
	Path       string `json:"path"`
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	Error      *error `json:"error,omitempty"`
	Timestamp  string `json:"timestamp"`
}
