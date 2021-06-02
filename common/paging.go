package common

import "strings"

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"total"`

	FakeCursor string `json:"cursor" form:"cursor"`
	NextCursor string `json:"next_cursor"`
}

func (paging *Paging) Fulfill() {
	if paging.Page <= 0 {
		paging.Page = 1
	}

	if paging.Limit <= 0 {
		paging.Limit = 50
	}

	paging.FakeCursor = strings.TrimSpace(paging.FakeCursor)
}
