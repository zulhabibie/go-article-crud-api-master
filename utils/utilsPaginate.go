package utils

import (
	"go-crud-article/structs"
	"net/http"
	"strconv"
)

func GeneratePagination(r *http.Request) structs.Pagination {
	limit := 10
	page := 1
	sort := "userid asc"
	query := r.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}
	return structs.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
