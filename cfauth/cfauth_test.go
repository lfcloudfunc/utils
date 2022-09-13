package cfauth_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/jfcote87/ctxclient"
	"github.com/lfcloudfunc/utils/cfauth"
)

type testHandler struct {
	u string
}

func (th *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	scopes := r.URL.Query().Get("scopes")

	switch scopes {
	case "":
		http.Error(w, "First", 404)
		return
	case "D":
		w.Write([]byte("{0"))
		return
	case "A,B,C":
		json.NewEncoder(w).Encode(map[string]string{
			"access_token":  "A Token",
			"refresh_token": "Refresh",
			"type":          "Bearer",
		})
		return
	}
	http.Error(w, "Invalid scope", 400)
}

func (th *testHandler) RoundTrip(r *http.Request) (*http.Response, error) {

	ux, _ := url.Parse(th.u + r.URL.Path)
	rx := *r
	ux.RawQuery = r.URL.RawQuery
	rx.URL = ux
	return http.DefaultTransport.RoundTrip(&rx)
}

func TestTokensource_Token(t *testing.T) {
	hx := &testHandler{}
	srv := httptest.NewServer(hx)
	hx.u = srv.URL
	cl := &http.Client{Transport: hx}
	var auths = cfauth.Tokensource{
		Func: func(ctx context.Context) (*http.Client, error) {
			return cl, nil
		},
		Scopes: []string{},
	}
	var werr00 *ctxclient.NotSuccess
	_, err := auths.Token(context.Background())
	if !errors.As(err, &werr00) {
		t.Errorf("expected 404 error; got %v", err)
	}
	auths.Scopes = append(auths.Scopes, "D")
	var werr01 *json.SyntaxError
	_, err = auths.Token(context.Background())
	if !errors.As(err, &werr01) {
		t.Errorf("expected json parse error; got %v", err)
	}
	auths.Scopes = []string{"A", "B", "C"}
	tk, err := auths.Token(context.Background())
	if err != nil || tk.AccessToken != "A Token" {
		t.Errorf("expected access token \"A Token\"; got %v", func() error {
			if err == nil {
				err = fmt.Errorf("%s", tk.AccessToken)
			}
			return err
		}())
	}
}
