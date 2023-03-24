package devtool

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/mafredri/cdp/internal/errors"
)

// DevToolsOption represents a function that sets a DevTools option.
type DevToolsOption func(*DevTools)

// WithClient returns a DevToolsOption that sets the http Client used
// for HTTP GET requests.
func WithClient(client *http.Client) DevToolsOption {
	return func(d *DevTools) {
		d.client = client
	}
}

// DevTools represents a devtools endpoint for managing and querying
// information about targets.
type DevTools struct {
	url    string
	client *http.Client

	mu     sync.Mutex // Protects following.
	lookup bool
}

// New returns a DevTools instance that uses URL.
func New(url string, opts ...DevToolsOption) *DevTools {
	devtools := &DevTools{url: url}
	for _, o := range opts {
		o(devtools)
	}
	if devtools.client == nil {
		devtools.client = &http.Client{}
	}
	return devtools
}

// Type represents the type of Target.
type Type string

// Type enums.
const (
	BackgroundPage Type = "background_page"
	Node           Type = "node"
	Other          Type = "other"
	Page           Type = "page"
	ServiceWorker  Type = "service_worker"
)

// Target represents a devtools target, e.g. a browser tab.
type Target struct {
	Description          string `json:"description"`
	DevToolsFrontendURL  string `json:"devtoolsFrontendUrl"`
	ID                   string `json:"id"`
	Title                string `json:"title"`
	Type                 Type   `json:"type"`
	URL                  string `json:"url"`
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

// Create a new Target, usually a page with about:blank as URL.
func (d *DevTools) Create(ctx context.Context) (*Target, error) {
	return d.CreateURL(ctx, "")
}

// CreateURL is like Create but opens the provided URL. The URL must be
// valid and begin with "http://" or "https://".
func (d *DevTools) CreateURL(ctx context.Context, openURL string) (*Target, error) {
	var escapedQueryURL string

	if openURL != "" {
		if parsed, err := url.Parse(openURL); err != nil || !parsed.IsAbs() {
			return nil, errors.New("devtool: CreateURL: invalid openURL: " + openURL)
		}
		escapedQueryURL = "?" + url.QueryEscape(openURL)
	}

	resp, err := d.httpPut(ctx, "/json/new"+escapedQueryURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if ok, err := d.handleNodeUnsupportedMethod(ctx, resp, "CreateURL"); ok {
		return nil, err
	}

	switch resp.StatusCode {
	// Returned by Headless Chrome that does
	// not support the "/json/new" endpoint.
	case http.StatusInternalServerError:
		err2 := parseError("CreateURL: StatusInternalServerError", resp.Body)

		v, err := d.Version(ctx)
		if err != nil {
			return nil, err2
		}

		if v.WebSocketDebuggerURL != "" {
			// This version is too new since it has a debugger URL set.
			return nil, err2
		}

		return fallbackHeadlessCreateURL(ctx, d, openURL)

	case http.StatusOK:
		t := new(Target)
		return t, json.NewDecoder(resp.Body).Decode(t)

	default:
		return nil, parseError("CreateURL", resp.Body)
	}
}

// Get the first Target that matches Type.
func (d *DevTools) Get(ctx context.Context, typ Type) (*Target, error) {
	list, err := d.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, t := range list {
		if t.Type == typ {
			return t, nil
		}
	}

	return nil, errors.New("devtool: Get: could not find target of type: " + string(typ))
}

// List returns a list with all devtools Targets.
func (d *DevTools) List(ctx context.Context) ([]*Target, error) {
	resp, err := d.httpPut(ctx, "/json/list")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, parseError("List", resp.Body)
	}

	var t []*Target
	return t, json.NewDecoder(resp.Body).Decode(&t)
}

// Activate brings focus to the Target.
func (d *DevTools) Activate(ctx context.Context, t *Target) error {
	resp, err := d.httpPut(ctx, "/json/activate/"+t.ID)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if ok, err := d.handleNodeUnsupportedMethod(ctx, resp, "Activate"); ok {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return parseError("Activate", resp.Body)
	}

	return nil
}

// Close the Target.
func (d *DevTools) Close(ctx context.Context, t *Target) error {
	resp, err := d.httpPut(ctx, "/json/close/"+t.ID)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if ok, err := d.handleNodeUnsupportedMethod(ctx, resp, "Close"); ok {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return parseError("Close", resp.Body)
	}

	return nil
}

// Version contains the version information for the DevTools endpoint.
type Version struct {
	// Present in Chrome, Edge, Node, etc.
	Browser  string `json:"Browser"`
	Protocol string `json:"Protocol-Version"`

	// Present in Chrome, Edge.
	UserAgent string `json:"User-Agent"`
	V8        string `json:"V8-Version"`
	WebKit    string `json:"WebKit-Version"`

	// Present on Android.
	AndroidPackage string `json:"Android-Package"`

	// Present in Chrome >= 62. Generic browser websocket URL.
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

// Version returns the version information for the DevTools endpoint.
func (d *DevTools) Version(ctx context.Context) (*Version, error) {
	resp, err := d.httpPut(ctx, "/json/version")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, parseError("Version", resp.Body)
	}

	v := new(Version)
	return v, json.NewDecoder(resp.Body).Decode(&v)
}

func (d *DevTools) httpPut(ctx context.Context, path string) (*http.Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	err := d.resolveHost(ctx)
	if err != nil {
		return nil, err
	}

	// New versions of Chromium require PUT requests as a security measure.
	req, err := http.NewRequest(http.MethodPut, d.url+path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := d.client.Do(req.WithContext(ctx))
	if err == nil {
		return resp, nil
	}

	// Fallback to old method, use GET request.
	req, err = http.NewRequest(http.MethodGet, d.url+path, nil)
	if err != nil {
		return nil, err
	}

	return d.client.Do(req.WithContext(ctx))
}

// resolveHost does a lookup on the hostname in d.url and tries to
// replace it with a valid IP address. Ever since Chrome 66, the
// DevTools endpoint disallows hostnames other than "localhost".
//
// Example error:
// < HTTP/1.1 500 Internal Server Error
// < Content-Length:63
// < Content-Type:text/html
// <
// Host header is specified and is not an IP address or localhost.
func (d *DevTools) resolveHost(ctx context.Context) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.lookup {
		return nil
	}
	d.lookup = true

	u, err := url.Parse(d.url)
	if err != nil {
		return err
	}
	host := strings.Split(u.Host, ":")
	origHost := host[0]

	if origHost == "localhost" {
		return nil // Nothing to do, localhost is allowed.
	}

	addrs, err := net.DefaultResolver.LookupHost(ctx, origHost)
	if err != nil {
		return err
	}

	newURL := ""
	for _, a := range addrs {
		host[0] = a
		u.Host = strings.Join(host, ":")
		try := u.String()

		// The selection of "/json/version" here is arbitrary,
		// it just needs to exist and not have side-effects.
		req, err := http.NewRequest(http.MethodGet, try+"/json/version", nil)
		if err != nil {
			return err
		}

		resp, err := d.client.Do(req.WithContext(ctx))
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			newURL = try
			break
		}
	}
	if newURL == "" {
		return errors.New("could not resolve IP for " + origHost)
	}
	d.url = newURL

	return nil
}

// handleNodeUnsupportedMethod returns a nicer error message for
// unsupported methods when used with Node.js.
func (d *DevTools) handleNodeUnsupportedMethod(ctx context.Context, resp *http.Response, from string) (ok bool, err error) {
	if resp.StatusCode != http.StatusBadRequest {
		return false, nil
	}

	v, err := d.Version(ctx)
	if err != nil {
		return false, nil
	}

	if strings.HasPrefix(v.Browser, "node.js") {
		return true, parseError(fmt.Sprintf("%s: not supported by Node.js", from), resp.Body)
	}

	return false, nil
}

func parseError(from string, r io.Reader) error {
	m, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return errors.New("devtool: " + from + ": " + string(m))
}
