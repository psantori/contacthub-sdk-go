package client

import "strconv"

// PageInfo contains the pagination info from list endpoints
type PageInfo struct {
	Size                    int `json:"size"`
	TotalElements           int `json:"totalElements"`
	TotalUnfilteredElements int `json:"totalUnfilteredElements"`
	TotalPages              int `json:"totalPages"`
	Page                    int `json:"number"`
}

// HasNextPage checks if exists a page after the current one
func (p *PageInfo) HasNextPage() bool {
	return p.Page+1 < p.TotalPages
}

// ListParams contains the params for list endpoints
// The Page param overrides the "page" value in the QueryParams
type ListParams struct {
	Page        int
	QueryParams QueryParams
}

func (p *ListParams) preparePagination() {
	if p.QueryParams == nil {
		p.QueryParams = QueryParams{}
	}
	p.QueryParams["page"] = strconv.Itoa(p.Page)
}
