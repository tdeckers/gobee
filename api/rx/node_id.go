package rx

import (
	"encoding/binary"
)

const (
	nodeIDAPIID byte = 0x95

	nodeIDAddr64Offset       = 0
	nodeIDAddr16Offset       = 8
	nodeIDOptionsOffset      = 10
	nodeIDRemoteAddr64Offset = 13
	nodeIDStringOffset       = 21
)

type NodeID struct {
	buffer                 []byte
	nodeIdDeviceTypeOffset int
}

func newNodeID() Frame {
	return &NodeID{
		buffer:                 make([]byte, 0),
		nodeIdDeviceTypeOffset: 0,
	}
}

// RX receive data frame byte.  Function is called for every byte received.
func (f *NodeID) RX(b byte) error {
	f.buffer = append(f.buffer, b)

	// Find nodeIdDeviceTypeOffset.  Search at position 21 and search for 0x00 (which
	// terminates the NI field)
	if f.nodeIdDeviceTypeOffset == 0 && len(f.buffer) > nodeIDStringOffset && f.buffer[len(f.buffer)-1] == 0x00 {
		f.nodeIdDeviceTypeOffset = len(f.buffer) + 2 // Skip two reserved bytes after NI string.
	}

	return nil
}

func (f *NodeID) Addr64() uint64 {
	return binary.BigEndian.Uint64(f.buffer[nodeIDAddr64Offset : nodeIDAddr64Offset+addr64Length])
}

func (f *NodeID) Addr16() uint16 {
	return binary.BigEndian.Uint16(f.buffer[nodeIDAddr16Offset : nodeIDAddr16Offset+addr16Length])
}

func (f *NodeID) Options() byte {
	return f.buffer[nodeIDOptionsOffset]
}

func (f *NodeID) RemoteAddr64() uint64 {
	return binary.BigEndian.Uint64(f.buffer[nodeIDRemoteAddr64Offset : nodeIDRemoteAddr64Offset+addr64Length])
}

func (f *NodeID) NodeID() []byte {
	return f.buffer[nodeIDStringOffset : f.nodeIdDeviceTypeOffset-3]
}

func (f *NodeID) NodeIDAsString() string {
	return string(f.NodeID())
}

func (f *NodeID) DeviceType() string {
	var devType string
	switch f.buffer[f.nodeIdDeviceTypeOffset] {
	case 0x00:
		devType = "coordinator"
	case 0x01:
		devType = "normal mode"
	case 0x02:
		devType = "end device"
	default:
		devType = "unknown"
	}
	return devType
}
