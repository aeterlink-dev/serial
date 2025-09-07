//go:build linux || darwin || freebsd || openbsd

package serial

// LinuxCSerialStruct is a C-interop struct for linux/serial.h: struct serial_struct
// Field order and types must exactly match the C definition for binary compatibility.
// See: https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/serial.h#L135
type LinuxCSerialStruct struct {
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
	// IOMapBase: In C, this is 'unsigned long'.
	IOMapBase uint
}
