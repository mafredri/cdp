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
	realEnum           = false
)

var (
	nonPtrMap = make(map[string]bool)
)

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func mkdir(name string) error {
	err := os.Mkdir(name, 0755)
	if os.IsExist(err) {
		return nil
	}
	return err
}

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
	panicErr(err)

	err = json.Unmarshal(protocolData, &protocol)
	panicErr(err)

	jsProtocolData, err := ioutil.ReadFile(jsProtoFileJSON)
	panicErr(err)

	err = json.Unmarshal(jsProtocolData, &jsProtocol)
	panicErr(err)

	protocol.Domains = append(protocol.Domains, jsProtocol.Domains...)
	sort.Slice(protocol.Domains, func(i, j int) bool {
		return protocol.Domains[i].Domain < protocol.Domains[j].Domain
	})

	protoDest := path.Join(destPkg, "protocol")
	imports := []string{
		"github.com/mafredri/cdp/rpcc",
		"github.com/mafredri/cdp/protocol/internal",
		"github.com/mafredri/cdp/protocol",
	}
	for i, d := range protocol.Domains {
		dLower := strings.ToLower(d.Name())
		imports = append(imports, path.Join(protoDest, dLower))

		for ii, t := range d.Types {
			nam := t.Name(d)
			if isNonPointer(d.Domain, d, t) {
				nonPtrMap[nam] = true
				nonPtrMap[d.Domain+"."+nam] = true
			}

			// Reference the FrameId in the Frame type.
			if d.Domain == "Page" && t.IDName == "Frame" {
				for iii, p := range t.Properties {
					if p.NameName == "id" || p.NameName == "parentId" {
						p.Ref = "FrameId"
						t.Properties[iii] = p
					}
				}
			}
			d.Types[ii] = t
		}
		protocol.Domains[i] = d
	}
	nonPtrMap["Timestamp"] = true
	nonPtrMap["protocol.Timestamp"] = true
	nonPtrMap["TimeSinceEpoch"] = true
	nonPtrMap["network.TimeSinceEpoch"] = true
	nonPtrMap["MonotonicTime"] = true
	nonPtrMap["network.MonotonicTime"] = true

	var cdp, g Generator
	cdp.imports = imports
	g.imports = imports

	// Define the cdp Client.
	cdp.pkg = "cdp"
	cdp.dir = destPkg
	cdp.PackageHeader("")
	cdp.CdpClient(protocol.Domains)
	cdp.writeFile("cdp_client.go")

	// Package cdp/protocol.
	g.pkg = "protocol"
	g.dir = protoDest
	err = mkdir(g.path())
	panicErr(err)
	// Generate circular dependency types inside protocol package.
	for _, d := range protocol.Domains {
		if d.Name() == "Page" {
			d.Domain = "protocol"
			g.Printf("// +build !go1.9\n\n")
			g.PackageHeader("")
			for _, t := range d.Types {
				if t.Name(d) == "FrameID" || t.Name(d) == "ResourceType" {
					t.IDName = "Page" + t.IDName
					g.DomainType(d, t)
				}
			}
			g.writeFile("page18.go")
		}
	}

	// Package cdp/protocol/internal.
	g.pkg = "internal"
	g.dir = path.Join(protoDest, "internal")
	for _, d := range protocol.Domains {
		if d.Name() == "Page" {
			d.Domain = "internal"
			g.Printf("// +build go1.9\n\n")
			g.PackageHeader("")
			for _, t := range d.Types {
				if t.Name(d) == "FrameID" || t.Name(d) == "ResourceType" {
					idName := t.Name(d)
					t.IDName = "Page" + t.IDName
					t.Description = fmt.Sprintf("%s\n//\n// This type cannot be used directly. Use page.%s instead.", t.Description, idName)
					g.DomainType(d, t)
				}
			}
			g.writeFile("page19.go")
		}
	}

	// Generate the protocol definitions.
	for _, d := range protocol.Domains {
		cdp.PackageHeader("")
		cdp.DomainInterface(d)

		dLower := strings.ToLower(d.Domain)
		g.dir = path.Join(protoDest, dLower)
		g.pkg = dLower
		err := mkdir(g.path())
		panicErr(err)

		g.PackageHeader(fmt.Sprintf("// Package %s implements the %s domain. %s", dLower, d.Name(), d.Desc()))
		g.DomainDefinition(d)
		g.writeFile("domain.go")

		if len(d.Types) > 0 {
			g.PackageHeader("")
			for _, t := range d.Types {
				g.DomainType(d, t)
			}
			g.writeFile("types.go")
		}

		if len(d.Commands) > 0 {
			g.PackageHeader("")
			for _, c := range d.Commands {
				if c.Redirect != "" {
					continue
				}
				g.DomainCmd(d, c)
			}
			g.writeFile("command.go")
		}

		if len(d.Events) > 0 {
			g.PackageHeader("")
			for _, e := range d.Events {
				g.DomainEvent(d, e)
			}
			g.writeFile("event.go")
		}
	}

	cdp.writeFile(cdp.pkg + ".go")

	g.dir = destPkg

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	goimports := exec.CommandContext(ctx, "goimports", "-w", g.path())
	out, err := goimports.CombinedOutput()
	if err != nil {
		log.Printf("goimports failed: %s", out)
		log.Println(err)
		os.Exit(1)
	}

	goinstall := exec.CommandContext(ctx, "go", "install", path.Join(destPkg, "..."))
	out, err = goinstall.CombinedOutput()
	if err != nil {
		log.Printf("install failed: %s", out)
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
	testbuf    bytes.Buffer // Accumulated test output.
	hasContent bool
	hasHeader  bool

	// For managing circular types for go 1.8/1.9.
	writingType    bool
	isCircularType bool
	typebuf        bytes.Buffer // Accumulated type output.
	typebuf18      bytes.Buffer
	typebuf19      bytes.Buffer
}

func (g *Generator) path() string {
	return path.Join(os.Getenv("GOPATH"), "src", g.dir)
}

// Printf prints to the Generator buffer.
func (g *Generator) Printf(format string, args ...interface{}) {
	buf := &g.buf
	if g.isCircularType {
		fmt.Fprintf(&g.typebuf19, format, args...)
		fmt.Fprintf(&g.typebuf18, format, args...)
		return
	}
	if g.writingType {
		buf = &g.typebuf
	}
	fmt.Fprintf(buf, format, args...)
}

// Printf19 prints to the type buffer for Go 1.9+.
func (g *Generator) Printf19(format string, args ...interface{}) {
	fmt.Fprintf(&g.typebuf19, format, args...)
}

// Printf18 prints to the type buffer for Go 1.8-.
func (g *Generator) Printf18(format string, args ...interface{}) {
	fmt.Fprintf(&g.typebuf18, format, args...)
}

func (g *Generator) beginType() {
	g.writingType = true
	g.isCircularType = false
}
func (g *Generator) markCircularType() {
	if !g.isCircularType {
		_, err := g.typebuf19.Write(g.typebuf.Bytes())
		panicErr(err)
		_, err = g.typebuf18.Write(g.typebuf.Bytes())
		panicErr(err)
		g.typebuf.Reset()
	}
	g.isCircularType = true

}
func (g *Generator) commitType() {
	g.writingType = false
	if g.isCircularType {
		return
	}

	_, err := g.typebuf.WriteTo(&g.buf)
	panicErr(err)
	g.typebuf.Reset()
}

// TestPrintf prints to the Generator test buffer.
func (g *Generator) TestPrintf(format string, args ...interface{}) {
	// No-op
	// fmt.Fprintf(&g.testbuf, format, args...)
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
	err := ioutil.WriteFile(fp, g.format(), 0644)
	panicErr(err)

	if g.testbuf.Len() > 0 {
		g.buf.Reset()
		g.Printf("package %s\n\n", g.pkg)
		_, err = g.testbuf.WriteTo(&g.buf)
		panicErr(err)
		fptest := strings.Replace(fp, ".go", "_test.go", 1)
		log.Printf("Writing %s...", fptest)

		err = ioutil.WriteFile(fptest, g.format(), 0644)
		panicErr(err)
	}
	g.clear()

	if g.typebuf19.Len() > 0 {
		g.Printf("// +build go1.9\n\n")
		g.hasHeader = false
		g.PackageHeader("")
		_, err = g.typebuf19.WriteTo(&g.buf)
		panicErr(err)
		fp19 := strings.Replace(fp, ".go", "19.go", 1)
		log.Printf("Writing %s...", fp19)
		err = ioutil.WriteFile(fp19, g.format(), 0644)
		panicErr(err)
	}
	g.clear()

	if g.typebuf18.Len() > 0 {
		g.Printf("// +build !go1.9\n\n")
		g.hasHeader = false
		g.PackageHeader("")
		_, err = g.typebuf18.WriteTo(&g.buf)
		panicErr(err)
		fp18 := strings.Replace(fp, ".go", "18.go", 1)
		log.Printf("Writing %s...", fp18)
		err = ioutil.WriteFile(fp18, g.format(), 0644)
		panicErr(err)
	}
	g.clear()
}

func (g *Generator) clear() {
	g.hasContent = false
	g.hasHeader = false
	g.writingType = false
	g.isCircularType = false
	g.buf.Reset()
	g.testbuf.Reset()
	g.typebuf.Reset()
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
		newFields.Printf("\t\t%s: %s.NewClient(conn),\n", d.Name(), strings.ToLower(d.Name()))
	}
	g.Printf(`
// Client represents a Chrome Debugging Protocol client that can be used to
// invoke methods or listen to events in every CDP domain. The Client consumes
// a rpcc connection, used to invoke the methods.
type Client struct {
	%s
}

// NewClient returns a new Client that uses conn
// for communication with the debugging target.
func NewClient(conn *rpcc.Conn) *Client {
	return &Client{
		%s
	}
}
`, fields.buf.Bytes(), newFields.buf.Bytes())
}

// PackageHeader writes the header for a package.
func (g *Generator) PackageHeader(comment string) {
	if g.hasHeader {
		return
	}
	g.hasHeader = true
	g.Printf(`// Code generated by cdpgen. DO NOT EDIT.

%s
package %s

import (
	"context"
	"encoding/json"
	"fmt"

	%s
)
`, comment, g.pkg, quotedImports(g.imports))
}

// DomainInterface defines the domain interface.
func (g *Generator) DomainInterface(d proto.Domain) {
	g.hasContent = true

	desc := d.Desc()
	if d.Deprecated {
		desc = "\n//\n// Deprecated: " + desc
	}
	if d.Experimental {
		desc += "\n//\n// Note: This domain is experimental."
	}
	g.Printf(`
// The %[1]s domain. %[2]s
type %[1]s interface{`, d.Name(), desc)
	for _, c := range d.Commands {
		if c.Redirect != "" {
			continue
		}
		request := ""
		reply := "error"
		if len(c.Parameters) > 0 {
			request = ", *" + strings.ToLower(d.Name()) + "." + c.ArgsName(d)
		}
		if len(c.Returns) > 0 {
			reply = fmt.Sprintf("(*%s.%s, error)", strings.ToLower(d.Name()), c.ReplyName(d))
		}
		desc := c.Desc(true)
		if c.Deprecated {
			desc = strings.Replace(desc, "Deprecated, ", "", 1)
			desc = "Deprecated: " + strings.ToUpper(desc[0:1]) + desc[1:]
		}
		if desc != "" {
			desc = "\n\t//\n\t// " + desc
		}
		if c.Experimental {
			desc += "\n\t//\n\t// Note: This command is experimental."
		}
		g.Printf("\n\t// Command %s%s\n\t%s(context.Context%s) %s\n", c.Name(), desc, c.Name(), request, reply)
	}
	for _, e := range d.Events {
		eventClient := fmt.Sprintf("%sClient", e.EventName(d))
		desc := e.Desc(true)
		if e.Deprecated {
			desc = strings.Replace(desc, "Deprecated, ", "", 1)
			desc = "Deprecated: " + strings.ToUpper(desc[0:1]) + desc[1:]
		}
		if desc != "" {
			desc = "\n\t//\n\t// " + desc
		}
		if e.Experimental {
			desc += "\n//\n// Note: This event is experimental."
		}
		g.Printf("\n\t// Event %s%s\n\t%s(context.Context) (%s.%s, error)\n", e.Name(), desc, e.Name(), strings.ToLower(d.Name()), eventClient)
	}
	g.Printf("}\n")
}

// DomainDefinition defines the entire domain.
func (g *Generator) DomainDefinition(d proto.Domain) {
	g.hasContent = true

	g.Printf(`
// %[1]s is a client for the %[2]s domain. %[3]s
type %[1]s struct{ conn *rpcc.Conn }

// NewClient returns a client for the %[2]s domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *%[1]s {
	return &%[1]s{conn: conn}
}
`, "domainClient", d.Name(), d.Desc())

	for _, c := range d.Commands {
		if c.Redirect != "" {
			continue
		}
		request := ""
		invokeReply := "nil"
		if len(c.Parameters) > 0 {
			request = ", args *" + c.ArgsName(d)
		}
		reply := "(err error)"
		if len(c.Returns) > 0 {
			reply = fmt.Sprintf("(reply *%s, err error)", c.ReplyName(d))
		}
		g.Printf(`
// %[1]s invokes the %[2]s method. %[5]s
func (d *domainClient) %[1]s(ctx context.Context%[3]s) %[4]s {`, c.Name(), d.Name(), request, reply, c.Desc(true))
		if len(c.Returns) > 0 {
			g.Printf(`
	reply = new(%s)`, c.ReplyName(d))
			invokeReply = "reply"
		}
		if len(c.Parameters) > 0 {
			g.Printf(`
	if args != nil {
		err = rpcc.Invoke(ctx, %[1]q, args, %[2]s, d.conn)
	} else {
		err = rpcc.Invoke(ctx, %[1]q, nil, %[2]s, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: %[3]q, Op: %[4]q, Err: err}
	}
	return
}
`, d.Domain+"."+c.NameName, invokeReply, d.Name(), c.Name())
		} else {
			g.Printf(`
	err = rpcc.Invoke(ctx, %q, nil, %s, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: %q, Op: %q, Err: err}
	}
	return
}
`, d.Domain+"."+c.NameName, invokeReply, d.Name(), c.Name())
		}

		// Generate method tests.
		g.TestPrintf(`
func Test%[1]s_%[2]s(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := New%[1]s(conn)
	var err error
`, d.Name(), c.Name())
		assign := "err"
		if len(c.Returns) > 0 {
			assign = "_, err"
		}
		if len(c.Parameters) > 0 {
			g.TestPrintf(`
	// Test nil args.
	%[1]s = dom.%[2]s(nil, nil)
	if err != nil {
		t.Error(err)
	}
	// Test args.
	%[1]s = dom.%[2]s(nil, &%[3]s{})`, assign, c.Name(), c.ArgsName(d))
		} else {
			g.TestPrintf(`
	%[1]s = dom.%[2]s(nil)`, assign, c.Name())
		}
		g.TestPrintf(`
	if err != nil {
		t.Error(err)
	}

	// Test error.
	codec.respErr = errors.New("bad request")`)
		if len(c.Parameters) > 0 {
			g.TestPrintf(`
	%[1]s = dom.%[2]s(nil, &%[3]s{})`, assign, c.Name(), c.ArgsName(d))
		} else {
			g.TestPrintf(`
	%[1]s = dom.%[2]s(nil)`, assign, c.Name())
		}
		g.TestPrintf(`
	if err == nil || err.(*internal.OpError).Err.(*rpcc.ResponseError).Message != codec.respErr.Error() {
		t.Errorf("unexpected error; got: %%v, want bad request", err)
	}`)
		g.TestPrintf(`
}
`)

	}
	for _, e := range d.Events {
		eventClient := fmt.Sprintf("%sClient", e.EventName(d))
		eventClientImpl := strings.ToLower(string(eventClient[0])) + eventClient[1:]

		// Implement event on domain.
		g.Printf(`
func (d *domainClient) %s(ctx context.Context) (%s, error) {
	s, err := rpcc.NewStream(ctx, %q, d.conn)
	if err != nil {
		return nil, err
	}
	return &%s{Stream: s}, nil
}
`, e.Name(), eventClient, d.Domain+"."+e.NameName, eventClientImpl)

		g.Printf(`
type %[4]s struct { rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *%[4]s) GetStream() rpcc.Stream { return c.Stream }

func (c *%[4]s) Recv() (*%[3]s, error) {
	event := new(%[3]s)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: %[5]q, Op: "%[6]s Recv", Err: err}
	}
	return event, nil
}
`, eventClient, "", e.ReplyName(d), eventClientImpl, d.Name(), e.Name())

		// Generate event tests.
		g.TestPrintf(`
func Test%[1]s_%[2]s(t *testing.T) {
	conn, codec, cleanup := newTestConn(t)
	defer cleanup()

	dom := New%[1]s(conn)

	stream, err := dom.%[2]s(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()

	codec.event = %[3]s.String()
	codec.conn <- nil
	_, err = stream.Recv()
	if err != nil {
		t.Error(err)
	}

	codec.eventArgs = []byte("invalid json")
	codec.conn <- nil
	_, err = stream.Recv()
	if err, ok := err.(*internal.OpError); !ok {
		t.Errorf("Recv() got %%v, want internal.OpError", err)
	}

	conn.Close()
	stream, err = dom.%[2]s(nil)
	if err == nil {
		t.Errorf("Open stream: got nil, want error")
	}
`, d.Name(), e.Name(), e.EventName(d))
		g.TestPrintf(`
}
`)
	}
}

// DomainType creates the type definition.
func (g *Generator) DomainType(d proto.Domain, t proto.AnyType) {
	g.hasContent = true

	g.beginType()
	var comment string
	if d.Name() == "Page" && (t.Name(d) == "FrameID" || t.Name(d) == "ResourceType") {
		g.markCircularType()
		comment = "//"
		g.Printf19(`
// %[1]s %[2]s
//
// Provided as an alias to prevent circular dependencies.
type %[1]s = internal.Page%[1]s
`, t.Name(d), t.Desc())
	}
	desc := t.Desc()
	if t.Deprecated {
		desc = "\n//\n// Deprecated: " + desc
	}
	if t.Experimental {
		desc += "\n//\n// Note: This type is experimental."
	}
	g.Printf(`
// %[1]s %[2]s
%[3]stype %[1]s `, t.Name(d), desc, comment)

	typ := t.GoType(g.pkg, d)
	switch typ {
	case "struct":
		g.domainTypeStruct(d, t)
	case "enum":
		g.domainTypeEnum(d, t)
	case "time.Time":
		g.domainTypeTime(d, t)
	case "RawMessage":
		g.domainTypeRawMessage(d, t)
	default:
		g.Printf(typ)
	}
	g.Printf("\n\n")
	g.commitType()
}

func (g *Generator) printStructProperties(d proto.Domain, name string, props []proto.AnyType, ptrOptional, renameOptional bool) {
	for _, prop := range props {
		jsontag := prop.NameName
		ptype := prop.GoType(g.pkg, d)
		ptype19 := ptype
		ptype18 := ptype

		if ptype == "page.FrameID" || ptype == "page.ResourceType" {
			g.markCircularType()
			if g.pkg == "network" || g.pkg == "dom" {
				// Domain-local type (alias) for go1.9.
				ptype19 = strings.Replace(ptype19, "page.", "internal.Page", 1)
			}
			ptype18 = strings.Replace(ptype18, ".", "", 1)
			ptype18 = "protocol." + strings.Title(ptype18)
		}
		if g.pkg == "page" && (ptype == "FrameID" || ptype == "ResourceType") {
			g.markCircularType()
			ptype18 = "protocol.Page" + ptype18
		}
		// Make all optional properties into pointers, unless they are slices.
		if prop.Optional {
			isNonPtr := nonPtrMap[ptype]
			if ptrOptional && !isNonPtr && !isNonPointer(g.pkg, d, prop) {
				ptype = "*" + ptype
				ptype19 = "*" + ptype19
				ptype18 = "*" + ptype18
			}
			jsontag += ",omitempty"
		}

		// Avoid recursive type definitions.
		if ptype == name {
			ptype = "*" + ptype
			ptype19 = "*" + ptype19
			ptype18 = "*" + ptype18
		}

		exportedName := prop.ExportedName(d)
		if renameOptional && prop.Optional {
			exportedName = OptionalPropPrefix + exportedName
		}

		var preDesc, postDesc string

		desc := prop.Desc()
		var deprecated, localEnum, experimental string
		if prop.Deprecated {
			if desc == "" {
				desc = "This property should not be used."
			}
			deprecated = "//\n// Deprecated: " + desc + "\n"
			desc = "is deprecated."
		}
		if prop.IsLocalEnum() {
			var enums []string
			for _, e := range prop.Enum {
				enums = append(enums, fmt.Sprintf("%q", e))
			}
			localEnum = "//\n// Values: " + strings.Join(enums, ", ") + ".\n"
		}
		if prop.Experimental {
			experimental = "//\n// Note: This property is experimental.\n"
		}
		if deprecated != "" || localEnum != "" || experimental != "" {
			preDesc = "// " + exportedName + " " + desc + "\n" + deprecated + localEnum + experimental
		} else {
			if desc == "" {
				desc = "No description."
			}
			postDesc = "// " + desc
		}

		if !g.isCircularType {
			g.Printf("\t%s%s %s `json:\"%s\"` %s\n", preDesc, exportedName, ptype, jsontag, postDesc)
		} else {
			g.Printf19("\t%s%s %s `json:\"%s\"` %s\n", preDesc, exportedName, ptype19, jsontag, postDesc)
			g.Printf18("\t%s%s %s `json:\"%s\"` %s\n", preDesc, exportedName, ptype18, jsontag, postDesc)
		}
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

	g.TestPrintf(`
func Test%[1]s_Marshal(t *testing.T) {
	var v %[1]s

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %%v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %%s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %%v, want no error", err)
	}

	// Test non-empty.
	v = 1
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %%v, want no error", err)
	}
	if string(b) != "1" {
		t.Errorf("Marshal() got %%s, want 1", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %%v, want no error", err)
	}
}
`, t.Name(d))
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

	g.TestPrintf(`
func Test%[1]s_Marshal(t *testing.T) {
	var v %[1]s

	// Test empty.
	b, err := json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %%v, want no error", err)
	}
	if string(b) != "null" {
		t.Errorf("Marshal() got %%s, want null", b)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %%v, want no error", err)
	}

	// Test non-empty.
	v = []byte("\"test\"")
	b, err = json.Marshal(&v)
	if err != nil {
		t.Errorf("Marshal() got %%v, want no error", err)
	}
	if !bytes.Equal(v, b) {
		t.Errorf("Marshal() got %%s, want %%s", b, v)
	}
	v = nil
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Errorf("Unmarshal() got %%v, want no error", err)
	}
	if !bytes.Equal(v, b) {
		t.Errorf("Unmarshal() got %%s, want %%s", b, v)
	}
}
`, t.Name(d))
}

func (g *Generator) domainTypeEnum(d proto.Domain, t proto.AnyType) {
	isCircular := g.isCircularType

	if t.Type != "string" {
		log.Panicf("unknown enum type: %s", t.Type)
	}

	// Verify assumption about enums never being the empty string, this
	// allows us to consider optional enums as type string instead of
	// *string when encoding to JSON (omitempty).
	for _, e := range t.Enum {
		if e == "" {
			panic("enum " + t.Name(d) + " has unexpected empty enum value")
		}
	}

	name := strings.Title(t.Name(d))
	if realEnum {
		g.Printf("int\n\n")

		if name != "PageResourceType" {
			format := `
// %s as enums.
const (
	%sNotSet %s = iota`
			if !isCircular {
				g.Printf(format, name, name, name)
			} else {
				g.Printf19(format, name, name, name)
				g.Printf18(format, name, name, "protocol."+d.Name()+name)
			}
			for _, e := range t.Enum {
				g.Printf("\n\t%s%s", name, e.Name())
			}
			g.Printf(`
)
`)
			if isCircular {
				g.commitType()
				g.beginType()
			}
		}
		g.Printf(`
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
	switch string(data) {
	case "null":
		*e = 0`, name, g.pkg)
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
		g.Printf("string\n\n")

		if name != "PageResourceType" {
			g.Printf(`
// %s as enums.
const (`, name)
			format := "\n\t%s%s %s = %q"
			if !isCircular {
				g.Printf(format, name, "NotSet", name, "")
			} else {
				g.Printf19(format, name, "NotSet", name, "")
				g.Printf18(format, name, "NotSet", "protocol."+d.Name()+name, "")
			}

			for _, e := range t.Enum {
				if !isCircular {
					g.Printf(format, name, e.Name(), name, e)
				} else {
					g.Printf19(format, name, e.Name(), name, e)
					g.Printf18(format, name, e.Name(), "protocol."+d.Name()+name, e)
				}
			}
			g.Printf(`
)
`)
			if isCircular {
				g.commitType()
				g.beginType()
			}
		}

		var enumValues []string
		for _, e := range t.Enum {
			enumValues = append(enumValues, fmt.Sprintf("%q", e))
		}

		g.Printf(`
func (e %[1]s) Valid() bool {
	switch e {
	case %[2]s:
		return true
	default:
		return false
	}
}

func (e %[1]s) String() string {
	return string(e)
}
`, t.Name(d), strings.Join(enumValues, ", "))
	}

	if isCircular {
		g.typebuf.Reset()
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
	g.beginType()
	if len(c.Parameters) > 0 {
		g.hasContent = true
		g.domainCmdArgs(d, c)
	}
	g.commitType()
	g.beginType()
	if len(c.Returns) > 0 {
		g.hasContent = true
		g.domainCmdReply(d, c)
	}
	g.commitType()
}

func (g *Generator) domainCmdArgs(d proto.Domain, c proto.Command) {
	g.Printf(`
// %[1]s represents the arguments for %[2]s in the %[3]s domain.
type %[1]s struct {
`, c.ArgsName(d), c.Name(), d.Name())
	g.printStructProperties(d, c.ArgsName(d), c.Parameters, true, true)
	g.Printf("}\n\n")

	newfmt := `
// New%[1]s initializes %[4]s with the required arguments.
func New%[1]s(%[2]s) *%[1]s {
	args := new(%[1]s)
	%[3]s
	return args
}
`
	if !g.isCircularType {
		g.Printf(newfmt, c.ArgsName(d), c.ArgsSignature(d), c.ArgsAssign("args", d), c.ArgsName(d))
	} else {
		sig := c.ArgsSignature(d)
		g.Printf19(newfmt, c.ArgsName(d), sig, c.ArgsAssign("args", d), c.ArgsName(d))
		sig18 := strings.Replace(sig, "page.FrameID", "protocol.PageFrameID", 1)
		sig18 = strings.Replace(sig18, "page.ResourceType", "protocol.PageResourceType", 1)
		if d.Name() == "Page" {
			sig18 = strings.Replace(sig, "FrameID", "protocol.PageFrameID", 1)
			sig18 = strings.Replace(sig18, "ResourceType", "protocol.PageResourceType", 1)
		}
		g.Printf18(newfmt, c.ArgsName(d), sig18, c.ArgsAssign("args", d), c.ArgsName(d))
	}

	// Test the new arguments.
	testInit := ""
	if c.ArgsSignature(d) != "" {
		testInit = fmt.Sprintf("func() (%s) { return }()", c.ArgsSignature(d))
	}
	g.TestPrintf(`
func TestNew%[1]s(t *testing.T) {
	args := New%[1]s(%[2]s)
	if args == nil {
		t.Errorf("New%[1]s returned nil args")
	}
}
`, c.ArgsName(d), testInit)

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
		desc := arg.Desc()
		if arg.Deprecated {
			if desc == "" {
				desc = "This property should not be used."
			}
			desc = "\n//\n// Deprecated: " + desc
		}
		if arg.IsLocalEnum() {
			var enums []string
			for _, e := range arg.Enum {
				enums = append(enums, fmt.Sprintf("%q", e))
			}
			desc += "\n//\n// Values: " + strings.Join(enums, ", ") + "."
		}
		if arg.Experimental {
			desc += "\n//\n// Note: This property is experimental."
		}
		g.Printf(`
// Set%[1]s sets the %[1]s optional argument. %[6]s
func (a *%[2]s) Set%[1]s(%[3]s %[4]s) *%[2]s {
	a.%[5]s%[1]s = %[7]s%[3]s
	return a
}
`, arg.ExportedName(d), c.ArgsName(d), name, arg.GoType("cdp", d), OptionalPropPrefix, desc, ptr)
	}
}

func (g *Generator) domainCmdReply(d proto.Domain, c proto.Command) {
	g.Printf(`
// %[1]s represents the return values for %[2]s in the %[3]s domain.
type %[1]s struct {
`, c.ReplyName(d), c.Name(), d.Name())
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

	g.beginType()
	g.domainEventReply(d, e)
	g.commitType()
}

func (g *Generator) domainEventClient(d proto.Domain, e proto.Event) {
	eventClient := fmt.Sprintf("%sClient", e.EventName(d))
	g.Printf(`
// %[1]s is a client for %[2]s events. %[4]s
type %[1]s interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*%[3]s, error)
	rpcc.Stream
}
`, eventClient, e.Name(), e.ReplyName(d), e.Desc(true))
}

func (g *Generator) domainEventReply(d proto.Domain, e proto.Event) {
	g.Printf(`
// %[1]s is the reply for %[2]s events.
type %[1]s struct {`, e.ReplyName(d), e.Name())
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
