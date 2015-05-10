// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package interfaceGUID

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/maruel/ut"
)

func ExampleCalculate() {
	type Foo interface {
		Baz1(int) error
		Baz2()
	}
	fmt.Printf("%s\n", Calculate(reflect.TypeOf((*Foo)(nil)).Elem()))
	// Output:
	// ee2fbefc01ff399e106213e5eab0f3a36c245f195195bd36a2058c115637e05d
}

func TestCalculatePrivateName(t *testing.T) {
	type value1 int
	type value2 int
	ut.AssertEqual(t, "26461de90bccabfde78889464de27c40b1a962fdc312a7b6284dd21593ffcd83", Calculate(reflect.TypeOf((*value1)(nil)).Elem()))
	ut.AssertEqual(t, "26461de90bccabfde78889464de27c40b1a962fdc312a7b6284dd21593ffcd83", Calculate(reflect.TypeOf((*value2)(nil)).Elem()))
}

func TestCalculateFooBase(t *testing.T) {
	// All the TestCalculateFoo* are relative to this test case.
	type Foo interface {
		Baz1(int) error
		Baz2()
	}
	type Bar interface {
		GetFoo(int) Foo
	}
	ut.AssertEqual(t, "18f7adec8708b2fd00336fb6d88911e465618c9429c823ca8f70c36f1794fcf1", Calculate(reflect.TypeOf((*Bar)(nil)).Elem()))
}

func TestCalculateFooMethodType(t *testing.T) {
	// Inner interface method has different argument.
	type Foo interface {
		Baz1(int64) error
		Baz2()
	}
	type Bar interface {
		GetFoo(int) Foo
	}
	ut.AssertEqual(t, "7945944c9adf174cbac6e4b6bac219bb4eedd802818cbdc6477b6b3f7b9f9cb1", Calculate(reflect.TypeOf((*Bar)(nil)).Elem()))
}

func TestCalculateFooMethodOrder(t *testing.T) {
	// Inner interface has same methods but in different order.
	type Foo interface {
		Baz2()
		Baz1(int) error
	}
	type Bar interface {
		GetFoo(int) Foo
	}
	ut.AssertEqual(t, "18f7adec8708b2fd00336fb6d88911e465618c9429c823ca8f70c36f1794fcf1", Calculate(reflect.TypeOf((*Bar)(nil)).Elem()))
}

func TestCalculateFooInnerName(t *testing.T) {
	// Inner interface has different name.
	type Foo2 interface {
		Baz1(int) error
		Baz2()
	}
	type Bar interface {
		GetFoo(int) Foo2
	}
	ut.AssertEqual(t, "ac703259173d15ced210292885a05c6bd0d7d766d97a6fec32fe6e950d81593a", Calculate(reflect.TypeOf((*Bar)(nil)).Elem()))
}

func TestCalculateFooPrivate(t *testing.T) {
	// Inner interface has private method.
	type Foo interface {
		Baz1(int) error
		Baz2()
		privateFn()
	}
	type Bar interface {
		GetFoo(int) Foo
	}
	ut.AssertEqual(t, "2afcbe99c3d2df8d2273723de6046a2922284a22bbe3ec91d2e312ea7b092fe5", Calculate(reflect.TypeOf((*Bar)(nil)).Elem()))
}

func TestCalculateFooPrivateDifferentName(t *testing.T) {
	// Inner interface has private method but with a different name.
	type Foo interface {
		Baz1(int) error
		Baz2()
		privateFn2()
	}
	type Bar interface {
		GetFoo(int) Foo
	}
	ut.AssertEqual(t, "20b6295934efe4f25ead7ccf461e8894d99ee94a2f4e34e4200010c67aa2d621", Calculate(reflect.TypeOf((*Bar)(nil)).Elem()))
}

func TestCalculateFooReturn(t *testing.T) {
	// Inner interface method has different return type.
	type Foo interface {
		Baz1(int) string
		Baz2()
	}
	type Bar interface {
		GetFoo(int) Foo
	}
	ut.AssertEqual(t, "aa65c9ef02f86f231200d0e3ac25f642e18a97730967a8a24203454ae20f73eb", Calculate(reflect.TypeOf((*Bar)(nil)).Elem()))
}

func TestCalculateFooBarRenamed(t *testing.T) {
	// The type has a different name, the hash value doesn't change.
	type Foo interface {
		Baz1(int) error
		Baz2()
	}
	type Bar interface {
		GetFoo(int) Foo
	}
	ut.AssertEqual(t, "18f7adec8708b2fd00336fb6d88911e465618c9429c823ca8f70c36f1794fcf1", Calculate(reflect.TypeOf((*Bar)(nil)).Elem()))
}

func TestCalculateMethodSelfReferencing(t *testing.T) {
	// The type is self-referencing.
	type Foo interface {
		Bar() Foo
	}
	ut.AssertEqual(t, "891e72350ead590aaa4b8ff9b4c829ef6c9af819d5c041fbccfa8f54304e25ca", Calculate(reflect.TypeOf((*Foo)(nil)).Elem()))
}

func TestCalculateTypeSelfReferencing(t *testing.T) {
	// The type is self-referencing.
	type Foo *Foo
	ut.AssertEqual(t, "5a4f72702c70491229d3751c92ea8b403c4aacd7d9df54f79e27c439e8ee8f74", Calculate(reflect.TypeOf((*Foo)(nil)).Elem()))
}

// These types are used in TestCalculateMutuallyReferenced. Go doesn't support
// function-local mutually referencing types so they have to be defined at file
// scope.
type MutuallyReferenced1 interface {
	Foo() MutuallyReferenced2
}
type MutuallyReferenced2 interface {
	Bar() MutuallyReferenced1
}

func TestCalculateMutuallyReferenced(t *testing.T) {
	// The types are mutually referencing.
	ut.AssertEqual(t, "109f2b3b0bcd2050177c557664fca8b7174e2bb6b6b688939f8b15d08cdd5ec1", Calculate(reflect.TypeOf((*MutuallyReferenced1)(nil)).Elem()))
}
