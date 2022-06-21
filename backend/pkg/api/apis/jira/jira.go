package pollon

import (
	"fmt"
	"net/http"

	"github.com/andygrunwald/go-jira"
)

type ClientFactory interface {
	NewJiraClient() (*jira.Client, error)
}

func NewTotClientFactory() ClientFactory {
	return &clientFactory{}
}

type clientFactory struct {
}

func (t *clientFactory) NewJiraClient() (*jira.Client, error) {
	transport := TokenAuthTransport{Token: "NzM3MzQ4NjI2NDIyOoZY5Wua1BN1VURUEjre6zI68KUC"}
	return jira.NewClient(transport.Client(), "https://issues.redhat.com")
}

type TokenAuthTransport struct {
	Token string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

func (t *TokenAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := cloneRequest(req) // per RoundTripper contract
	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.Token))
	return t.transport().RoundTrip(req2)
}

func (t *TokenAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *TokenAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}
