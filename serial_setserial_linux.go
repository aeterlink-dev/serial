//go:build linux
// +build linux

package serial

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// CSerialStruct is a C-interop struct for linux/serial.h: struct serial_struct
// Field order and types must exactly match the C definition for binary compatibility.
// See: https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/serial.h#L135
type CSerialStruct struct {
	Type          int32
	Line          int32
	Port          uint32
	IRQ           int32
	Flags         int32
	XmitFifoSize  int32
	CustomDivisor int32
	BaudBase      int32
	CloseDelay    uint16
	IOType        byte
	ReservedChar  byte
	Hub6          int32
	ClosingWait   uint16
	ClosingWait2  uint16
	IOMemBase     uintptr
	IOMemRegShift uint16
	PortHigh      uint32
	// IOMapBase: In C, this is 'unsigned long'. Use uint64 for 64-bit systems.
	// For 32-bit systems, you may need to use uint32 for binary compatibility.
	IOMapBase uint64
}

// GetSerialStruct opens the device and retrieves the CSerialStruct using TIOCGSERIAL ioctl.
func GetSerialStruct(device string) (*CSerialStruct, error) {
	f, err := os.OpenFile(device, os.O_RDWR|syscall.O_NOCTTY|syscall.O_NONBLOCK, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open device: %w", err)
	}
	defer f.Close()

	var ser CSerialStruct
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), unix.TIOCGSERIAL, uintptr(unsafe.Pointer(&ser)))
	if errno != 0 {
		return nil, fmt.Errorf("ioctl TIOCGSERIAL failed: %v", errno)
	}
	return &ser, nil
}

// SetSerialPortMode sets the port mode using ioctl TIOCSSERIAL
func SetSerialPortMode(device string, portMode uint32) error {
	f, err := os.OpenFile(device, os.O_RDWR|syscall.O_NOCTTY|syscall.O_NONBLOCK, 0666)
	if err != nil {
		return fmt.Errorf("failed to open device: %w", err)
	}
	defer f.Close()

	var ser CSerialStruct
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), unix.TIOCGSERIAL, uintptr(unsafe.Pointer(&ser)))
	if errno != 0 {
		return fmt.Errorf("ioctl TIOCGSERIAL failed: %v", errno)
	}

	ser.Port = portMode

	_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), unix.TIOCSSERIAL, uintptr(unsafe.Pointer(&ser)))
	if errno != 0 {
		return fmt.Errorf("ioctl TIOCSSERIAL failed: %v", errno)
	}
	return nil
}
