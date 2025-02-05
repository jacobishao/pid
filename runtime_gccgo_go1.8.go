// This file may have been modified by CloudWeGo authors. (“CloudWeGo Modifications”). All CloudWeGo Modifications are Copyright 2021 CloudWeGo authors.

//go:build gccgo && go1.8
// +build gccgo,go1.8

package goid

// https://github.com/gcc-mirror/gcc/blob/gcc-7-branch/libgo/go/runtime/runtime2.go#L329-L422

type g struct {
	_panic       uintptr
	_defer       uintptr
	m            uintptr
	syscallsp    uintptr
	syscallpc    uintptr
	param        uintptr
	atomicstatus uint32
	goid         int64 // Here it is!
}
