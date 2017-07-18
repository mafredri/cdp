package runtime_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/mafredri/cdp/protocol/runtime"
)

var update = flag.Bool("update", false, "update .golden files")

func TestRuntimeRemoteObject_String(t *testing.T) {
	in := filepath.Join("testdata", "log.input")
	input, err := ioutil.ReadFile(in)
	if err != nil {
		t.Fatal(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(input))
	var buf bytes.Buffer

	for scanner.Scan() {
		var logEv runtime.ConsoleAPICalledReply
		err := json.Unmarshal(scanner.Bytes(), &logEv)
		if err != nil {
			t.Fatal(err)
		}
		for _, arg := range logEv.Args {
			buf.WriteString(arg.String())
			buf.WriteByte('\n')
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}

	out := filepath.Join("testdata", "log.golden")
	want, err := ioutil.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}

	if got := buf.Bytes(); !bytes.Equal(got, want) {
		if *update {
			err := ioutil.WriteFile(out, got, 0666)
			if err != nil {
				t.Error(err)
			}
			return
		}
		t.Error("output does not match golden file")
		showDiff(t, got, want)
	}
}

func showDiff(t testing.TB, got, want []byte) {
	gr := bufio.NewReader(bytes.NewReader(got))
	wr := bufio.NewReader(bytes.NewReader(want))
	var lineno int
	for {
		lineno++
		g, err1 := gr.ReadBytes('\n')
		w, err2 := wr.ReadBytes('\n')
		if err1 != nil || err2 != nil {
			if err1 != io.EOF || err2 != io.EOF {
				t.Error(err1, err2)
			}
			break
		}
		if !bytes.Equal(g, w) {
			t.Errorf("line %d: got %s; want %s", lineno, g[:len(g)-1], w[:len(w)-1])
		}
	}
}
