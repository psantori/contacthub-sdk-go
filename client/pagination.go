/**
 * This file is part of contacthub-sdk-go.
 *
 * contacthub-sdk-go is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 2 of the License, or
 * (at your option) any later version.
 *
 * contacthub-sdk-go is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with contacthub-sdk-go. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (C) 2017 Arduino AG
 *
 * @author Luca Osti
 *
 */

package client

import "strconv"

// DefaultPageSize is the default page size for pagination. Max is 50.
const DefaultPageSize int = 20

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
	PageSize    int
	QueryParams QueryParams
}

func (p *ListParams) preparePagination() {
	if p.QueryParams == nil {
		p.QueryParams = QueryParams{}
	}
	if p.PageSize == 0 {
		p.PageSize = DefaultPageSize
	}
	p.QueryParams["page"] = strconv.Itoa(p.Page)
	p.QueryParams["size"] = strconv.Itoa(p.PageSize)
}
