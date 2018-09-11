package lint

import (
	"regexp"

	lint "github.com/mafredri/go-lint"
)

var (
	reIDs  = regexp.MustCompile("^(.*)Ids$")
	reURLs = regexp.MustCompile("^(.*)Urls$")
)

func init() {
	lint.SetInitialism("DOM", true)
	lint.SetInitialism("GPU", true)
	lint.SetInitialism("SSL", true)
	lint.SetInitialism("MAC", true)
}

// Name returns a different name if it should be different.
func Name(name string) (should string) {
	should = lint.Name(name)
	// Rename SomethingIds to SomethingIDs.
	should = reIDs.ReplaceAllString(should, "${1}IDs")
	should = reURLs.ReplaceAllString(should, "${1}URLs")
	if should == "Idref" {
		return "IDRef"
	}
	return
}
