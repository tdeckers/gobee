package tx

import (
	"testing"

	"github.com/pauleyj/gobee/api"
)

type dummyFrame struct {
	data []byte
}

func (f *dummyFrame) Bytes() ([]byte, error) {
	return f.data, nil
}

type txFrameTest struct {
	name     string
	input    Frame
	f        *APIFrame
	expected []byte
}

var txFrameTests = []txFrameTest{
	{"API Frame Default",
		&dummyFrame{data: []byte{0x08, 0x01, 'N', 'J'}},
		New(),
		[]byte{0x7E, 0x00, 0x04, 0x08, 0x01, 'N', 'J', 0x5e}},
	{"API Frame nil Options",
		&dummyFrame{data: []byte{0x08, 0x01, 'N', 'J'}},
		New(nil),
		[]byte{0x7E, 0x00, 0x04, 0x08, 0x01, 'N', 'J', 0x5e}},
	{"API Frame No Escape",
		&dummyFrame{data: []byte{0x08, 0x01, 'N', 'J'}},
		New(api.APIEscapeMode(api.EscapeModeInactive)),
		[]byte{0x7E, 0x00, 0x04, 0x08, 0x01, 'N', 'J', 0x5e}},
	{"API Frame With Escape",
		&dummyFrame{[]byte{0x23, 0x7E, 0x7D, 0x11, 0x13}},
		New(api.APIEscapeMode(api.EscapeModeActive)),
		[]byte{0x7E, 0x00, 0x05, 0x23, 0x7D, 0x5E, 0x7D, 0x5D, 0x7D, 0x31, 0x7D, 0x33, 0xBD}},
}

func TestTXAPIFrame(t *testing.T) {
	t.Parallel()

	t.Run("TX API Suite", func(t *testing.T) {
		for _, tt := range txFrameTests {
			tt := tt

			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				actual, err := tt.f.Bytes(tt.input)
				if err != nil {
					t.Fatalf("Expected no error, but got: %v", err)
				}
				if len(actual) != len(tt.expected) {
					t.Fatalf("Expected API frame to be %d bytes in length, got: %d", len(tt.expected), len(actual))
				}
				for i, b := range actual {
					if b != tt.expected[i] {
						t.Fatalf("Expected 0x%02x, but got 0x%02x at index %d", b, actual[i], i)
					}
				}
			})
		}
	})
}
