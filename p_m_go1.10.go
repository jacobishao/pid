// Copyright 2022 Cholerae Hu.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.

// This file may have been modified by CloudWeGo authors. (“CloudWeGo Modifications”). All CloudWeGo Modifications are Copyright 2021 CloudWeGo authors.

//go:build gc && go1.10 && !go1.13
// +build gc,go1.10,!go1.13

package goid

type mutex struct {
	// Futex-based impl treats it as uint32 key,
	// while sema-based impl as M* waitm.
	// Used to be a union, but unions break precise GC.
	key uintptr
}

type p struct {
	lock mutex
	id   int32 // Here is pid
}

type m struct {
	g0      uintptr // goroutine with scheduling stack
	morebuf gobuf   // gobuf arg to morestack
	divmod  uint32  // div/mod denominator for arm - known to liblink

	// Fields not known to debuggers.
	procid     uint64       // for debuggers, but offset not hard-coded
	gsignal    uintptr      // signal-handling g
	goSigStack gsignalStack // Go-allocated signal handling stack
	sigmask    sigset       // storage for saved signal mask
	tls        [6]uintptr   // thread-local storage (for x86 extern register)
	mstartfn   func()
	curg       uintptr // current running goroutine
	caughtsig  uintptr // goroutine running during fatal signal
	p          *p      // attached p for executing go code (nil if not executing go code)
}
