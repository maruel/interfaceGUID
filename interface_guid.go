// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Package interface_guid exposes one function to calculate a unique identifier
// for an interface.
package interface_guid

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"reflect"
)

type set map[string]bool

func write(h io.Writer, item string) {
	h.Write([]byte(item))
	h.Write([]byte{0})
}

func recurseType(h io.Writer, t reflect.Type, seen set) {
	kind := t.Kind()
	write(h, kind.String())
	if kind == reflect.Interface {
		name := t.Name()
		write(h, name)
		if seen[name] {
			return
		}
		seen[name] = true
		for i := 0; i < t.NumMethod(); i++ {
			recurseMethod(h, t.Method(i), seen)
		}
	} else if kind == reflect.Struct {
		name := t.Name()
		write(h, name)
		if seen[name] {
			return
		}
		seen[name] = true
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			write(h, f.Name)
			recurseType(h, f.Type, seen)
		}
		for i := 0; i < t.NumMethod(); i++ {
			recurseMethod(h, t.Method(i), seen)
		}
	} else if kind == reflect.Array || kind == reflect.Chan || kind == reflect.Ptr || kind == reflect.Slice {
		recurseType(h, t.Elem(), seen)
	} else if kind == reflect.Map {
		recurseType(h, t.Key(), seen)
		recurseType(h, t.Elem(), seen)
	} else if kind >= reflect.Bool && kind <= reflect.Complex128 || kind == reflect.String {
		// Base types.
	} else {
		panic(kind.String())
	}
}

func recurseMethod(h io.Writer, m reflect.Method, seen set) {
	write(h, m.Name)
	for i := 0; i < m.Type.NumIn(); i++ {
		recurseType(h, m.Type.In(i), seen)
	}
	for i := 0; i < m.Type.NumOut(); i++ {
		recurseType(h, m.Type.Out(i), seen)
	}
}

// CalculateGUID returns the hex encoded string of the SHA-256 hash of a
// reflected type, normally an interface.
//
// The reflected type is traversed recursively up to all native types
// referenced. The value is dependent on the referenced type names, methods and
// their order. The name of the type itself is not taken in account.
//
// The purpose of this function is to get into a quick common agreement between
// two remote parties, so that follow up communication can be done with gob or
// another communication mechanism.
//
// See test cases for more details.
func CalculateGUID(t reflect.Type) string {
	h := sha256.New()
	recurseType(h, t, make(set))
	return hex.EncodeToString(h.Sum(nil))
}
