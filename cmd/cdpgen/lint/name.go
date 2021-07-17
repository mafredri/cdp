package lint

import (
	"fmt"
	"regexp"

	lint "github.com/mafredri/go-lint"
)

var (
	reIDs   = regexp.MustCompile("^(.*)Ids($|[A-Z].*$)")
	reURLs  = regexp.MustCompile("^(.*)Urls($|[A-Z].*$)")
	reIDRef = regexp.MustCompile("^(.*)Idref($|[A-Z].*$)")
)

func init() {
	lint.SetInitialism("DOM", true)
	lint.SetInitialism("GPU", true)
	lint.SetInitialism("SSL", true)
	lint.SetInitialism("MAC", true)
	lint.SetInitialism("U2F", true)
	lint.SetInitialism("CTAP", true)
	lint.SetInitialism("USB", true)
	lint.SetInitialism("NFC", true)
	lint.SetInitialism("BLE", true)
	lint.SetInitialism("RP", true) // RPID => Relaying Party (ID).
	lint.SetInitialism("JPEG", true)
	lint.SetInitialism("WEBP", true)
	lint.SetInitialism("AVIF", true)
	lint.SetInitialism("RGB", true)
	lint.SetInitialism("HSL", true)
	lint.SetInitialism("CORS", true)
	lint.SetInitialism("WASM", true)
	lint.SetInitialism("JXL", true)
	lint.SetInitialism("BR", true) // Brotli compression.
}

// Name returns a different name if it should be different.
func Name(name string) (should string) {
	should = lint.Name(name)

	for _, replace := range []struct {
		re *regexp.Regexp
		to string
	}{
		{re: reIDs, to: "IDs"},
		{re: reURLs, to: "URLs"},
		{re: reIDRef, to: "IDRef"},
	} {
		should = replace.re.ReplaceAllString(should, fmt.Sprintf("${1}%s${2}", replace.to))
	}

	return should
}
