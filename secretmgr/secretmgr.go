// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package secretmgr is a lightweight implementation of
// Google's secret manager client api
package secretmgr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jfcote87/ctxclient"
)

// SecretMgr handles secret manager api calls
type SecretMgr struct {
	ctxclient.Func
}

// VersionPayload is the response to the secret manager access call
type VersionPayload struct {
	Data  []byte `json:"data"`
	Crc32 string `json:"datCrc32c"`
}

// DecodeVersion decodes the secret value into result
func (sm *SecretMgr) DecodeVersion(ctx context.Context, secretID string, result interface{}) error {
	b, err := sm.Access(ctx, secretID)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, result)
}

// Access returns the secret version value
func (sm *SecretMgr) Access(ctx context.Context, secretID string) ([]byte, error) {
	var payload = struct {
		Name    string          `json:"name"`
		Payload *VersionPayload `json:"payload"`
	}{}
	var accessURL = fmt.Sprintf("https://secretmanager.googleapis.com/v1/{name=%s}:access", secretID)
	r, _ := http.NewRequest("GET", accessURL, nil)
	r.Header.Set("Accept", "application/pdf")
	res, err := sm.Do(ctx, r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return nil, err
	}
	if payload.Payload == nil || len(payload.Payload.Data) == 0 {
		return nil, errors.New("payload data is empty")
	}
	return payload.Payload.Data, nil
}
