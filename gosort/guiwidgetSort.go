package gosort

import (
	"fmt"
	"sort"
)

type compareFunc func(x, y *Track, key string) bool

type sortByClicked struct {
	t           []*Track
	keySort     *[]string
	less, equal compareFunc
}

func (s sortByClicked) Len() int      { return len(s.t) }
func (s sortByClicked) Swap(i, j int) { s.t[i], s.t[j] = s.t[j], s.t[i] }
func (s sortByClicked) Less(i, j int) bool {
	for _, key := range *s.keySort {
		if !s.equal(s.t[i], s.t[j], key) {
			return s.less(s.t[i], s.t[j], key)
		}
	}
	return false
}

func (s sortByClicked) updateKeysort(key string) {
	var i = -1
	var keyfound string
	for index, sortkey := range *s.keySort {
		if sortkey == key {
			i = index
			keyfound = sortkey
		}
	}
	if i != -1 {
		for index := i; index >= 1; index-- {
			(*s.keySort)[index] = (*s.keySort)[index-1]
		}
		(*s.keySort)[0] = keyfound
	}
}

func isequal(x, y *Track, key string) bool {
	switch key {
	case "Title":
		return x.Title == y.Title
	case "Artist":
		return x.Artist == y.Artist
	case "Album":
		return x.Album == y.Album
	case "Year":
		return x.Year == y.Year
	case "Length":
		return x.Length == y.Length
	}
	return false
}

func less(x, y *Track, key string) bool {
	switch key {
	case "Title":
		return x.Title < y.Title
	case "Artist":
		return x.Artist < y.Artist
	case "Album":
		return x.Album < y.Album
	case "Year":
		return x.Year < y.Year
	case "Length":
		return x.Length < y.Length
	}
	return false
}

// TestDynamicSort is the test function for dynamic column sort
func TestDynamicSort() {
	testString := []string{"Title", "Artist", "Album", "Year", "Length"}
	sortmethod := sortByClicked{t: tracks, keySort: &testString, less: less, equal: isequal}
	sort.Sort(sortmethod)
	printTracks(tracks)
	sortmethod.updateKeysort("Length")
	fmt.Println("click Length-----------", testString)
	sort.Sort(sortmethod)
	printTracks(tracks)
	sortmethod.updateKeysort("Year")
	fmt.Println("click Year-------------", testString)
	sort.Sort(sortmethod)
	printTracks(tracks)
}
