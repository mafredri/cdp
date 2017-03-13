package lint

import (
	"regexp"
)

var reIDs = regexp.MustCompile("^(.*)Ids$")

// Name returns a different name if it should be different.
func Name(name string) (should string) {
	should = lintName(name)
	// Rename SomethingIds to SomethingIDs.
	should = reIDs.ReplaceAllString(should, "${1}IDs")
	if should == "Idref" {
		return "IDRef"
	}
	return
}
