package rx

import "fmt"

const (
	atndAddressOffset = 2
)

type atnd struct {
	Data []byte
}

func NewATND(data []byte) *atnd {
	return &atnd{
		Data: data,
	}
}

// AddressFromND AT command
func (a *atnd) AddressAsString() string {
	hex := a.Data[atndAddressOffset : atndAddressOffset+addr64Length]
	return fmt.Sprintf("%#0.16x", hex)
}
