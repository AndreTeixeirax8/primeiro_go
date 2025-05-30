package pagination

import (
	"net/url"
	"strconv"
	"strings"
)

type Filter struct {
	Field  string `json:"field"`  // Field to filter by
	Action string `json:"action"` // Action to perform (e.g., "equals", "contains", "greater_than", etc.)
	Value  string `json:"value"`  // Value to filter by
}

type PaginationQuery struct {
	Limit   int      `json:"limit"`   // Number of items per page
	Page    int      `json:"page"`    // Current page number
	Sort    string   `json:"sort"`    // Sort field
	Search  string   `json:"search"`  // Search term
	Filters []Filter `json:"filters"` // List of filters to apply
}

type PaginationMeta struct {
	TotalRows   int `json:"total_rows"`   // Total number of rows
	FromRow     int `json:"from_row"`     // Starting row number
	ToRow       int `json:"to_row"`       // Ending row number
	TotalPages  int `json:"total_pages"`  // Total number of pages
	PerPage     int `json:"per_page"`     // Number of items per page
	CurrentPage int `json:"current_page"` // Current page number
}

type PaginationResponse[E interface{}] struct {
	Rows []E            `json:"rows"`
	Meta PaginationMeta `json:"meta"` // Metadata about the pagination
}

func GeneratePaginationRequest(query url.Values) *PaginationQuery {

	limit := 10 // Default limit
	page := 1   // Default page
	sort := ""
	search := ""

	var filters []Filter

	for key, values := range query {
		queryValue := values[0]
		switch key {
		case "limit":
			if len(values) > 0 {
				limit, _ = strconv.Atoi(queryValue)
			}
		case "page":
			if len(values) > 0 {
				page, _ = strconv.Atoi(queryValue)
			}
		case "sort":
			sort = queryValue
		case "search":
			search = queryValue
		}
		if strings.Contains(key, ".") {
			filterKeys := strings.Split(key, ".")
			filter := Filter{
				Field:  filterKeys[0],
				Action: filterKeys[1],
				Value:  queryValue,
			}
			filters = append(filters, filter)
		}
	}

	return &PaginationQuery{
		Limit:   limit,
		Page:    page,
		Sort:    sort,
		Search:  search,
		Filters: filters,
	}
}
