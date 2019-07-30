package ustr

import (
	"sort"
)

type sorter struct {
	s    []string
	less func(string, string) bool
}

func (me *sorter) Swap(i int, j int)      { me.s[i], me.s[j] = me.s[j], me.s[i] }
func (me *sorter) Len() int               { return len(me.s) }
func (me *sorter) Less(i int, j int) bool { return me.less(me.s[i], me.s[j]) }

func Sortable(s []string, less func(string, string) bool) sort.Interface {
	return &sorter{s: s, less: less}
}
