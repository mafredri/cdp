package devtool

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

var update = flag.Bool("update", false, "update .golden files")

type testHandler struct {
	status int
	body   []byte
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.RequestURI, "/json") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Only /json endpoint is supported!"))
		return
	}

	w.WriteHeader(h.status)
	w.Write(h.body)
}

func read(t *testing.T, name string) []byte {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestDevTools_WithClient(t *testing.T) {
	c := &http.Client{}
	devt := New("", WithClient(c))
	if devt.client != c {
		t.Error("DevTools client was not set")
	}
}

func TestDevTools(t *testing.T) {
	th := &testHandler{}
	srv := httptest.NewServer(th)
	defer srv.Close()

	devt := New(srv.URL)

	var buf bytes.Buffer

	tests := []struct {
		name   string
		status int
		body   []byte
		fn     func() (interface{}, error)
	}{
		{"Create", http.StatusOK, read(t, filepath.Join("testdata", "new.json")), func() (interface{}, error) {
			target, err := devt.Create(nil)
			return target, err
		}},
		{"Get", http.StatusOK, read(t, filepath.Join("testdata", "list.json")), func() (interface{}, error) {
			target, err := devt.Get(nil, Page)
			return target, err
		}},
		{"Close", http.StatusOK, []byte("Target is closing"), func() (interface{}, error) {
			return nil, devt.Close(nil, &Target{ID: "ddd908ca-4d8c-4783-a089-c9456c463eef"})
		}},
		{"Activate", http.StatusOK, []byte("Target activated"), func() (interface{}, error) {
			return nil, devt.Activate(nil, &Target{ID: "ddd908ca-4d8c-4783-a089-c9456c463eef"})
		}},
		{"Version", http.StatusOK, read(t, filepath.Join("testdata", "version.json")), func() (interface{}, error) {
			v, err := devt.Version(nil)
			return v, err
		}},
	}

	for _, tt := range tests {
		th.status = tt.status
		th.body = tt.body
		v, err := tt.fn()
		fmt.Fprintf(&buf, "%s: %v %v\n", tt.name, v, err)
	}

	out := filepath.Join("testdata", "test.golden")
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

func TestDevTools_Error(t *testing.T) {
	th := &testHandler{}
	srv := httptest.NewServer(th)
	defer srv.Close()

	devt := New(srv.URL)

	var buf bytes.Buffer

	tests := []struct {
		name   string
		status int
		body   []byte
		fn     func() (interface{}, error)
	}{
		{"Create", http.StatusNotFound, []byte("Not found"), func() (interface{}, error) {
			target, err := devt.Create(nil)
			return target, err
		}},
		{"Get", http.StatusNotFound, []byte("Not found"), func() (interface{}, error) {
			target, err := devt.Get(nil, Page)
			return target, err
		}},
		{"Get ServiceWorker", http.StatusOK, read(t, filepath.Join("testdata", "list.json")), func() (interface{}, error) {
			target, err := devt.Get(nil, ServiceWorker)
			return target, err
		}},
		{"Close", http.StatusNotFound, []byte("Could not close target id: ddd908ca-4d8c-4783-a089-c9456c463eef"), func() (interface{}, error) {
			return nil, devt.Close(nil, &Target{ID: "ddd908ca-4d8c-4783-a089-c9456c463eef"})
		}},
		{"Activate", http.StatusNotFound, []byte("Could not close target id: ddd908ca-4d8c-4783-a089-c9456c463eef"), func() (interface{}, error) {
			return nil, devt.Activate(nil, &Target{ID: "ddd908ca-4d8c-4783-a089-c9456c463eef"})
		}},
		{"Version", http.StatusNotFound, []byte("Not found"), func() (interface{}, error) {
			v, err := devt.Version(nil)
			return v, err
		}},
	}

	for _, tt := range tests {
		th.status = tt.status
		th.body = tt.body
		v, err := tt.fn()
		fmt.Fprintf(&buf, "%s: %v %v\n", tt.name, v, err)
	}

	out := filepath.Join("testdata", "error.golden")
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
