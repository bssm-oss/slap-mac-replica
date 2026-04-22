package platform

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// SensorStatus summarizes whether the undocumented Apple Silicon SPU sensor
// appears to be present on the current machine.
type SensorStatus struct {
	Present bool
	Summary string
}

// DetectSensor inspects ioreg output for AppleSPUHIDDevice services.
func DetectSensor() (SensorStatus, error) {
	if runtime.GOOS != "darwin" {
		return SensorStatus{
			Present: false,
			Summary: "macOS가 아니라서 SPU 센서를 확인할 수 없습니다.",
		}, nil
	}

	cmd := exec.Command("/usr/sbin/ioreg", "-l", "-w0")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return SensorStatus{}, fmt.Errorf("ioreg -l -w0: %w: %s", err, strings.TrimSpace(string(output)))
	}

	if HasSPUSensor(string(output)) {
		return SensorStatus{
			Present: true,
			Summary: "AppleSPUHIDDevice 가 확인되었습니다.",
		}, nil
	}

	return SensorStatus{
		Present: false,
		Summary: "AppleSPUHIDDevice 가 보이지 않습니다.",
	}, nil
}

// HasSPUSensor is a small pure helper that keeps unit tests independent from ioreg.
func HasSPUSensor(output string) bool {
	return strings.Contains(output, "AppleSPUHIDDevice")
}
