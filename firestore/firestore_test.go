// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package firestore_test

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	fs "github.com/lfcloudfunc/utils/firestore"
)

func TestMarshal(t *testing.T) {
	tm := time.Now()
	niltime, tm := (*time.Time)(nil), time.Now()
	var fsrecord = fs.Document{
		Name: "Record Name",
		Fields: map[string]fs.Value{
			"a": fs.StringValue("abc"),
			"b": fs.IntegerValue(6),
			"c": fs.DoubleValue(8.8),
			"d": fs.NullValue(),
			"e": fs.TimestampValue(tm),
			"f": fs.TimestampPtrValue(niltime),
			"g": fs.BoolValue(false),
			"h": fs.MapValue(map[string]fs.Value{
				"x1": fs.StringValue("x1"),
				"x2": fs.StringValue("x2"),
			}),
		},
	}

	var result = make(map[string]map[string]interface{})
	b, _ := json.Marshal(fsrecord.Fields)
	if err := json.Unmarshal(b, &result); err != nil {
		t.Errorf("unable to parse json: %v", err)
		return
	}

	xval1 := map[string]interface{}{"stringValue": "x1"}
	xval2 := map[string]interface{}{"stringValue": "x2"}
	ix := map[string]interface{}{
		"fields": map[string]interface{}{
			"x1": xval1,
			"x2": xval2,
		},
	}
	tmval, _ := json.Marshal(tm)
	if len(tmval) > 1 {
		tmval = tmval[1 : len(tmval)-1]
	}
	testResult := map[string]map[string]interface{}{
		"a": {"stringValue": "abc"},
		"b": {"integerValue": float64(6)},
		"c": {"doubleValue": 8.8},
		"d": {"nullValue": nil},
		"e": {"timestampValue": string(tmval)},
		"f": {"nullValue": nil},
		"g": {"booleanValue": false},
		"h": {"mapValue": ix},
	}

	for k, v := range result {
		tr, ok := testResult[k]
		if !ok {
			t.Errorf("unexpected key %s", k)
			continue
		}
		if !reflect.DeepEqual(tr, v) {
			t.Errorf("expected value for %s = %#v; got %#v", k, tr, v)
		}
	}

}
