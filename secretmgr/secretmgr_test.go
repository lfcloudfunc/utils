// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package secretmgr_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/jfcote87/ctxclient"
	"github.com/lfcloudfunc/utils/secretmgr"
)

type testHandler struct {
	u string
}

func (th *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/v1/{name=abc}:access":
		http.Error(w, "No Rec", 404)
		return
	case "/v1/{name=def}:access":
		w.Write([]byte(`{ "name": "def", "payload": {"data": "ImRhdGFzdHJpbmci"} }`))
	case "/v1/{name=ghi}:access":
		w.Write([]byte(`{ name": "ghi", "payload": {"data": "ImRhdGFzdHJpbmci"} }`))
	case "/v1/{name=jkl}:access":
		w.Write([]byte(`{ "name": "jkl" }`))
	case "/v1/{name=xyz}:access":
		var data = "ewoJImludGFjY3QiOiB7CgkgICJzZW5kZXJfaWQiOiAiWFhYWFhYWFgiLAoJICAic2VuZGVyX3B3ZCI6ICIxMjM0NTYiLAoJICAibG9naW4iOiB7CgkJInVzZXJfaWQiOiAicDEiLAoJCSJjb21wYW55IjogIllZWVlZWVlZIiwKCQkicGFzc3dvcmQiOiAiYWJjZGVmZyIsCgkJImxvY2F0aW9uX2lkIjogIkxvYyIKCSAgfSwKCSAgInNlc3Npb24iOiB7CgkJImV4cGlyeURlbHRhIjogMTAKCSAgfQoJfSwKCSJzZiI6IHsKCSAgImhvc3QiOiAiZXhhbXBsZS5teS5zYWxlc2ZvcmNlLmNvbSIsCgkgICJjb25zdW1lcl9rZXkiOiAiWlpaWlpaWlpaWlpaWlpaWlpaWlpaWlpaWlpaWlpaWlpaWlpaWlpaWlpaWiIsCgkgICJ1c2VyX2lkIjogInBlcnNvbkBleGFtcGxlLmNvbSIsCgkgICJrZXkiOiAiLS0tLUJFR0lOIFBSSVZBVEUgS0VZLS0tLS1cbk1JSUV2Z0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktnd2dnU2tBZ0VBQW9JQkFRQ2dIZHAvUGtGb2NVenZcbnJBOWhOZnc1VTc0cXpVd21oc0ZSQnJveE1yVGZxUURrVUx5OHE4SWhlR2ExNHlhbThudXNtdFVXd0V0ZnFSak9cbjBueXVjSTRzKzd4TVRsWnBBVEN3Rlo0SW1tMHpPODVkQnJ4N3VUMFZkWmVkOEU1dkdkUXVIYUdrV2pEVmhnSU1cbk9yT0lUM2U0ODB3aldqN0l3WWpTdm9qalZTVUlxSXo2djZXYVNLcjc4VU4zcXJLbVY3VEpjRVQ3ZWkxWDhlMjRcbjhzc0VlRjhnWGV5RHFDS2s0NGZYQ3VGMHl6VXQ0Q05CYVN5UVUvcVFUNlgrbEJBTDRzWjFMajB3anpWOHBMdUFcbmh5Q0tqaW1YTENJSUpMMDk2aVpzdkNpN3BYRjNVcWtpVklzenpWWXkxaElmZFZXYTdLVFdhc29HMDNvQ3huUmZcbnNJa0wyaTZwQWdNQkFBRUNnZ0VBSSsvL2FoN1poRzBsdFFlYjNoaEZWOEtkaExMV2ZERXpzNUY4ZWUxbEhtd3FcbjJKUHNnTFpXV0xmUzZkRWxqRVFSa1NDaFlqMWZ6WVZCSGE2dHNHTnY3ZFhFb2lYVkREVnByYVAwZC83ZE1xb2lcbk84djJ3eGhGQnd1Qkw4QlBnbTA4Syt1WHE3RllrNXhEMm1YVWhPdzlOV1JhT3ZjVzcyMU0zZXhKQWlsS3FpSDdcbnEzREVNc1haVERCakgxR0tsRG9BdUZvSEpqYzN1blJNS2EvRjAyL3VxMWVDRFp0empQU2hkUXdmdk1IeFVPQWRcbjBOQjMwSWZEUEE0YU52bm12OVRJbjZJNnhYQ0RtQ3hDYWFqeVB1bm9SK0tGaEV2M3k1dFpCelV2a2Q3WExqOEFcbkZPSzFHMzRwSVB0UWhRdVZQeUNWYS9RdzJjZFd6ZEFQOTZxQlErNzhPUUtCZ1FEU1VkYXNsc0hubFFtWWJoQmdcbktPeXg1VXJDS2s3dTErSFpsakt4enZtSDF5Zk50TVI3Y3puUXJsaGdwaFBsZzErSUN1QnJ6V2h4cFlONzROZEZcblRrVnBlTVdncW5ieUtoT0wvN1NSTG9nYTcyRnAxSFdBMVM4R2VEdmd6V2xLV1lKVnFIU2tQekczU2N2VXVHVkFcbjRtbVNkenBiRG1KQ1dQN05uc0tRK3JYd2hRS0JnUURDNUtITU83c2N5V3lYb1BnRVZYZGk0a25wTG5iQVMwL0tcbnFGTTJQbENOelg1MThsSzZ4bGsxRnRTNC9SNjBsYXZaNlQ5L25GRTRXOGtKTGFJOUdMSk0rZHBVTjVUbUxMRERcblc1L2ZrU01iLzc2NHJNbmU2ajgzZWFsMERvZDFNbDFLWGdHWWpGcmVPelExZWtaVSt2WmNvZ0pCa0pDYWhicjhcblBKblN2dW5RMVFLQmdRQ1lnTkVaNm1NRjZsSy8zYWN3SXR5bWplMnNad00vT3ZqQVRxSEp2cHUvZlNxV2hXb25cblhjRGhhVm9hRXJMQXJJYjczVzhubkJVTXV3aGhMZ0xLZFZ2dGc2NUJ0dmx5MVVIQk5SUVlaZFk1T2JTNENWOStcbm9KWnRxWWM1bFppd1djRWhIWG9la3RnWTRJRFZhVE82WGJKUml2U2NqQXlCc1locFRxSWtnRGFuQlFLQmdRQ1BcbkZYSWdVbDRPUldjL2IwT0hVd1FDaWlZL1Z3aUNKd2lldmJQUjZEQVYyNStUa21QZ2tremVLRFBtMkpiWmQrVWlcbmFCZGhaZTlQVkR1bFRHYUhPb0YwK0I2dG85emluRDd1UG0rQ29MakpEcHNFVkg4azlJK1M5L1JMcTNzUnRQYytcbjZRTjg3dGFydjRXV0dYNVhWNlRWMzIrT21VZEEyeHVyb2daQTdqOFZQUUtCZ0hLK2ZWNy9HRnVtUmRXMmcxUDdcblBpdTNzNndqbWcwRDN4YnBJQTZHa3FLWWRaY0RmRVZxWk1xM0dDWGVDQXkrMDh2SVM5TVYwWEpsU1Q5dyswNUFcbkFZSktTWGFNd0ZKMXRFeWZ5dkFKakIwOGhkMGpaZnlxRGkyM1FYWVVYa0lOOTg5VGlOYVpvQitlT2kycmMvMnlcbk5MaExyQ3JzZTlKV1FlS1hjVWw1RlZhZlxuLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLVxuIgoJfSwKCSJkcyI6IHsKCSAgImFjY291bnRfaWQiOiAiMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwKCSAgImFjY2Vzc190b2tlbiI6ICJhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhPSIsCgkgICJzY29wZSI6ICJhcGkiLAoJICAidG9rZW5fdHlwZSI6ICJiZWFyZXIiLAoJICAiaG9zdCI6ICJ3d3cuZG9jdXNpZ24ubmV0IiwKCSAgIm9uQmVoYWxmT2YiOiAiY29uZmVyZW5jZS1pbnZpdGF0aW9uQGV4YW1wbGUuY29tIgoJfQogIH0"
		var result = struct {
			Name    string                    `json:"name"`
			Payload *secretmgr.VersionPayload `json:"payload"`
		}{
			Name:    "xyz",
			Payload: &secretmgr.VersionPayload{Data: []byte(data)},
		}
		json.NewEncoder(w).Encode(result)
	}
	return
}

func (th *testHandler) RoundTrip(r *http.Request) (*http.Response, error) {
	ux, _ := url.Parse(th.u + r.URL.Path)
	rx := *r
	rx.URL = ux
	return http.DefaultTransport.RoundTrip(&rx)
}

func TestSecretMgr_DecodeVersion(t *testing.T) {
	var th = &testHandler{}
	srv := httptest.NewServer(th)
	defer srv.Close()
	th.u = srv.URL
	sm := &secretmgr.SecretMgr{
		Func: func(ctx context.Context) (*http.Client, error) {
			return &http.Client{Transport: th}, nil
		},
	}
	ctx := context.Background()
	var ix interface{}
	err := sm.DecodeVersion(ctx, "abc", &ix)
	var ex *ctxclient.NotSuccess
	if !errors.As(err, &ex) || ex.StatusCode != 404 {
		t.Errorf("expected 404 error; got %v", err)
		return
	}
	var datastring string
	err = sm.DecodeVersion(ctx, "def", &datastring)
	if err != nil {
		t.Errorf("expected success; got %v", err)
		return
	}
	if datastring != "datastring" {
		t.Errorf("expected datastring; got %s", datastring)
		return
	}
	var expectedJsonErr *json.SyntaxError
	err = sm.DecodeVersion(ctx, "ghi", nil)
	if !errors.As(err, &expectedJsonErr) {
		t.Errorf("expected json.SyntaxError; got %v", err)
		return
	}

	err = sm.DecodeVersion(ctx, "jkl", &datastring)
	if err == nil || err.Error() != "payload data is empty" {
		t.Errorf("payload data is empty; got %v", err)
		return
	}

}

/*
{
	"intacct": {
	  "sender_id": "XXXXXXXX",
	  "sender_pwd": "123456",
	  "login": {
		"user_id": "p1",
		"company": "YYYYYYYY",
		"password": "abcdefg",
		"location_id": "Loc"
	  },
	  "session": {
		"expiryDelta": 10
	  }
	},
	"sf": {
	  "host": "example.my.salesforce.com",
	  "consumer_key": "ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ",
	  "user_id": "person@example.com",
	  "key": "----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCgHdp/PkFocUzv\nrA9hNfw5U74qzUwmhsFRBroxMrTfqQDkULy8q8IheGa14yam8nusmtUWwEtfqRjO\n0nyucI4s+7xMTlZpATCwFZ4Imm0zO85dBrx7uT0VdZed8E5vGdQuHaGkWjDVhgIM\nOrOIT3e480wjWj7IwYjSvojjVSUIqIz6v6WaSKr78UN3qrKmV7TJcET7ei1X8e24\n8ssEeF8gXeyDqCKk44fXCuF0yzUt4CNBaSyQU/qQT6X+lBAL4sZ1Lj0wjzV8pLuA\nhyCKjimXLCIIJL096iZsvCi7pXF3UqkiVIszzVYy1hIfdVWa7KTWasoG03oCxnRf\nsIkL2i6pAgMBAAECggEAI+//ah7ZhG0ltQeb3hhFV8KdhLLWfDEzs5F8ee1lHmwq\n2JPsgLZWWLfS6dEljEQRkSChYj1fzYVBHa6tsGNv7dXEoiXVDDVpraP0d/7dMqoi\nO8v2wxhFBwuBL8BPgm08K+uXq7FYk5xD2mXUhOw9NWRaOvcW721M3exJAilKqiH7\nq3DEMsXZTDBjH1GKlDoAuFoHJjc3unRMKa/F02/uq1eCDZtzjPShdQwfvMHxUOAd\n0NB30IfDPA4aNvnmv9TIn6I6xXCDmCxCaajyPunoR+KFhEv3y5tZBzUvkd7XLj8A\nFOK1G34pIPtQhQuVPyCVa/Qw2cdWzdAP96qBQ+78OQKBgQDSUdaslsHnlQmYbhBg\nKOyx5UrCKk7u1+HZljKxzvmH1yfNtMR7cznQrlhgphPlg1+ICuBrzWhxpYN74NdF\nTkVpeMWgqnbyKhOL/7SRLoga72Fp1HWA1S8GeDvgzWlKWYJVqHSkPzG3ScvUuGVA\n4mmSdzpbDmJCWP7NnsKQ+rXwhQKBgQDC5KHMO7scyWyXoPgEVXdi4knpLnbAS0/K\nqFM2PlCNzX518lK6xlk1FtS4/R60lavZ6T9/nFE4W8kJLaI9GLJM+dpUN5TmLLDD\nW5/fkSMb/764rMne6j83eal0Dod1Ml1KXgGYjFreOzQ1ekZU+vZcogJBkJCahbr8\nPJnSvunQ1QKBgQCYgNEZ6mMF6lK/3acwItymje2sZwM/OvjATqHJvpu/fSqWhWon\nXcDhaVoaErLArIb73W8nnBUMuwhhLgLKdVvtg65Btvly1UHBNRQYZdY5ObS4CV9+\noJZtqYc5lZiwWcEhHXoektgY4IDVaTO6XbJRivScjAyBsYhpTqIkgDanBQKBgQCP\nFXIgUl4ORWc/b0OHUwQCiiY/VwiCJwievbPR6DAV25+TkmPgkkzeKDPm2JbZd+Ui\naBdhZe9PVDulTGaHOoF0+B6to9zinD7uPm+CoLjJDpsEVH8k9I+S9/RLq3sRtPc+\n6QN87tarv4WWGX5XV6TV32+OmUdA2xurogZA7j8VPQKBgHK+fV7/GFumRdW2g1P7\nPiu3s6wjmg0D3xbpIA6GkqKYdZcDfEVqZMq3GCXeCAy+08vIS9MV0XJlST9w+05A\nAYJKSXaMwFJ1tEyfyvAJjB08hd0jZfyqDi23QXYUXkIN989TiNaZoB+eOi2rc/2y\nNLhLrCrse9JWQeKXcUl5FVaf\n-----END PRIVATE KEY-----\n"
	},
	"ds": {
	  "account_id": "00000000-0000-0000-0000-000000000000",
	  "access_token": "aaaaaaaaaaaaaaaaaaaaaaaaaaaa=",
	  "scope": "api",
	  "token_type": "bearer",
	  "host": "www.docusign.net",
	  "onBehalfOf": "conference-invitation@example.com"
	}
  }
*/
