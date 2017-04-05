package lint

import (
	"regexp"
)

var (
	reIDs  = regexp.MustCompile("^(.*)Ids$")
	reURLs = regexp.MustCompile("^(.*)Urls$")
)

// Name returns a different name if it should be different.
func Name(name string) (should string) {
	should = lintName(name)
	// Rename SomethingIds to SomethingIDs.
	should = reIDs.ReplaceAllString(should, "${1}IDs")
	should = reURLs.ReplaceAllString(should, "${1}URLs")
	if should == "Idref" {
		return "IDRef"
	}
	return
}
