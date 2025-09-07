//go:build linux || darwin || freebsd || openbsd

package serial

type UnixPort interface {
	Port

	GetSerialStruct() (*LinuxCSerialStruct, error)
	SetSerialPortMode(portMode uint32) error
}

// UnixOpen opens the serial port using the specified modes
func UnixOpen(portName string, mode *Mode) (UnixPort, error) {
	port, err := nativeOpen(portName, mode)
	if err != nil {
		// Return a nil interface, for which var==nil is true (instead of
		// a nil pointer to a struct that satisfies the interface).
		return nil, err
	}
	return port, err
}
