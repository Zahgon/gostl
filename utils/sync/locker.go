package sync

import (
	gosync "sync"
)

// Locker define an abstract locker interface
type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

var _ Locker = (*gosync.RWMutex)(nil)

// FakeLocker is a fake locker
type FakeLocker struct {
}

// Lock does nothing
func (l FakeLocker) Lock() {
	_ = "STUB: not implemented"

	// Unlock does nothing
	return
}

func (l FakeLocker) Unlock() {
	_ = "STUB: not implemented"

	// RLock does nothing
	return
}

func (l FakeLocker) RLock() {
	_ = "STUB: not implemented"

	// RUnlock does nothing
	return
}

func (l FakeLocker) RUnlock() { _ = "STUB: not implemented"; return }
