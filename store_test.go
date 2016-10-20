// Copyright 2016 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package bolthold_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/timshannon/bolthold"
)

func TestOpen(t *testing.T) {
	filename := tempfile()
	store, err := bolthold.Open(filename, 0666, nil)
	if err != nil {
		t.Fatalf("Error opening %s: %s", filename, err)
	}

	if store == nil {
		t.Fatalf("store is null!")
	}

	defer store.Close()
	defer os.Remove(filename)
}

func TestBolt(t *testing.T) {
	testWrap(t, func(store *bolthold.Store, t *testing.T) {
		b := store.Bolt()
		if b == nil {
			t.Fatalf("Bolt is null in bolthold")
		}
	})
}

func TestRemoveIndex(t *testing.T) {
	testWrap(t, func(store *bolthold.Store, t *testing.T) {
		//TODO
	})
}

func TestReIndex(t *testing.T) {
	testWrap(t, func(store *bolthold.Store, t *testing.T) {
		//TODO
	})
}

// utilities

// testWrap creates a temporary database for testing and closes and cleans it up when
// completed.
func testWrap(t *testing.T, tests func(store *bolthold.Store, t *testing.T)) {
	filename := tempfile()
	store, err := bolthold.Open(filename, 0666, nil)
	if err != nil {
		t.Fatalf("Error opening %s: %s", filename, err)
	}

	if store == nil {
		t.Fatalf("store is null!")
	}

	defer store.Close()
	defer os.Remove(filename)

	tests(store, t)
}

// tempfile returns a temporary file path.
func tempfile() string {
	f, err := ioutil.TempFile("", "bolthold-")
	if err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
	if err := os.Remove(f.Name()); err != nil {
		panic(err)
	}
	return f.Name()
}
