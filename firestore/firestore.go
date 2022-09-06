// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package firestore provides lightweight document read and write
// funcs to the firestore database
package firestore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jfcote87/ctxclient"
)

// ArrayValue holds an array in a firestore doc
type ArrayValue struct {
	Values []Value `json:"values,omitempty"`
}

// StringValue return Value for string
func StringValue(s string) Value {
	var v Value
	v.String = s
	return v
}

// IntegerValue handle ints
func IntegerValue(i int64) Value {
	var v Value
	v.Int = i
	return v
}

// DoubleValue handles floats
func DoubleValue(f float64) Value {
	var v Value
	v.Double = f
	return v
}

// TimestampValue handles floats
func TimestampValue(t time.Time) Value {
	var v Value
	v.Timestamp = &t
	return v

}

// TimestampPtrValue handle pointers with null handling
func TimestampPtrValue(t *time.Time) Value {
	var v Value
	if t != nil {
		return TimestampValue(*t)
	}
	v.Null = true
	return v
}

// BoolValue handles bools
func BoolValue(b bool) Value {
	var v Value
	v.Bool = &b
	return v
}

// BytesValue handles byte slices
func BytesValue(b []byte) Value {
	if b == nil {
		return NullValue()
	}
	var v Value
	v.Bytes = b
	return v
}

// RefValue handle ref type (check if nulls allowed)
func RefValue(ref string) Value {
	var v Value
	v.Ref = ref
	return v
}

// GeoValue handles geo points
func GeoValue(g *FsGeo) Value {
	if g == nil {
		return NullValue()
	}
	var v Value
	v.Geo = g
	return v
}

// ArrayValues handles arrays
func ArrayValues(a []Value) Value {
	if a == nil {
		return NullValue()
	}
	var v Value
	v.Array = &ArrayValue{
		Values: a,
	}
	return v
}

// MapValue handles Maps
func MapValue(a map[string]Value) Value {
	if a == nil {
		return NullValue()
	}
	var v Value
	v.Map = &Fields{a}
	return v
}

// NullValue returns null
func NullValue() Value {
	var v Value
	v.Null = true
	return v
}

// Value represents a field value in firestore database.  One and
// only one field may be set.
type Value struct {
	Null      FsNull      `json:"nullValue,omitempty"`
	String    string      `json:"stringValue,omitempty"`
	Timestamp *time.Time  `json:"timestampValue,omitempty"`
	Bool      *bool       `json:"booleanValue,omitempty"`
	Int       int64       `json:"integerValue,omitempty"`
	Double    float64     `json:"doubleValue,omitempty"`
	Bytes     []byte      `json:"bytesValue,omitempty"`
	Ref       string      `json:"referenceValue,omitempty"`
	Geo       *FsGeo      `json:"geoValue,omitempty"`
	Array     *ArrayValue `json:"arrayValue,omitempty"`
	Map       *Fields     `json:"mapValue,omitempty"`
}

// FsNull represents an explicit null value in firestore db
type FsNull bool

// MarshalJSON outputs null
func (n FsNull) MarshalJSON() ([]byte, error) {
	if n {
		return []byte("null"), nil
	}
	return nil, nil
}

// FsGeo handle locations in firestore
type FsGeo struct {
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
}

// Fields contains map for mapvalue
type Fields struct {
	Fields map[string]Value `json:"fields,omitempty"`
}

// UpdateDoc upserts a record
func UpdateDoc(ctx context.Context, f ctxclient.Func, projectID string, docPath string, doc interface{}) error {
	var body = struct {
		Fields interface{} `json:"fields,omitempty"`
	}{
		Fields: doc,
	}
	b, _ := json.Marshal(body)

	var url = fmt.Sprintf("https://firestore.googleapis.com/v1/projects/%s/databases/(default)/documents/%s", projectID, docPath)

	req, _ := http.NewRequest("PATCH", url, bytes.NewReader(b))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := f.Do(ctx, req)
	if err != nil {
		return err
	}
	res.Body.Close()
	return nil
}

// Document is a return document
type Document struct {
	Name       string
	Fields     map[string]Value
	CreateTime time.Time
	UpdateTime time.Time
}

// GetString returns a string field from Document
func (d *Document) GetString(key string) string {
	val, ok := d.Fields[key]
	if !ok {
		return ""
	}
	return val.String
}

// GetDoc returns a single firestore document
func GetDoc(ctx context.Context, f ctxclient.Func, projectID string, docPath string) (*Document, error) {
	var doc *Document
	var url = fmt.Sprintf("https://firestore.googleapis.com/v1/projects/%s/databases/(default)/documents/%s", projectID, docPath)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	res, err := f.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	return doc, json.NewDecoder(res.Body).Decode(&doc)
}
