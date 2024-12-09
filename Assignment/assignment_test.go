package BuildAnEncoder

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNas5GSUpdateType_Encode(t *testing.T) {
	// Create an instance of Nas5GSUpdateType for testing
	informationElement := Nas5GSUpdateType{
		IEI:           1,
		Length:        2,
		SMS_requested: 1,
		NG_RAN_RCU:    1,
		GS5_PNB_CIoT:  0,
		EPS_PNB_CIoT:  0,
		Spare_1:       0,
		Spare_2:       0,
	}
	// Create a buffer to capture encoded data
	var buffer bytes.Buffer
	informationElement.Encode(&buffer)

	// Define the expected encoded bytes
	expectedBytes := []byte{0x01, 0x02, 0x03}

	// Get the actual encoded bytes from the buffer
	actualBytes := buffer.Bytes()
	assert.Equal(t, expectedBytes, actualBytes, "Expected Result {0x01, 0x02, 0x03}")
}
