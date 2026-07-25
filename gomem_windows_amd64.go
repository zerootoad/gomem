package gomem

import (
	"bytes"
	"errors"
	"unsafe"

	"github.com/zerootoad/gomem/internal/kernel32"
	"github.com/zerootoad/gomem/internal/user32"
)

// Process is a struct representing a windows process.
type Process struct {
	ID     uint32
	Name   string
	Handle uintptr
}

// Module describes a loaded module in a process.
type Module struct {
	Name        string
	Path        string
	BaseAddress uintptr
	Size        uint32
	Handle      uintptr
}

// IsOpen reports whether the process currently has an open handle.
func (p *Process) IsOpen() bool {
	return p != nil && p.Handle != 0
}

// GetProcessFromName converts a process name to a Process struct.
func GetProcessFromName(name string) (*Process, error) {
	pid, err := kernel32.GetProcessID(name)

	if err != nil {
		return nil, err
	}

	process := Process{ID: pid, Name: name}

	return &process, nil
}

// GetOpenProcessFromName converts a process name to a Process struct with open handle.
func GetOpenProcessFromName(name string) (*Process, error) {
	process, err := GetProcessFromName(name)

	if err != nil {
		return nil, err
	}

	_, err = process.Open()

	if err != nil {
		return nil, err
	}

	return process, nil
}

// GetOpenProcessFromPID converts a process pid to a Process struct with open handle.
func GetOpenProcessFromPID(pid uint32) (*Process, error) {
	process := Process{ID: pid, Name: ""}

	_, err := process.Open()
	if err != nil {
		return nil, err
	}

	return &process, nil
}

// Open process handle.
func (p *Process) Open() (uintptr, error) {
	handle, err := kernel32.OpenProcess(kernel32.PROCESS_ALL_ACCESS, false, p.ID)

	if err != nil {
		return 0, err
	}

	p.Handle = handle

	return handle, err
}

// Close process handle.
func (p *Process) Close() error {
	if p == nil || p.Handle == 0 {
		return nil
	}

	if !kernel32.CloseHandle(p.Handle) {
		return errors.New("failed to close process handle")
	}

	p.Handle = 0

	return nil
}

// Read process memory.
func (p *Process) Read(offset uintptr, buffer uintptr, length uintptr) error {
	_, err := kernel32.ReadProcessMemory(p.Handle, offset, buffer, length)

	return err
}

// Read bytes from process memory.
func (p *Process) ReadBytes(offset uintptr, length uintptr) ([]byte, error) {
	buffer := make([]byte, int(length))
	if length == 0 {
		return buffer, nil
	}

	bufferPtr := uintptr(unsafe.Pointer(&buffer[0]))
	err := p.Read(offset, bufferPtr, length)

	return buffer, err
}

// Read bool from process memory.
func (p *Process) ReadBool(offset uintptr) (bool, error) {
	value, err := p.ReadByteAt(offset)
	return value != 0, err
}

// Read byte at an address in process memory.
func (p *Process) ReadByteAt(offset uintptr) (byte, error) {
	var (
		value    byte
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	err := p.Read(offset, valuePtr, unsafe.Sizeof(value))

	return value, err
}

// Read uint32 from process memory.
func (p *Process) ReadUInt32(offset uintptr) (uint32, error) {
	var (
		value    uint32
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	err := p.Read(offset, valuePtr, unsafe.Sizeof(value))

	return value, err
}

// Read int32 from process memory.
func (p *Process) ReadInt32(offset uintptr) (int32, error) {
	var (
		value    int32
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	err := p.Read(offset, valuePtr, unsafe.Sizeof(value))

	return value, err
}

// Read uint64 from process memory.
func (p *Process) ReadUInt64(offset uintptr) (uint64, error) {
	var (
		value    uint64
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	err := p.Read(offset, valuePtr, unsafe.Sizeof(value))

	return value, err
}

// Read int64 from process memory.
func (p *Process) ReadInt64(offset uintptr) (int64, error) {
	var (
		value    int64
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	err := p.Read(offset, valuePtr, unsafe.Sizeof(value))

	return value, err
}

// Read pointer-sized value from process memory.
func (p *Process) ReadPointer(offset uintptr) (uintptr, error) {
	var (
		value    uintptr
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	err := p.Read(offset, valuePtr, unsafe.Sizeof(value))

	return value, err
}

// Read float32 from process memory.
func (p *Process) ReadFloat32(offset uintptr) (float32, error) {
	var (
		value    float32
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	err := p.Read(offset, valuePtr, unsafe.Sizeof(value))

	return value, err
}

// Read float64 from process memory.
func (p *Process) ReadFloat64(offset uintptr) (float64, error) {
	var (
		value    float64
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	err := p.Read(offset, valuePtr, unsafe.Sizeof(value))

	return value, err
}

// Read c-string from process memory up to maxLength bytes.
func (p *Process) ReadCString(offset uintptr, maxLength uintptr) (string, error) {
	buffer, err := p.ReadBytes(offset, maxLength)
	if err != nil {
		return "", err
	}

	if index := bytes.IndexByte(buffer, 0); index >= 0 {
		return string(buffer[:index]), nil
	}

	return string(buffer), nil
}

// Read string16 from process memory.
func (p *Process) ReadString16(offset uintptr) (string, error) {
	var (
		value    [16]byte
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	err := p.Read(offset, valuePtr, unsafe.Sizeof(value))

	return string(value[:]), err
}

// Write process memory.
func (p *Process) Write(offset uintptr, buffer uintptr, length uintptr) error {
	_, err := kernel32.WriteProcessMemory(p.Handle, offset, buffer, length)

	return err
}

// Write bytes to process memory.
func (p *Process) WriteBytes(offset uintptr, buffer []byte) error {
	if len(buffer) == 0 {
		return nil
	}

	bufferPtr := uintptr(unsafe.Pointer(&buffer[0]))
	return p.Write(offset, bufferPtr, uintptr(len(buffer)))
}

// Write bool to process memory.
func (p *Process) WriteBool(offset uintptr, value bool) error {
	var byteValue byte
	if value {
		byteValue = 1
	}

	return p.WriteByteAt(offset, byteValue)
}

// Write byte at an address in process memory.
func (p *Process) WriteByteAt(offset uintptr, value byte) error {
	var (
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	return p.Write(offset, valuePtr, unsafe.Sizeof(value))
}

// Write pointer-sized value to process memory.
func (p *Process) WritePointer(offset uintptr, value uintptr) error {
	var (
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	return p.Write(offset, valuePtr, unsafe.Sizeof(value))
}

// GetModule address.
func (p *Process) GetModule(name string) (uintptr, error) {
	ptr, err := kernel32.GetModule(name, p.ID)

	return ptr, err
}

// Modules returns the loaded modules for the process.
func (p *Process) Modules() ([]Module, error) {
	rawModules, err := kernel32.GetModules(p.ID)
	if err != nil {
		return nil, err
	}

	modules := make([]Module, 0, len(rawModules))
	for _, raw := range rawModules {
		modules = append(modules, Module{
			Name:        raw.Name,
			Path:        raw.Path,
			BaseAddress: raw.BaseAddress,
			Size:        raw.Size,
			Handle:      raw.Handle,
		})
	}

	return modules, nil
}

// FindModule returns a loaded module by name.
func (p *Process) FindModule(name string) (*Module, error) {
	modules, err := p.Modules()
	if err != nil {
		return nil, err
	}

	for i := range modules {
		if modules[i].Name == name {
			return &modules[i], nil
		}
	}

	return nil, errors.New("module not found")
}

// ResolvePointer follows a pointer chain and returns the final value.
func (p *Process) ResolvePointer(base uintptr, offsets ...uintptr) (uintptr, error) {
	address := base

	for _, offset := range offsets {
		if address == 0 {
			return 0, errors.New("null pointer in chain")
		}

		nextAddress, err := p.ReadPointer(address + offset)
		if err != nil {
			return 0, err
		}

		address = nextAddress
	}

	return address, nil
}

// ReadPointerChain is an alias for ResolvePointer for readability at call sites.
func (p *Process) ReadPointerChain(base uintptr, offsets ...uintptr) (uintptr, error) {
	return p.ResolvePointer(base, offsets...)
}

// IsKeyDown https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
func IsKeyDown(v int) bool {
	return user32.GetAsyncKeyState(v) != 0
}
