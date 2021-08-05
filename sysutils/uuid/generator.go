// Copyright (C) 2013-2018 by Maxim Bublis <b@codemonkey.ru>
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"hash"
	"net"
	"os"
	"sync"
	"time"
)

// Difference in 100-nanosecond intervals between
// UUID epoch (October 15, 1582) and Unix epoch (January 1, 1970).
const epochStart = 122192928000000000

var (
	global = newDefaultGenerator()

	epochFunc = unixTimeFunc
	posixUID  = uint32(os.Getuid())
	posixGID  = uint32(os.Getgid())
)

// NewV4 returns random generated UUID.
func NewV4() UUID {
	return global.NewV4()
}

// Generator provides interface for generating UUIDs.
type Generator interface {
	NewV4() UUID
}

// Default generator implementation.
type generator struct {
	storageOnce  sync.Once
	storageMutex sync.Mutex

	lastTime      uint64
	clockSequence uint16
	hardwareAddr  [6]byte
}

func newDefaultGenerator() Generator {
	return &generator{}
}

// NewV4 returns random generated UUID.
func (g *generator) NewV4() UUID {
	u := UUID{}
	g.safeRandom(u[:])
	u.SetVersion(V4)
	u.SetVariant(VariantRFC4122)

	return u
}

func (g *generator) initStorage() {
	g.initClockSequence()
	g.initHardwareAddr()
}

func (g *generator) initClockSequence() {
	buf := make([]byte, 2)
	g.safeRandom(buf)
	g.clockSequence = binary.BigEndian.Uint16(buf)
}

func (g *generator) initHardwareAddr() {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			if len(iface.HardwareAddr) >= 6 {
				copy(g.hardwareAddr[:], iface.HardwareAddr)
				return
			}
		}
	}

	// Initialize hardwareAddr randomly in case
	// of real network interfaces absence
	g.safeRandom(g.hardwareAddr[:])

	// Set multicast bit as recommended in RFC 4122
	g.hardwareAddr[0] |= 0x01
}

func (g *generator) safeRandom(dest []byte) {
	if _, err := rand.Read(dest); err != nil {
		panic(err)
	}
}

// Returns UUID v1/v2 storage state.
// Returns epoch timestamp, clock sequence, and hardware address.
func (g *generator) getStorage() (uint64, uint16, []byte) {
	g.storageOnce.Do(g.initStorage)

	g.storageMutex.Lock()
	defer g.storageMutex.Unlock()

	timeNow := epochFunc()
	// Clock changed backwards since last UUID generation.
	// Should increase clock sequence.
	if timeNow <= g.lastTime {
		g.clockSequence++
	}
	g.lastTime = timeNow

	return timeNow, g.clockSequence, g.hardwareAddr[:]
}

// Returns difference in 100-nanosecond intervals between
// UUID epoch (October 15, 1582) and current time.
// This is default epoch calculation function.
func unixTimeFunc() uint64 {
	return epochStart + uint64(time.Now().UnixNano()/100)
}

// Returns UUID based on hashing of namespace UUID and name.
func newFromHash(h hash.Hash, ns UUID, name string) UUID {
	u := UUID{}
	h.Write(ns[:])
	h.Write([]byte(name))
	copy(u[:], h.Sum(nil))

	return u
}
