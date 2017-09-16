// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux

package ipv4

import (
	"unsafe"

	"github.com/gy-games-libs/golang/x/net/bpf"
	"github.com/gy-games-libs/golang/x/net/internal/socket"
)

func (so *sockOpt) setAttachFilter(c *socket.Conn, f []bpf.RawInstruction) error {
	prog := sockFProg{
		Len:    uint16(len(f)),
		Filter: (*sockFilter)(unsafe.Pointer(&f[0])),
	}
	b := (*[sizeofSockFprog]byte)(unsafe.Pointer(&prog))[:sizeofSockFprog]
	return so.Set(c, b)
}
