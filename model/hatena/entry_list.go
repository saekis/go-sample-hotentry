package hatena

import "sort"

type EntryList []Entry

func (el *EntryList) SortByBookmarkUser() {
	sort.SliceStable((*el), func(i, j int) bool { return (*el)[i].Bookmarkcount < (*el)[j].Bookmarkcount })
}
