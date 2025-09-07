//go:build darwin || freebsd || openbsd

package serial

import "fmt"

func (port *unixPort) SetSerialPortMode(portMode uint32) error {
	return fmt.Errorf("SetSerialPortMode is not supported on this OS")
}
func (port *unixPort) GetSerialStruct() (*LinuxCSerialStruct, error) {
	return nil, fmt.Errorf("GetSerialStruct is not supported on this OS")
}
