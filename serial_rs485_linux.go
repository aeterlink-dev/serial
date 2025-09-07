//go:build linux

package serial

import (
	"time"

	"golang.org/x/sys/unix"
)

const (
	rs485Enabled      = 1 << 0
	rs485RTSOnSend    = 1 << 1
	rs485RTSAfterSend = 1 << 2
	rs485RXDuringTX   = 1 << 4
	rs485Tiocs        = unix.TIOCSRS485
)

// EnableRS485 enables RS485 functionality of driver via an ioctl if the config says so
func (port *unixPort) EnableRS485(config *LinuxRS485Config) error {
	rs485 := rs485IoctlOpts{
		rs485Enabled,
		int(config.DelayRtsBeforeSend / time.Millisecond),
		int(config.DelayRtsAfterSend / time.Millisecond),
		[5]int{0, 0, 0, 0, 0},
	}

	if config.RtsHighDuringSend {
		rs485.flags |= rs485RTSOnSend
	}
	if config.RtsHighAfterSend {
		rs485.flags |= rs485RTSAfterSend
	}
	if config.RxDuringTx {
		rs485.flags |= rs485RXDuringTX
	}

	return unix.IoctlSetPointerInt(port.handle, rs485Tiocs, rs485.flags)
}
