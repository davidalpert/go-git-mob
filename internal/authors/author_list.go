package authors

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"sort"
)

// AuthorList is a helper type to support string format when printing a list of authors
type AuthorList struct {
	Members []*Author `json:"members"`
}

// String implements Stringer to format our list output as a string
func (m AuthorList) String() string {
	var s string
	for _, a := range m.Members {
		s += fmt.Sprintf("%s\n", a.String())
	}
	return s
}

func (m AuthorList) SortBy(less func(left, right *Author) bool) {
	sort.Slice(m.Members, func(i, j int) bool {
		left := m.Members[i]
		right := m.Members[j]
		return less(left, right)
	})
}

func (m AuthorList) WriteToTable(table *tablewriter.Table) {
	table.SetHeader([]string{"Name", "Email"})
	for _, a := range m.Members {
		table.Append([]string{a.Name, a.Email})
	}
}
