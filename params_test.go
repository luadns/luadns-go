package luadns_test

import (
	"testing"

	"github.com/luadns/luadns-go"
	"github.com/stretchr/testify/assert"
)

func TestListOptionsQueryString(t *testing.T) {
	tests := []struct {
		opts *luadns.ListParams
		qs   string
	}{
		{&luadns.ListParams{}, ""},
		{&luadns.ListParams{Query: "foo bar"}, "query=foo+bar"},
		{&luadns.ListParams{SortBy: "name"}, "sort_by=name"},
		{&luadns.ListParams{SortOrder: "asc"}, "sort_order=asc"},
		{&luadns.ListParams{SortOrder: "desc"}, "sort_order=desc"},
		{&luadns.ListParams{Limit: 1}, "limit=1"},
		{&luadns.ListParams{Page: 1}, "page=1"},
		{&luadns.ListParams{Limit: 1, Page: 1}, "limit=1&page=1"},
	}

	for _, test := range tests {
		assert.Equal(t, test.qs, test.opts.QueryString())
	}
}
