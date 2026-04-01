package markdown

import (
	"sort"
	"strings"
)

func SortSectionsByTitle(sections []Section) {
	sort.SliceStable(sections, func(i, j int) bool {
		return strings.ToLower(sections[i].Title) < strings.ToLower(sections[j].Title)
	})
}
