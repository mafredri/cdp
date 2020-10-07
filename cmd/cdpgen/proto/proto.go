// Package proto is used to parse the CDP protocol definitions (JSON).
package proto

import (
	"log"
	"strings"

	"github.com/mafredri/cdp/cmd/cdpgen/lint"

	"github.com/client9/misspell"
)

var misspellReplacer = misspell.New()

func init() {
	misspellReplacer.AddRuleList(misspell.DictAmerican)
	misspellReplacer.Compile()
}

// Protocol represents the JSON protocol structure.
type Protocol struct {
	Version Version  `json:"version,omitempty"`
	Domains []Domain `json:"domains,omitempty"`
}

// Version contains protocol version information.
type Version struct {
	Major string `json:"major,omitempty"`
	Minor string `json:"minor,omitempty"`
}

// Domain represents a domain, e.g. Page, Network, etc.
type Domain struct {
	Domain       string    `json:"domain,omitempty"`
	Experimental bool      `json:"experimental,omitempty"`
	Description  string    `json:"description,omitempty"`
	Dependencies []string  `json:"dependencies,omitempty"`
	Types        []AnyType `json:"types,omitempty"`
	Commands     []Command `json:"commands,omitempty"`
	Events       []Event   `json:"events,omitempty"`
	Deprecated   bool      `json:"deprecated,omitempty"`
}

// Name returns the domain name.
func (d Domain) Name() string {
	return d.Domain
}

// Type returns the domain type.
func (d Domain) Type() string {
	return d.Domain
}

// Desc returns the domain decription.
func (d Domain) Desc(indent, startOffset int) string {
	return cleanDescription(d.Description, indent, startOffset)
}

// Command represents a command belonging to a domain, e.g. Network.setCookie.
type Command struct {
	NameName     string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	Parameters   []AnyType `json:"parameters,omitempty"`
	Returns      []AnyType `json:"returns,omitempty"`
	Optional     bool      `json:"optional,omitempty"`
	Experimental bool      `json:"experimental,omitempty"`
	Redirect     string    `json:"redirect,omitempty"`
	Handlers     []string  `json:"handlers,omitempty"`
	Deprecated   bool      `json:"deprecated,omitempty"`
}

// Name rturns the linted command name.
func (c Command) Name() string {
	return lint.Name(strings.Title(c.NameName))
}

// Desc returns a cleaned description.
func (c Command) Desc(lineEndComment bool, indent, startOffset int) string {
	if lineEndComment {
		return cleanDescription(c.Description, indent, startOffset)
	}
	return lowerFirst(cleanDescription(c.Description, indent, startOffset))
}

// CmdName returns the full name of a command.
func (c Command) CmdName(d Domain, export bool) string {
	name := d.Name()
	if !export {
		name = strings.ToLower(d.Name())
	}
	return name + c.Name() + ""
}

// ArgsName returns the name of command arguments.
func (c Command) ArgsName(d Domain) string {
	return c.Name() + "Args"
}

// ReplyName returns the name of the command reply.
func (c Command) ReplyName(d Domain) string {
	return c.Name() + "Reply"
}

// ArgsSignature returns the signature (for use as function parameters).
func (c Command) ArgsSignature(pkg string, d Domain) string {
	var args []string
	for _, arg := range filter(optional(false), c.Parameters...) {
		name := arg.Name(d)
		if name == "range" || name == "type" {
			name = name[0 : len(name)-1]
		}
		name += " "
		args = append(args, name+arg.GoType(pkg, d))
	}
	return strings.Join(args, ", ")
}

// ArgsInit returns the code for initializing arguments.
func (c Command) ArgsInit(d Domain) string {
	var args []string
	for _, arg := range filter(optional(false), c.Parameters...) {
		name := arg.Name(d)
		if name == "range" || name == "type" {
			name = name[0 : len(name)-1]
		}
		args = append(args, arg.ExportedName(d)+": "+name+",")
	}
	return strings.Join(args, "\n")
}

// ArgsAssign returns the argument assignment for args.
func (c Command) ArgsAssign(receiver string, d Domain) string {
	var args []string
	for _, arg := range filter(optional(false), c.Parameters...) {
		name := arg.Name(d)
		if name == "range" || name == "type" {
			name = name[0 : len(name)-1]
		}
		args = append(args, receiver+"."+arg.ExportedName(d)+" = "+name)
	}
	return strings.Join(args, "\n")
}

// ReplySignature returns the reply signature. Not used.
func (c Command) ReplySignature(d Domain) string {
	var args []string
	for _, arg := range c.Returns {
		name := arg.Name(d)
		if name == "range" || name == "type" {
			name = name[0 : len(name)-1]
		}

		typ := arg.GoType("cdp", d)
		if arg.Optional && !strings.HasPrefix(typ, "[]") {
			typ = "*" + typ
		}
		args = append(args, name+" "+typ)
	}
	return strings.Join(args, ", ")
}

// ReplyAssign assigns the parameters of the reply. Not used.
func (c Command) ReplyAssign(receiver string, d Domain) string {
	var args []string
	for _, arg := range c.Returns {
		name := arg.Name(d)
		if name == "range" || name == "type" {
			name = name[0 : len(name)-1]
		}
		args = append(args, name+" = "+receiver+"."+arg.ExportedName(d))
	}
	return strings.Join(args, "\n")
}

// Event represents an subscribeable event.
type Event struct {
	NameName     string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	Parameters   []AnyType `json:"parameters,omitempty"`
	Experimental bool      `json:"experimental,omitempty"`
	Deprecated   bool      `json:"deprecated,omitempty"`
}

// Name returns the name of the event.
func (e Event) Name() string {
	return lint.Name(strings.Title(e.NameName))
}

// Desc returns the cleaned description.
func (e Event) Desc(lineEndComment bool, indent, startOffset int) string {
	if lineEndComment {
		return cleanDescription(e.Description, indent, startOffset)
	}
	return lowerFirst(cleanDescription(e.Description, indent, startOffset))
}

// EventName returns the name of the event as a go type.
func (e Event) EventName(d Domain) string {
	return nameInDomain(d, e.Name(), "")
}

// ReplyName returns the name of the event reply struct.
func (e Event) ReplyName(d Domain) string {
	return e.EventName(d) + "Reply"
}

// Enum represents an enumerable value.
type Enum string

// Name returns the Go-ified name for the enum.
func (e Enum) Name() string {
	switch e {
	case "-Infinity", "-0":
		return strings.Replace(string(e), "-", "Negative", 1)
	}
	s := strings.Replace(string(e), "-", " ", -1)
	s = strings.Title(s)
	return lint.Name(strings.Replace(s, " ", "", -1))
}

// AnyType is a catch-all struct for properties, parameters, etc.
type AnyType struct {
	IDName       string    `json:"id,omitempty"`
	NameName     string    `json:"name,omitempty"`
	Type         string    `json:"type,omitempty"`
	Description  string    `json:"description,omitempty"`
	Ref          string    `json:"$ref,omitempty"`
	Properties   []AnyType `json:"properties,omitempty"`
	Enum         []Enum    `json:"enum,omitempty"`
	Items        *AnyType  `json:"items,omitempty"`
	MinItems     int       `json:"minItems,omitempty"`
	MaxItems     int       `json:"maxItems,omitempty"`
	Optional     bool      `json:"optional,omitempty"`
	Deprecated   bool      `json:"deprecated,omitempty"`
	Experimental bool      `json:"experimental,omitempty"`
}

// Desc returns the cleaned description.
func (at AnyType) Desc(indent, startOffset int) string {
	return cleanDescription(at.Description, indent, startOffset)
}

// ExportedName returns an exported name.
func (at AnyType) ExportedName(d Domain) string {
	if at.IDName != "" {
		return at.Name(d)
	}
	return lint.Name(strings.Title(at.NameName))
}

// Name returns a Go-ified name for the AnyType.
func (at AnyType) Name(d Domain) string {
	if at.IDName != "" {
		return nameInDomain(d, at.IDName, "")
	}

	return lint.Name(at.NameName)
}

// Recvr returns the receiver for the type.
func (at AnyType) Recvr(d Domain) string {
	return strings.ToLower(at.Name(d)[0:1])
}

func nameInDomain(d Domain, name, _ string) string {
	name = lint.Name(strings.Title(name))
	if name != d.Name() && strings.Index(name, d.Name()) == 0 {
		name = strings.Replace(name, d.Name(), "", 1)
	}
	return name
}

func (at AnyType) getType(pkg string) string {
	if at.Ref != "" {
		return "reference"
	}
	if (at.IDName == "Timestamp" || at.IDName == "NetworkTimeSinceEpoch" || (pkg != "network" && at.IDName == "TimeSinceEpoch") || at.IDName == "MonotonicTime") && at.Type == "number" {
		return "time"
	}
	if at.IsEnum() {
		return "enum"
	}
	if at.Type == "string" && strings.HasPrefix(at.Description, "Base64-encoded") {
		return "base64"
	}
	if at.Type == "object" && len(at.Properties) == 0 {
		if at.IDName != "" {
			return "raw"
		}
		return "any"
	}
	return at.Type
}

// GoType returns the Go representation for a protocol type.
func (at AnyType) GoType(pkg string, d Domain) string {
	// Special case for circular types.
	if pkg == "page" && at.Name(d) == "FrameID" {
		return "= internal.PageFrameID"
	}
	if pkg == "network" && at.Name(d) == "TimeSinceEpoch" {
		return "= internal.NetworkTimeSinceEpoch"
	}
	if pkg == "browser" && at.Name(d) == "ContextID" {
		return "= internal.BrowserContextID"
	}

	switch at.getType(pkg) {
	case "any":
		return "json.RawMessage"
	case "boolean":
		return "bool"
	case "string":
		return "string"
	case "number":
		return "float64"
	case "integer":
		return "int"
	case "object":
		return "struct"
	case "array":
		return "[]" + at.Items.GoType(pkg, d)
	case "reference":
		return at.refType(pkg, d)
	case "time":
		return "time.Time"
	case "enum":
		// Enums are handled specially.
		return "enum"
	case "base64":
		// By using a []byte here, Base64-encoded images are
		// automatically decoded by json.Unmarshal.
		return "[]byte"
	case "raw":
		// Special []byte, without bas64 decoding.
		return "RawMessage"
	default:
		log.Panicf("unhandled type: %#v", at)
	}

	panic("unreachable")
}

func (at AnyType) refType(pkg string, d Domain) string {
	var prefix string
	if strings.ContainsRune(at.Ref, '.') {
		s := strings.Split(at.Ref, ".")
		prefix = strings.ToLower(s[0]) + "."
		s[0] = lint.Name(strings.Title(s[0]))
		s[1] = lint.Name(strings.Title(s[1]))

		// Remove stutter, e.g. SecuritySecurityState.
		if strings.Index(s[1], s[0]) == 0 || s[1] == s[0] {
			s[1] = strings.Replace(s[1], s[0], "", 1)
		}

		// Fix types that reference their own domain.
		if s[0] == d.Name() {
			prefix = ""
			at.Ref = s[1]
		}
		name := prefix + strings.Title(lint.Name(s[1]))

		// Special cases for circular types.
		switch {
		case (pkg == "network" || pkg == "dom") && name == "page.FrameID":
			return "internal.PageFrameID"
		case pkg == "target" && name == "browser.ContextID":
			return "internal.BrowserContextID"
		case pkg == "security" && name == "network.TimeSinceEpoch":
			return "internal.NetworkTimeSinceEpoch"
		}
		return name
	}
	return prefix + nameInDomain(d, at.Ref, "")
}

// IsEnum returns true if type is an enum.
func (at AnyType) IsEnum() bool {
	return at.IDName != "" && len(at.Enum) > 0
}

// IsLocalEnum returns true if type is enumerated without an exported type.
func (at AnyType) IsLocalEnum() bool {
	return at.IDName == "" && len(at.Enum) > 0
}

func lowerFirst(d string) string {
	desc := strings.Split(d, " ")
	if desc[0] != strings.ToUpper(desc[0]) {
		desc[0] = strings.ToLower(desc[0])
	}
	return strings.Join(desc, " ")
}

// Account 3 for comment prefix (// ).
const maxCommentLineLen = 80 - 3

func cleanDescription(d string, indent, startOffset int) string {
	replace := []struct {
		old string
		new string
	}{
		{"<code>", ""}, {"</code>", ""},
		// <p> is only used by DOM description.
		{"<p>", "\n\n"}, {"</p>", ""},
		{"&lt;", "<"}, {"&gt;", ">"},
		// Fix typo...
		{"&gt ", "> "},
	}

	for _, r := range replace {
		d = strings.Replace(d, r.old, r.new, -1)
	}

	d, _ = misspellReplacer.Replace(d)

	p := strings.Split(d, "\n\n")
	for i, s := range p {
		ss := strings.Fields(s)

		var split []string
		n := startOffset
		startOffset = 0 // Zero after first use.
		s = ""
		for _, sss := range ss {
			n += len(sss) + 1
			if n < maxCommentLineLen-indent {
				split = append(split, sss)
				continue
			}
			n = len(sss)
			s += strings.Join(split, " ") + "\n"
			split = nil
			split = append(split, sss)
		}

		if len(split) > 0 {
			s += strings.Join(split, " ") + "\n"
		}

		s = strings.TrimSuffix(s, "\n")

		p[i] = strings.Replace(s, "\n", "\n// ", -1)
	}

	return strings.Join(p, "\n//\n// ")
}

type filterFunc func(at AnyType) bool

func filter(f filterFunc, at ...AnyType) []AnyType {
	var ret []AnyType
	for _, a := range at {
		if f(a) {
			ret = append(ret, a)
		}
	}
	return ret
}

func optional(o bool) filterFunc {
	return func(at AnyType) bool {
		return at.Optional == o
	}
}
