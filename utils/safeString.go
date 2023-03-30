package utils

import "sync"

// SafeString is a thread safe string.
type SafeString struct {
	sync.RWMutex
	s string
}

// NewSafeString creates a new SafeString.
func NewSafeString(s string) *SafeString {
	return &SafeString{s: s}
}

// Get returns the string.
func (ss *SafeString) Get() string {
	ss.RLock()
	defer ss.RUnlock()
	return ss.s
}

// Set sets the string.
func (ss *SafeString) Set(s string) {
	ss.Lock()
	ss.s = s
	ss.Unlock()
}
