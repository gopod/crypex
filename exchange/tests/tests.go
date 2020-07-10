package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// DefaultTimeout default tests timeout
var DefaultTimeout = 3 * time.Second

// ReceiveWithTimeout receive data from channel with timeout.
func ReceiveWithTimeout(t *testing.T, ch <-chan interface{}, timeout *time.Duration) {
	if timeout == nil {
		timeout = &DefaultTimeout
	}

	for stay, timeout := true, time.After(*timeout); stay; {
		select {
		case <-timeout:
			stay = false
		case _, ok := <-ch:
			assert.True(t, ok)
		}
	}
}
