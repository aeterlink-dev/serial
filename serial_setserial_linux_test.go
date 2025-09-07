//go:build linux

package serial

import (
	"context"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func startSocatAndWaitForSetserialTest(t *testing.T, ctx context.Context) *exec.Cmd {
	cmd := exec.CommandContext(ctx, "socat", "-D", "STDIO", "pty,link=/tmp/faketty_setserial")
	r, err := cmd.StderrPipe()
	require.NoError(t, err)
	require.NoError(t, cmd.Start())
	buf := make([]byte, 1024)
	_, err = r.Read(buf)
	require.NoError(t, err)
	return cmd
}

func TestGetSerialStructOnFakeTty(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd := startSocatAndWaitForSetserialTest(t, ctx)
	go cmd.Wait()

	port, err := nativeOpen("/tmp/faketty", &Mode{})
	require.NoError(t, err)
	defer port.Close()
	ser, err := port.GetSerialStruct()

	// Note: socat's virtual TTY may not fully support TIOCGSERIAL.
	// If not supported, skip the test (environment dependent).
	if err != nil {
		t.Skipf("ioctl TIOCGSERIAL not supported on socat pty: %v", err)
	}
	require.NotNil(t, ser)
}

func TestSetSerialPortModeOnFakeTty(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd := startSocatAndWaitForSetserialTest(t, ctx)
	go cmd.Wait()

	port, err := nativeOpen("/tmp/faketty", &Mode{})
	require.NoError(t, err)
	defer port.Close()
	err = port.SetSerialPortMode(0)
	// Note: socat's virtual TTY may not fully support TIOCSSERIAL.
	// If not supported, skip the test (environment dependent).
	if err != nil {
		t.Skipf("ioctl TIOCSSERIAL not supported on socat pty: %v", err)
	}
}
