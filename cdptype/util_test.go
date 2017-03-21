package cdptype_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/mafredri/cdp/cdpevent"
)

var update = flag.Bool("update", false, "update .golden files")

func open(t testing.TB, name string) (*os.File, func()) {
	f, err := os.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return f, func() { f.Close() }
}

func TestRuntimeRemoteObject_String(t *testing.T) {
	inputFile, inputClose := open(t, filepath.Join("testdata", "log.input"))
	defer inputClose()

	input := bufio.NewScanner(inputFile)
	var buf bytes.Buffer

	for input.Scan() {
		var logEv cdpevent.RuntimeConsoleAPICalledReply
		err := json.Unmarshal(input.Bytes(), &logEv)
		if err != nil {
			t.Error(err)
		}
		for _, arg := range logEv.Args {
			buf.WriteString(arg.String())
			buf.WriteByte('\n')
		}
	}

	if input.Err() != nil {
		t.Error(input.Err())
	}

	gname := filepath.Join("testdata", "log.golden")
	if *update {
		err := ioutil.WriteFile(gname, buf.Bytes(), 0666)
		if err != nil {
			t.Fatal(err)
		}
		return
	}

	goldenFile, goldenClose := open(t, gname)
	defer goldenClose()

	var lineno int
	br := bufio.NewReader(&buf)
	gr := bufio.NewReader(goldenFile)
	for {
		lineno++
		b, err1 := br.ReadBytes('\n')
		g, err2 := gr.ReadBytes('\n')
		if err1 != nil || err2 != nil {
			if err1 != io.EOF || err2 != io.EOF {
				t.Error(err1, err2)
			}
			break
		}
		if !bytes.Equal(b, g) {
			t.Errorf("line %d does not match golden file", lineno)
			t.Errorf("got %s; want %s", b[:len(b)-1], g[:len(g)-1])
		}
	}
}
