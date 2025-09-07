//go:build linux || darwin || freebsd || openbsd

package serial

import "time"

// RS485Config holds configuration for RS485 mode on Linux systems (TIOCSRS485).
// Note: This is only supported on Linux systems.
type LinuxRS485Config struct {
	// Delay RTS prior to send
	DelayRtsBeforeSend time.Duration
	// Delay RTS after send
	DelayRtsAfterSend time.Duration
	// Set RTS high during send
	RtsHighDuringSend bool
	// Set RTS high after send
	RtsHighAfterSend bool
	// Rx during Tx
	RxDuringTx bool
}
