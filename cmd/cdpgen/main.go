// The cdpgen command generates the package cdp from the provided protocol definitions.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/mafredri/cdp/cmd/cdpgen/proto"
)

// Global constants.
const (
	OptionalPropPrefix = ""
	realEnum           = true
)

var (
	nonPtrMap = make(map[string]bool)
)

func main() {
	var (
		destPkg          string
		browserProtoJSON string
		jsProtoFileJSON  string
	)
	flag.StringVar(&destPkg, "dest-pkg", "", "Destination for generated cdp package (inside $GOPATH)")
	flag.StringVar(&browserProtoJSON, "browser-proto", "./protodef/browser_protocol.json", "Path to browser protocol")
	flag.StringVar(&jsProtoFileJSON, "js-proto", "./protodef/js_protocol.json", "Path to JS protocol")
	flag.Parse()

	if destPkg == "" {
		fmt.Fprintln(os.Stderr, "error: dest-pkg must be set")
		os.Exit(1)
	}

	var protocol, jsProtocol proto.Protocol
	protocolData, err := ioutil.ReadFile(browserProtoJSON)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(protocolData, &protocol)
	if err != nil {
		panic(err)
	}
	jsProtocolData, err := ioutil.ReadFile(jsProtoFileJSON)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsProtocolData, &jsProtocol)
	if err != nil {
		panic(err)
	}

	protocol.Domains = append(protocol.Domains, jsProtocol.Domains...)
	sort.Slice(protocol.Domains, func(i, j int) bool {
		return protocol.Domains[i].Domain < protocol.Domains[j].Domain
	})

	var cdpgen, typegen, cmdgen, eventgen, domgen Generator

	cdpgen.dir = destPkg
	cdpgen.pkg = "cdp"
	os.Mkdir(cdpgen.path(), 0755)

	typegen.pkg = "cdptype"
	typegen.dir = path.Join(cdpgen.dir, typegen.pkg)
	os.Mkdir(typegen.path(), 0755)

	cmdgen.pkg = "cdpcmd"
	cmdgen.dir = path.Join(cdpgen.dir, cmdgen.pkg)
	os.Mkdir(cmdgen.path(), 0755)

	eventgen.pkg = "cdpevent"
	eventgen.dir = path.Join(cdpgen.dir, eventgen.pkg)
	os.Mkdir(eventgen.path(), 0755)

	domgen.pkg = "cdpdom"
	domgen.dir = path.Join(cdpgen.dir, domgen.pkg)
	os.Mkdir(domgen.path(), 0755)

	cdpgen.imports = []string{
		"github.com/mafredri/cdp/rpcc",
		typegen.dir, cmdgen.dir, eventgen.dir,
	}
	cmdgen.imports = []string{typegen.dir}
	eventgen.imports = []string{typegen.dir}
	domgen.imports = []string{typegen.dir}

	// Define the cdp Client.
	cdpgen.PackageHeader()
	cdpgen.CdpClient(protocol.Domains)
	cdpgen.writeFile("cdp_client.go")

	// Define all CDP command methods.
	cmdgen.PackageHeader()
	cmdgen.CmdType(protocol.Domains)
	cmdgen.writeFile("cmd.go")

	// Define all CDP event methods.
	eventgen.PackageHeader()
	eventgen.EventType(protocol.Domains)
	eventgen.writeFile("event.go")

	for _, d := range protocol.Domains {
		for _, t := range d.Types {
			nam := t.Name(d)
			if isNonPointer(typegen.pkg, d, t) {
				nonPtrMap[nam] = true
				nonPtrMap[typegen.pkg+"."+nam] = true
			}
		}
	}
	nonPtrMap["Timestamp"] = true
	nonPtrMap[typegen.pkg+"."+"Timestamp"] = true

	for _, d := range protocol.Domains {
		cdpgen.PackageHeader()
		cdpgen.DomainInterface(d)

		domgen.PackageHeader()
		domgen.DomainDefinition(d)

		if len(d.Types) > 0 {
			typegen.PackageHeader()
			for _, t := range d.Types {
				typegen.DomainType(d, t)
			}
		}

		if len(d.Commands) > 0 {
			cmdgen.PackageHeader()
			for _, c := range d.Commands {
				cmdgen.DomainCmd(d, c)
			}
		}

		if len(d.Events) > 0 {
			eventgen.PackageHeader()
			for _, e := range d.Events {
				eventgen.DomainEvent(d, e)
			}
		}

		// Write dom definitions into separate files.
		// domfile := fmt.Sprintf("%s.go", strings.ToLower(d.Domain))
		// tg.writeFile(domfile)
		// cg.writeFile(domfile)
		// eg.writeFile(domfile)
		// g.writeFile(domfile)
	}

	// Add a custom Timestamp type.
	typegen.Printf(`
// Timestamp represents a timestamp (since epoch).
type Timestamp `)
	typegen.domainTypeTime(proto.Domain{}, proto.AnyType{NameName: "Timestamp", Type: "number"})
	typegen.Printf("\n\n")

	typegen.writeFile(typegen.pkg + ".go")
	cmdgen.writeFile(cmdgen.pkg + ".go")
	eventgen.writeFile(eventgen.pkg + ".go")
	domgen.writeFile(domgen.pkg + ".go")
	cdpgen.writeFile(cdpgen.pkg + ".go")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	goimports := exec.CommandContext(ctx, "goimports", "-w", cdpgen.path(), typegen.path(), cmdgen.path(), eventgen.path())
	out, err := goimports.CombinedOutput()
	if err != nil {
		log.Printf("%s", out)
		log.Println(err)
		os.Exit(1)
	}

	goinstall := exec.CommandContext(ctx, "go", "install", path.Join(cdpgen.path(), "..."))
	out, err = goinstall.CombinedOutput()
	if err != nil {
		log.Printf("%s", out)
		log.Println(err)
		os.Exit(1)
	}
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	dir        string
	pkg        string
	imports    []string
	buf        bytes.Buffer // Accumulated output.
	hasContent bool
	hasHeader  bool
}

func (g *Generator) path() string {
	return path.Join(os.Getenv("GOPATH"), "src", g.dir)
}

// Printf prints to the Generator buffer.
func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) writeFile(f string) {
	fp := path.Join(g.path(), f)
	if !g.hasContent {
		log.Printf("No content, skipping %s...", fp)
		g.clear()
		return
	}
	if g.buf.Len() == 0 {
		log.Printf("Empty buffer, skipping %s...", fp)
		return
	}
	log.Printf("Writing %s...", fp)
	if err := ioutil.WriteFile(fp, g.format(), 0644); err != nil {
		panic(err)
	}
	g.clear()
}

func (g *Generator) clear() {
	g.hasContent = false
	g.hasHeader = false
	g.buf.Truncate(0)
}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

// CdpClient creates the cdp.Client type.
func (g *Generator) CdpClient(domains []proto.Domain) {
	g.hasContent = true
	var fields, newFields Generator
	for _, d := range domains {
		fields.Printf("\t%s %s\n", d.Name(), d.Type())
		newFields.Printf("\t\t%s: cdpdom.New%s(conn),\n", d.Name(), d.Name())
	}
	g.Printf(`
// Client represents a Chrome Debugging Protocol client that can be used to
// invoke methods or listen to events in every CDP domain. The Client consumes
// a rpcc connection, used to invoke the methods.
type Client struct {
	%s
}

// NewClient returns a new Client.
func NewClient(conn *rpcc.Conn) *Client {
	return &Client{
		%s
	}
}
`, fields.buf.Bytes(), newFields.buf.Bytes())
}

// PackageHeader writes the header for a package.
func (g *Generator) PackageHeader() {
	if g.hasHeader {
		return
	}
	g.hasHeader = true
	g.Printf(`// Code generated by cdpgen; DO NOT EDIT!

package %s

import (
	"context"
	"encoding/json"
	"fmt"

	%s
)
`, g.pkg, quotedImports(g.imports))
}

// DomainInterface defines the domain interface.
func (g *Generator) DomainInterface(d proto.Domain) {
	g.hasContent = true

	g.Printf(`
// The %[1]s domain. %[2]s
type %[1]s interface{`, d.Name(), d.Desc())
	for _, c := range d.Commands {
		request := ""
		reply := "error"
		if len(c.Parameters) > 0 {
			request = ", *cdpcmd." + c.ArgsName(d)
		}
		if len(c.Returns) > 0 {
			reply = fmt.Sprintf("(*cdpcmd.%s, error)", c.ReplyName(d))
		}
		g.Printf("\n\t// Command %s\n\t//\n\t// %s\n\t%s(context.Context%s) %s\n", c.Name(), c.Desc(true), c.Name(), request, reply)
	}
	for _, e := range d.Events {
		eventClient := fmt.Sprintf("%sClient", e.EventName(d))
		g.Printf("\n\t// Event %s\n\t//\n\t// %s\n\t%s(context.Context) (cdpevent.%s, error)\n", e.Name(), e.Desc(true), e.Name(), eventClient)
	}
	g.Printf("}\n")
}

// DomainDefinition defines the entire domain.
func (g *Generator) DomainDefinition(d proto.Domain) {
	g.hasContent = true

	g.Printf(`
// The %[1]s domain. %[2]s
type %[1]s struct{ conn *rpcc.Conn }

// New%[1]s returns the domain with the connection set to conn.
func New%[1]s(conn *rpcc.Conn) *%[1]s {
	return &%[1]s{conn: conn}
}
`, d.Name(), d.Desc())

	for _, c := range d.Commands {
		request := ""
		invokeReply := "nil"
		if len(c.Parameters) > 0 {
			request = ", args *cdpcmd." + c.ArgsName(d)
		}
		reply := "(err error)"
		if len(c.Returns) > 0 {
			reply = fmt.Sprintf("(reply *cdpcmd.%s, err error)", c.ReplyName(d))
		}
		g.Printf(`
// %[1]s invokes the %[2]s method. %[5]s
func (d *%[2]s) %[1]s(ctx context.Context%[3]s) %[4]s {`, c.Name(), d.Name(), request, reply, c.Desc(true))
		if len(c.Returns) > 0 {
			g.Printf(`
	reply = new(cdpcmd.%s)`, c.ReplyName(d))
			invokeReply = "reply"
		}
		if len(c.Parameters) > 0 {
			g.Printf(`
	if args != nil {
		err = rpcc.Invoke(ctx, cdpcmd.%[1]s.String(), args, %[2]s, d.conn)
	} else {
		err = rpcc.Invoke(ctx, cdpcmd.%[1]s.String(), nil, %[2]s, d.conn)
	}
	if err != nil {
		err = &OpError{Domain: %[3]q, Op: %[4]q, Err: err}
	}
	return
}
`, c.CmdName(d, true), invokeReply, d.Name(), c.Name())
		} else {
			g.Printf(`
	err = rpcc.Invoke(ctx, cdpcmd.%s.String(), nil, %s, d.conn)
	if err != nil {
		err = &OpError{Domain: %q, Op: %q, Err: err}
	}
	return
}
`, c.CmdName(d, true), invokeReply, d.Name(), c.Name())
		}
	}
	for _, e := range d.Events {
		eventClient := fmt.Sprintf("%sClient", e.EventName(d))

		// Implement event on domain.
		g.Printf(`
// %s creates the event client. %s
func (d *%s) %s(ctx context.Context) (cdpevent.%s, error) {
	s, err := rpcc.NewStream(ctx, cdpevent.%s.String(), d.conn)
	if err != nil {
		return nil, err
	}
	return &%s{Stream: s}, nil
}
`, e.Name(), e.Desc(true), d.Name(), e.Name(), eventClient, e.EventName(d), eventClient)

		g.Printf(`
// %[4]s implements cdpevent.%[1]s.
type %[4]s struct { rpcc.Stream }

// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
// triggered, context canceled or connection closed.
func (c *%[4]s) Recv() (*cdpevent.%[3]s, error) {
	event := new(cdpevent.%[3]s)
	if err := c.RecvMsg(event); err != nil {
		return nil, &OpError{Domain: %[5]q, Op: "%[6]s Recv", Err: err}
	}
	return event, nil
}
`, eventClient, "", e.ReplyName(d), eventClient, d.Name(), e.Name())

	}
}

// DomainType creates the type definition.
func (g *Generator) DomainType(d proto.Domain, t proto.AnyType) {
	g.hasContent = true
	g.Printf(`
// %[1]s %[2]s
type %[1]s `, t.Name(d), t.Desc())
	switch t.GoType(g.pkg, d) {
	case "struct":
		g.domainTypeStruct(d, t)
	case "enum":
		g.domainTypeEnum(d, t)
	case "time.Time":
		g.domainTypeTime(d, t)
	case "RawMessage":
		g.domainTypeRawMessage(d, t)
	default:
		g.Printf(t.GoType(g.pkg, d))
	}
	g.Printf("\n\n")
}

func (g *Generator) printStructProperties(d proto.Domain, name string, props []proto.AnyType, ptrOptional, renameOptional bool) {
	for _, prop := range props {
		jsontag := prop.NameName
		ptype := prop.GoType(g.pkg, d)
		// Make all optional properties into pointers, unless they are slices.
		if prop.Optional {
			isNonPtr := nonPtrMap[ptype]
			if ptrOptional && !isNonPtr && !isNonPointer(g.pkg, d, prop) {
				ptype = "*" + ptype
			}
			jsontag += ",omitempty"
		}

		// Avoid recursive type definitions.
		if ptype == name {
			ptype = "*" + ptype
		}

		exportedName := prop.ExportedName(d)
		if renameOptional && prop.Optional {
			exportedName = OptionalPropPrefix + exportedName
		}

		g.Printf("\t%s %s `json:\"%s\"` // %s\n", exportedName, ptype, jsontag, prop.Desc())
	}
}

func (g *Generator) domainTypeStruct(d proto.Domain, t proto.AnyType) {
	g.Printf("struct{\n")
	g.printStructProperties(d, t.Name(d), t.Properties, true, false)
	g.Printf("}")
}

func (g *Generator) domainTypeTime(d proto.Domain, t proto.AnyType) {
	g.Printf(`float64

// String calls (time.Time).String().
func (t %[1]s) String() string {
	return t.Time().String()
}

// Time parses the Unix time with millisecond accuracy.
func (t %[1]s) Time() time.Time {
	secs := int64(t)
	// The Unix time in t only has ms accuracy.
	ms := int64((float64(t)-float64(secs))*1000000)
	return time.Unix(secs, ms*1000)
}

// MarshalJSON implements json.Marshaler. Encodes to null if t is zero.
func (t %[1]s) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("null"), nil
	}
	f := float64(t)
	return json.Marshal(&f)
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *%[1]s) UnmarshalJSON(data []byte) error {
	*t = 0
	if len(data) == 0 {
		return nil
	}
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return errors.New("%[2]s.%[1]s: " + err.Error())
	}
	*t = %[1]s(f)
	return nil
}

var _ json.Marshaler = (*%[1]s)(nil)
var _ json.Unmarshaler = (*%[1]s)(nil)
`, t.Name(d), g.pkg)
}

func (g *Generator) domainTypeRawMessage(d proto.Domain, t proto.AnyType) {
	g.Printf(`[]byte

// MarshalJSON copies behavior of json.RawMessage.
func (%[3]s %[1]s) MarshalJSON() ([]byte, error) {
	if %[3]s == nil {
		return []byte("null"), nil
	}
	return %[3]s, nil
}

// UnmarshalJSON copies behavior of json.RawMessage.
func (%[3]s *%[1]s) UnmarshalJSON(data []byte) error {
	if %[3]s == nil {
		return errors.New("%[2]s.%[1]s: UnmarshalJSON on nil pointer")
	}
	*%[3]s = append((*%[3]s)[0:0], data...)
	return nil
}

var _ json.Marshaler = (*%[1]s)(nil)
var _ json.Unmarshaler = (*%[1]s)(nil)
`, t.Name(d), g.pkg, t.Recvr(d))
}

func (g *Generator) domainTypeEnum(d proto.Domain, t proto.AnyType) {
	if t.Type != "string" {
		log.Panicf("unknown enum type: %s", t.Type)
	}
	if realEnum {
		name := strings.Title(t.Name(d))
		g.Printf(`int

// %s as enums.
const (
	%sNotSet %s = iota`, name, name, name)
		for _, e := range t.Enum {
			g.Printf("\n\t%s%s", name, e.Name())
		}
		g.Printf(`
)

// Valid returns true if enum is set.
func (e %[1]s) Valid() bool {
	return e >= 1 && e <= %[2]d
}

func (e %[1]s) String() string {
	switch e {
	case 0:
		return "%[1]sNotSet"`, name, len(t.Enum))
		for i, e := range t.Enum {
			g.Printf(`
	case %d:
		return "%s"`, i+1, e)
		}
		g.Printf(`
	}
	return fmt.Sprintf("%[1]s(%%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e %[1]s) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("%[2]s.%[1]s: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *%[1]s) UnmarshalJSON(data []byte) error {
	if data == nil {
		*e = 0
		return nil
	}
	switch string(data) {`, name, g.pkg)
		for i, e := range t.Enum {
			g.Printf(`
	case "\"%s\"":
		*e = %d`, e, i+1)
		}
		g.Printf(`
	default:
		return fmt.Errorf("%s.%s: UnmarshalJSON on bad input: %%s", data)
	}
	return nil
}`, g.pkg, name)
	} else {
		g.Printf(`string

func (e %[1]s) String() string {
	return string(e)
}

// %[1]s types.
const (
`, t.Name(d))
		for _, e := range t.Enum {
			g.Printf("\t%s %s = %q\n", t.Name(d)+e.Name(), t.Name(d), e)
		}
		g.Printf(")")
	}
}

// CmdType generates the type for CDP methods names.
func (g *Generator) CmdType(doms []proto.Domain) {
	g.hasContent = true
	g.Printf(`
// CmdType is the type for CDP methods names.
type CmdType string

func (c CmdType) String() string {
	return string(c)
}

// Cmd methods.
const (`)
	for _, d := range doms {
		for _, c := range d.Commands {
			g.Printf("\n\t%s CmdType = %q", c.CmdName(d, true), d.Domain+"."+c.NameName)
		}
	}
	g.Printf("\n)\n")
}

// DomainCmd defines the command args and reply.
func (g *Generator) DomainCmd(d proto.Domain, c proto.Command) {
	if len(c.Parameters) > 0 {
		g.hasContent = true
		g.domainCmdArgs(d, c)
	}
	if len(c.Returns) > 0 {
		g.hasContent = true
		g.domainCmdReply(d, c)
	}
}

func (g *Generator) domainCmdArgs(d proto.Domain, c proto.Command) {
	g.Printf(`
// %[1]s represents the arguments for %[2]s in the %[3]s domain.
type %[1]s struct {`, c.ArgsName(d), c.Name(), d.Name())
	g.printStructProperties(d, c.ArgsName(d), c.Parameters, true, true)
	g.Printf("}\n\n")

	g.Printf(`
// New%[1]s initializes %[4]s with the required arguments.
func New%[1]s(%[2]s) *%[1]s {
	args := new(%[1]s)
	%[3]s
	return args
}
`, c.ArgsName(d), c.ArgsSignature(d), c.ArgsAssign("args", d), c.ArgsName(d))

	for _, arg := range c.Parameters {
		if !arg.Optional {
			continue
		}
		typ := arg.GoType(g.pkg, d)
		isNonPtr := nonPtrMap[typ]
		ptr := "&"
		if isNonPtr || isNonPointer(g.pkg, d, arg) {
			ptr = ""
		}
		name := arg.Name(d)
		if name == "range" || name == "type" {
			name = name[0 : len(name)-1]
		}
		g.Printf(`
// Set%[1]s sets the %[1]s optional argument. %[6]s
func (a *%[2]s) Set%[1]s(%[3]s %[4]s) *%[2]s {
	a.%[5]s%[1]s = %[7]s%[3]s
	return a
}
`, arg.ExportedName(d), c.ArgsName(d), name, arg.GoType("cdp", d), OptionalPropPrefix, arg.Desc(), ptr)
	}
}

func (g *Generator) domainCmdReply(d proto.Domain, c proto.Command) {
	g.Printf(`
// %[1]s represents the return values for %[2]s in the %[3]s domain.
type %[1]s struct {`, c.ReplyName(d), c.Name(), d.Name())
	g.printStructProperties(d, c.ReplyName(d), c.Returns, true, false)
	g.Printf("}\n\n")
}

// EventType generates the type for CDP event names.
func (g *Generator) EventType(doms []proto.Domain) {
	g.hasContent = true
	g.Printf(`
// EventType is the type for CDP event names.
type EventType string

func (e EventType) String() string {
	return string(e)
}

// Event methods.
const (`)
	for _, d := range doms {
		for _, e := range d.Events {
			g.Printf("\n\t%s EventType = %q", e.EventName(d), d.Domain+"."+e.NameName)
		}
	}
	g.Printf("\n)\n")
}

// DomainEvent defines the event client and reply.
func (g *Generator) DomainEvent(d proto.Domain, e proto.Event) {
	g.hasContent = true
	g.domainEventClient(d, e)
	g.domainEventReply(d, e)
}

func (g *Generator) domainEventClient(d proto.Domain, e proto.Event) {
	eventClient := fmt.Sprintf("%sClient", e.EventName(d))
	g.Printf(`
// %[1]s receives %[2]s events.
type %[1]s interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*%[3]s, error)
	rpcc.Stream
}
`, eventClient, e.Name(), e.ReplyName(d))
}

func (g *Generator) domainEventReply(d proto.Domain, e proto.Event) {
	g.Printf(`
// %[1]s %[2]s
type %[1]s struct {`, e.ReplyName(d), e.Desc(false))
	g.printStructProperties(d, e.ReplyName(d), e.Parameters, true, false)
	g.Printf("}\n")
}

func quotedImports(imports []string) string {
	if len(imports) == 0 {
		return ""
	}

	return "\"" + strings.Join(imports, "\"\n\"") + "\""
}

func isNonPointer(pkg string, d proto.Domain, t proto.AnyType) bool {
	typ := t.GoType(pkg, d)
	switch {
	case t.IsEnum():
	case strings.HasPrefix(typ, "[]"):
	case strings.HasPrefix(typ, "map["):
	case typ == "time.Time":
	case typ == "json.RawMessage":
	case typ == "RawMessage":
	case typ == "interface{}":
	default:
		return false
	}
	return true
}
