package dto

type SortDto struct {
	ColId string `json:"colId"`
	Sort  string `json:"sort"`
}

type Filter struct {
	Type       string `json:"type"`
	From       string `json:"from"`
	To         string `json:"to"`
	FilterType string `json:"filterType"`
}

type DynamicFilter struct {
	Sort   *[]SortDto        `json:"sort"`
	Filter map[string]Filter `json:"filter"`
}

type PageList[T any] struct {
	PageNumber  int   `json:"page_number"`
	TotalRows   int64 `json:"total_rows"`
	TotalPages  int   `json:"total_pages"`
	HasPervious bool  `json:"has_pervious"`
	HasNextPage bool  `json:"has_next_page"`
	Items       *[]T  `json:"items"`
}

type PaginationInput struct {
	PageSize   int `json:"page_size"`
	PageNumber int `json:"page_number"`
}

type PaginationInputWithFilter struct {
	PaginationInput
	DynamicFilter
}

func (p *PaginationInputWithFilter) GetOffSet() int {
	return p.GetPageNumber() - 1*p.GetPageSize()
}

func (p *PaginationInputWithFilter) GetPageSize() int {
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	return p.PageSize
}

func (p *PaginationInputWithFilter) GetPageNumber() int {
	if p.PageNumber == 0 {
		p.PageNumber = 1
	}
	return p.PageNumber
}
