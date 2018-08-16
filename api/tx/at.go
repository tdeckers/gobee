package tx

import (
	"bytes"
)

const atAPIID byte = 0x08

func NewAT(options ...func(interface{})) *AT {
	f := &AT{Cmd: [2]byte{'N', 'I'}}

	if options == nil {
		return f
	}

	for _, option := range options {
		option(f)
	}

	return f
}

// AT transmit frame
type AT struct {
	FrameID   byte
	Cmd       [2]byte
	Parameter []byte
}

func (f *AT) SetFrameID(id byte) {
	f.FrameID = id
}

func (f *AT) SetCommand(cmd [2]byte) {
	copy(f.Cmd[:], cmd[:])
}

func (f *AT) SetParameter(parameter []byte) {
	f.Parameter = make([]byte, len(parameter))
	copy(f.Parameter, parameter)
}

// Bytes turn AT frame into bytes
func (f *AT) Bytes() ([]byte, error) {
	var b bytes.Buffer

	b.WriteByte(atAPIID)
	b.WriteByte(f.FrameID)
	b.Write(f.Cmd[:])

	if f.Parameter != nil && len(f.Parameter) > 0 {
		b.Write(f.Parameter)
	}

	return b.Bytes(), nil
}
