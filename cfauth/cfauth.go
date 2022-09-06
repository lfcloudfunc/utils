package cfauth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/jfcote87/ctxclient"
	"github.com/jfcote87/oauth2"
)

// Tokensource returns access tokens for the default Service Account
type Tokensource struct {
	Scopes []string
	ctxclient.Func
}

const serviceAccountTokenURL = "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/token"

// Token retrieves an access_token from GCE
func (ts *Tokensource) Token(ctx context.Context) (*oauth2.Token, error) {
	r, _ := http.NewRequest("GET", serviceAccountTokenURL, nil)
	r.Header.Set("Metadata-Flavor", "Google")
	if len(ts.Scopes) > 0 {
		r.URL.RawQuery = url.Values{"scopes": []string{strings.Join(ts.Scopes, ",")}}.Encode()
	}
	res, err := ts.Do(ctx, r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var tkmap = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&tkmap); err != nil {
		return nil, err
	}
	return oauth2.TokenFromMap(tkmap, 30*time.Second)
}
