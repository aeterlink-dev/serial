//go:build darwin || freebsd || openbsd

package serial

import (
	"fmt"
)

func (port *unixPort) EnableRS485(config *LinuxRS485Config) error {
	return fmt.Errorf("EnableRS485 is not supported on this OS")
}
