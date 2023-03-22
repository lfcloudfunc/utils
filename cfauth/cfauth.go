// Package cfauth provides a wrapper to a cloud function/gce
// tokensource
package cfauth

import (
	"context"
	"errors"
	"net/http"

	"github.com/jfcote87/oauth2"
	"github.com/jfcote87/oauth2/google"
)

// TokenSource returns access tokens for the default Service Account
type TokenSource struct {
	TS oauth2.TokenSource
}

type tsOverride struct{}

var tsOverrideKey = (*tsOverride)(nil)

// New returns an oauth2.TokenSource that can be overridden with
// a context.Context
func New(scopes ...string) oauth2.TokenSource {
	return &TokenSource{
		TS: google.ComputeTokenSource("", scopes...),
	}
}

// ContextTokenSource creates a context with an override tokensource
func ContextTokenSource(ctx context.Context, ts oauth2.TokenSource) context.Context {
	return context.WithValue(ctx, tsOverrideKey, ts)
}

// Token retrieves an access_token from GCE
func (ts *TokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	if tsCtx, ok := ctx.Value(tsOverrideKey).(oauth2.TokenSource); ok {
		return tsCtx.Token(ctx)
	}
	if ts == nil || ts.TS == nil {
		return nil, errors.New("cfauth: nil tokensource")
	}
	return ts.TS.Token(ctx)
}

// Client returns an authorizing client
func (ts *TokenSource) Client() *http.Client {
	return oauth2.Client(ts, nil)
}
