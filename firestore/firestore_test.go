// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package firestore_test

import (
	"testing"
	"time"

	fs "github.com/lfcloudfunc/utils/firestore"
)

type DSConnect struct {
	EnvelopeID string            `xml:"EnvelopeID" firestore:"envelopeId,omitempty"`
	Subject    string            `xml:"Subject" firestore:"subject,omitempty"`
	UserName   string            `xml:"UserName" firestore:"userName,omitempty"`
	Email      string            `xml:"Email" firestore:"email,omitempty"`
	Status     string            `xml:"Status" firestore:"status,omitempty"`
	Created    *time.Time        `xml:"Created" firestore:"created"`
	Sent       *time.Time        `xml:"Sent" firestore:"sent"`
	Delivered  *time.Time        `xml:"Delivered" firestore:"delivered"`
	Signed     *time.Time        `xml:"Signed" firestore:"signed"`
	Completed  *time.Time        `xml:"Completed" firestore:"completed"`
	VoidReason string            `xml:"VoidReason" firestore:"voidReason"`
	LibertyID  string            `firestore:"libertyId"`
	DriveID    string            `firestore:"driveId"`
	Headers    map[string]string `firestore:"hdrMap"`
}

// MarshalJSON format json for firestore rest api
func (ds DSConnect) valueMap() map[string]fs.Value {
	hdrMap := make(map[string]fs.Value)
	for k, v := range ds.Headers {
		hdrMap[k] = fs.StringValue(v)
	}
	return map[string]fs.Value{
		"envelopeId": fs.StringValue(ds.EnvelopeID),
		"subject":    fs.StringValue(ds.Subject),
		"userName":   fs.StringValue(ds.UserName),
		"email":      fs.StringValue(ds.UserName),
		"status":     fs.StringValue(ds.Status),
		"created":    fs.TimestampPtrValue(ds.Created),
		"sent":       fs.TimestampPtrValue(ds.Sent),
		"delivered":  fs.TimestampPtrValue(ds.Delivered),
		"completed":  fs.TimestampPtrValue(ds.Completed),
		"voidReason": fs.StringValue(ds.VoidReason),
		"libertyId":  fs.StringValue(ds.LibertyID),
		"driveId":    fs.StringValue(ds.DriveID),
		"headers":    fs.MapValue(hdrMap),
	}
}

func TestMarshal(t *testing.T) {
	/*	var blank *time.Time
		var x = make(map[string]fs.Value)
		x["a"] = fs.StringValue("abc")
		x["b"] = fs.IntegerValue(6)
		x["c"] = fs.DoubleValue(8.8)
		x["d"] = fs.NullValue()
		x["e"] = fs.TimestampValue(time.Now())
		x["f"] = fs.TimestampPtrValue(blank)
		x["g"] = fs.BoolValue(false)
		x["h"] = fs.MapValue(map[string]fs.Value{
			"x1": fs.StringValue("x1"),
			"x2": fs.StringValue("x2"),
		})
	*/
	/*
		tm := time.Now()

		dsx := DSConnect{
			EnvelopeID: "envXXX",
			Subject:    "subc",
			UserName:   "894992240294",
			Email:      "jfc@lf.org",
			Status:     "New",
			Created:    &tm,
			LibertyID:  "EXX-XXXX",
			DriveID:    "asdfasdfasdfadfs",
			Headers: map[string]string{
				"Content-type": "app/json",
				"Accept":       "app/json",
			},
		}

		var testToken = "ya29.c.EmFtB6ivmwIBKGy_DdGb8M-9JVRHWPciasIpafB3tQ4_kK6Id68SaQ_OE2_iEnX1WL3q1SpbmQK9D0EB9WlYr6rzmUTmVFSEFvld_DQMK0J1NJ3Y_VDzsWBp5Ibl0FwFTvFA"
		ctx := context.Background()
		appTS := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: testToken})
		cl := oauth2.Client(appTS, nil)
		var f ctxclient.Func = func(ctx context.Context) (*http.Client, error) {
			return cl, nil
		}

		err := fs.UpdateDoc(ctx, f, "lf-domain-project", "testrecs/"+dsx.EnvelopeID, dsx.valueMap())
		t.Errorf("%v", err)
	*/
	/*var body = map[string]interface{}{
		"fields": x,
	}

	b, _ := json.Marshal(body)
	t.Errorf("%s", b)

	var url = "https://firestore.googleapis.com/v1/projects/lf-domain-project/databases/(default)/documents/testrecs/56789"
	rx := bytes.NewReader(b)
	_ = rx
	req, _ := http.NewRequest("PATCH", url, bytes.NewReader(b))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Accept", "application/json")
	//req, _ := http.NewRequest("GET", url, nil) // bytes.NewReader(b))
	res, err := f.Do(ctx, req)
	if err != nil {
		t.Fatalf("%v", err)
	}
	bx, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	t.Errorf("%s", bx)
	*/
}
