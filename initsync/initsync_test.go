// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package initsync_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/lfcloudfunc/utils/initsync"
)

type one int

func (o *one) Increment() {
	*o++
}

func run(t *testing.T, once *initsync.Once, o *one, c chan error) {
	c <- once.Do(func() error {
		o.Increment()
		if int(*o) < 3 {
			return fmt.Errorf("ERR%02d", *o)
		}
		return nil
	})

}

func TestOnce(t *testing.T) {
	o := new(one)
	once := new(initsync.Once)
	c := make(chan error)
	const N = 10
	for i := 0; i < N; i++ {
		go run(t, once, o, c)
	}
	cnt := 0
	for i := 0; i < N; i++ {
		if err := <-c; err == nil {
			cnt++
		}
	}
	if *o != 3 {
		t.Errorf("once failed outside run: %d is not 3", *o)
	}
	if cnt != 8 {
		t.Errorf("expected 7 skipped; got %d", cnt)
	}
}

var testErr = errors.New("test")

func TestOncePanic(t *testing.T) {
	var once initsync.Once
	cnt := 0
	f := func() error {
		if cnt == 0 {
			panic("failed")
		}
		return testErr
	}
	func() {
		defer func() {
			if r := recover(); r != nil {
				cnt++
			}
		}()
		_ = once.Do(f)
		t.Error("expected panic, should not be here")
	}()

	if err := once.Do(f); err != testErr {
		t.Errorf("expected testErr; got %v", err)
	}

}

func BenchmarkOnce(b *testing.B) {
	var once initsync.Once
	f := func() error { return nil }
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			once.Do(f)
		}
	})
}
