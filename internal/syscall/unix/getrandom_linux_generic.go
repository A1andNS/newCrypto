// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && arm64
// +build linux,arm64

package unix

// Linux getrandom system call number.
// See GetRandom in getrandom_linux.go.
//
// This file is named "generic" because at a certain point Linux
// started standardizing on system call numbers across
// architectures. So far this means only arm64 uses the standard
// numbers.
const randomTrap uintptr = 278
