package sportmonks

type (
	// Meta struct.
	Meta struct {
		Pagination *Pagination `json:"pagination,omitempty"`
		Plan       *Plan       `json:"plan,omitempty"`
		Sports     *Sports     `json:"sports,omitempty"`
	}

	// Pagination struct.
	Pagination struct {
		Total       int `json:"total"`
		Count       int `json:"count"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
	}

	// Plan struct.
	Plan struct {
		Name         string `json:"name"`
		Price        string `json:"price"`
		RequestLimit string `json:"request_limit"`
	}

	// Sports struct.
	Sports []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Current bool   `json:"current"`
	}

	// DateTime struct.
	DateTime struct {
		Date         string `json:"date"`
		TimezoneType int    `json:"timezone_type"`
		Timezone     string `json:"timezone"`
	}
)
