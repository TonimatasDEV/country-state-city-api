package util

import (
	"net/http"
	"strconv"
	"strings"
)

type Filter struct {
	Value string
}

func GetFilters(r *http.Request) map[string]Filter {
	filters := strings.Split(r.RequestURI, "?")
	result := make(map[string]Filter)

	if len(filters) != 2 {
		return result
	}

	for _, filter := range strings.Split(filters[1], "&") {
		split := strings.Split(filter, "=")

		if len(split) != 2 {
			continue
		}

		result[split[0]] = Filter{split[1]}
	}

	return result
}

func (filter *Filter) IsEmpty() bool {
	return filter.Value == ""
}

func (filter *Filter) Int() int {
	value, err := strconv.Atoi(filter.Value)

	if err != nil {
		return 0
	}

	return value
}
