package cfauth_test

import (
	"context"
	"testing"

	"github.com/jfcote87/oauth2"
	"github.com/lfcloudfunc/utils/cfauth"
)

type tsTest struct {
	tk string
}

func (tx tsTest) Token(ctx context.Context) (*oauth2.Token, error) {
	return &oauth2.Token{AccessToken: tx.tk}, nil
}

func TestNew(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string

		scopes  []string
		ctx     context.Context
		ts      *cfauth.TokenSource
		want    string
		wanterr bool
	}{
		{name: "Test00", ts: &cfauth.TokenSource{}, ctx: ctx, wanterr: true},
		{name: "Test01", ctx: cfauth.ContextTokenSource(ctx, tsTest{tk: "AAAAA"}), want: "AAAAA"},
		{name: "Test02", ctx: ctx, ts: &cfauth.TokenSource{TS: tsTest{tk: "ZZZZZ"}}, want: "ZZZZZ"},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk, err := tt.ts.Token(tt.ctx)
			if err != nil {
				if tt.wanterr {
					return
				}
				t.Errorf("expected %s; got error %v", tt.want, err)
				return
			}
			if tk == nil {
				t.Errorf("expected %s; got nil", tt.want)
				return
			}
			if tt.wanterr {
				t.Errorf("expected error; got %s", tk.AccessToken)
				return
			}
			if tt.want != tk.AccessToken {
				t.Errorf("expected %s; got %s", tt.want, tk.AccessToken)
				return
			}

		})
	}
}
