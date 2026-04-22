package platform

import "testing"

func TestHasSPUSensor(t *testing.T) {
	if !HasSPUSensor(`some text AppleSPUHIDDevice more text`) {
		t.Fatal("expected HasSPUSensor to detect AppleSPUHIDDevice")
	}

	if HasSPUSensor(`no relevant sensor here`) {
		t.Fatal("expected HasSPUSensor to return false when sensor is missing")
	}
}
