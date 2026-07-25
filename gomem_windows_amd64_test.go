package gomem

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unsafe"
)

func TestGetProcessFromName(t *testing.T) {
	name := executableName()

	process, err := GetProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	if process.ID == 0 {
		t.Errorf("unexpected process id")
	}

	if process.Name != name {
		t.Errorf("unexpected process name")
	}
}

func TestProcessOpen(t *testing.T) {
	name := executableName()

	process, _ := GetProcessFromName(name)

	handle, err := process.Open()

	if err != nil {
		t.Errorf(err.Error())
	}

	if handle == 0 {
		t.Errorf("unexpected handle id")
	}
}

func TestProcessReadByte(t *testing.T) {
	name := executableName()

	var value = (byte)(0x42)
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadByteAt(valuePtr)

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != assertValue {
		t.Errorf("unexpected value")
	}
}

func TestProcessReadUInt32(t *testing.T) {
	name := executableName()

	var value = (uint32)(42)
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadUInt32(valuePtr)

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != assertValue {
		t.Errorf("unexpected value")
	}
}

func TestProcessReadUInt64(t *testing.T) {
	name := executableName()

	var value = (uint64)(42)
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadUInt64(valuePtr)

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != assertValue {
		t.Errorf("unexpected value")
	}
}

func TestProcessReadFloat32(t *testing.T) {
	name := executableName()

	var value = (float32)(42.0)
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadFloat32(valuePtr)

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != assertValue {
		t.Errorf("unexpected value")
	}
}

func TestProcessReadFloat64(t *testing.T) {
	name := executableName()

	var value = (float64)(42.0)
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadFloat64(valuePtr)

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != assertValue {
		t.Errorf("unexpected value")
	}
}

func TestProcessReadString16(t *testing.T) {
	name := executableName()

	var value = [16]byte{1, 2}
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadString16(valuePtr)
	fmt.Println(assertValue)

	if err != nil {
		t.Errorf(err.Error())
	}

	if string(value[:]) != assertValue {
		t.Errorf("unexpected value")
	}
}

func TestProcessReadBytes(t *testing.T) {
	name := executableName()

	value := []byte{0x11, 0x22, 0x33, 0x44}
	valuePtr := (uintptr)(unsafe.Pointer(&value[0]))

	process, err := GetOpenProcessFromName(name)
	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadBytes(valuePtr, uintptr(len(value)))
	if err != nil {
		t.Errorf(err.Error())
	}

	for i := range value {
		if value[i] != assertValue[i] {
			t.Errorf("unexpected value")
		}
	}
}

func TestProcessReadCString(t *testing.T) {
	name := executableName()

	value := [16]byte{'g', 'o', 'm', 'e', 'm', 0, 'x'}
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetOpenProcessFromName(name)
	if err != nil {
		t.Errorf(err.Error())
	}

	assertValue, err := process.ReadCString(valuePtr, uintptr(len(value)))
	if err != nil {
		t.Errorf(err.Error())
	}

	if assertValue != "gomem" {
		t.Errorf("unexpected value")
	}
}

func TestProcessWriteByte(t *testing.T) {
	name := executableName()

	var (
		value    = (byte)(0x42)
		valuePtr = (uintptr)(unsafe.Pointer(&value))
		newValue = (byte)(0x43)
	)

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	err = process.WriteByteAt(valuePtr, newValue)

	if err != nil {
		t.Errorf(err.Error())
	}

	if value != newValue {
		t.Errorf("unexpected value")
	}
}

func TestGetModuleNotFound(t *testing.T) {
	name := executableName()

	process, err := GetOpenProcessFromName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	ptr, err := process.GetModule("unknown.dll")

	if err.Error() != "module not found" {
		t.Errorf(err.Error())
	}

	if ptr != 0 {
		t.Errorf("unexpected value")
	}
}

func TestProcessClose(t *testing.T) {
	name := executableName()

	process, err := GetOpenProcessFromName(name)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !process.IsOpen() {
		t.Errorf("unexpected value")
	}

	if err := process.Close(); err != nil {
		t.Errorf(err.Error())
	}

	if process.IsOpen() {
		t.Errorf("unexpected value")
	}
}

func TestProcessModules(t *testing.T) {
	name := executableName()

	process, err := GetOpenProcessFromName(name)
	if err != nil {
		t.Errorf(err.Error())
	}

	modules, err := process.Modules()
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(modules) == 0 {
		t.Errorf("unexpected value")
	}

	found := false
	for _, module := range modules {
		if strings.EqualFold(module.Name, name) {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("unexpected value")
	}
}

func TestResolvePointerChain(t *testing.T) {
	name := executableName()

	var target uintptr = 0x42424242
	level1 := uintptr(unsafe.Pointer(&target))
	level2 := uintptr(unsafe.Pointer(&level1))

	process, err := GetOpenProcessFromName(name)
	if err != nil {
		t.Errorf(err.Error())
	}

	resolved, err := process.ResolvePointer(level2, 0, 0)
	if err != nil {
		t.Errorf(err.Error())
	}

	if resolved != target {
		t.Errorf("unexpected value")
	}

	resolved, err = process.ReadPointerChain(level2, 0, 0)
	if err != nil {
		t.Errorf(err.Error())
	}

	if resolved != target {
		t.Errorf("unexpected value")
	}
}

func TestIsKeyDown(t *testing.T) {
	value := IsKeyDown(0x20) // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes

	if value != false {
		t.Errorf("unexpected value")
	}
}

func executableName() string {
	path, _ := os.Executable()

	return filepath.Base(path)
}
