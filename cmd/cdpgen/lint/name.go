package lint

import (
	"fmt"
	"regexp"

	lint "github.com/mafredri/go-lint"
)

var (
	reIDs   = regexp.MustCompile("^(.*)Ids($|[A-Z].*$)")
	reUUIDs = regexp.MustCompile("^(.*)Uuids($|[A-Z].*$)")
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
	lint.SetInitialism("JS", true)
	lint.SetInitialism("JXL", true)
	lint.SetInitialism("BR", true)  // Brotli compression.
	lint.SetInitialism("HWB", true) // Color format.
	lint.SetInitialism("IME", true)
	lint.SetInitialism("CM", true)   // FedCm => FedCM.
	lint.SetInitialism("OS", true)   // Operating system.
	lint.SetInitialism("COEP", true) // Cross-Origin-Embedder-Policy.
	lint.SetInitialism("COOP", true) // Cross-Origin-Opener-Policy.
	lint.SetInitialism("CORP", true) // Cross-Origin-Resource-Policy.
	lint.SetInitialism("CVC", true)  // Card verification code.
	lint.SetInitialism("IDP", true)  // Identity Provider.
	lint.SetInitialism("CSP", true)  // Content Security Policy.
	lint.SetInitialism("DIP", true)  // Document isolation policy.
	lint.SetInitialism("RSSI", true)
	lint.SetInitialism("CH", true)
	lint.SetInitialism("DPR", true)
	lint.SetInitialism("UA", true)
	lint.SetInitialism("ECT", true)
	lint.SetInitialism("RTT", true)
	lint.SetInitialism("HID", true)
	lint.SetInitialism("OTP", true)
	lint.SetInitialism("XHR", true)
}

// Name returns a different name if it should be different.
func Name(name string) (should string) {
	should = lint.Name(name)

	for _, replace := range []struct {
		re *regexp.Regexp
		to string
	}{
		{re: reIDs, to: "IDs"},
		{re: reUUIDs, to: "UUIDs"},
		{re: reURLs, to: "URLs"},
		{re: reIDRef, to: "IDRef"},
	} {
		should = replace.re.ReplaceAllString(should, fmt.Sprintf("${1}%s${2}", replace.to))
	}

	return should
}
