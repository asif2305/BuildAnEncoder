package main

import (
	"bytes"
	"testing"
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

    // Compare the lengths of expected and actual encoded bytes
	if len(expectedBytes) != len(actualBytes) {
		t.Errorf("Expected byte length %d, but got %d", len(expectedBytes), len(actualBytes))
		return
	}
    // Compare each byte in the expected and actual encoded bytes
	for i, expected := range expectedBytes {
		if expected != actualBytes[i] {
			t.Errorf("Byte at index %d mismatch, expected: 0x%02X, actual: 0x%02X", i, expected, actualBytes[i])
		}
	}
}