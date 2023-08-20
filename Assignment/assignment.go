package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/HewlettPackard/structex"
)

// Nas5GSUpdateType represents a data structure for testing purposes.
type Nas5GSUpdateType  struct {
    IEI                     uint8                  // Byte 0         // Information Element Identifier
    Length                  uint8                  // Byte 1         // Length of the encoded data
    SMS_requested           uint8 `bitfield:"1"`   // Byte 2 Start   // Indicator for SMS request
    NG_RAN_RCU              uint8 `bitfield:"1"`                     // Indicator for NG-RAN RCU
    GS5_PNB_CIoT            uint8 `bitfield:"2"`                     // Indicator for 5GS PNB CIoT
    EPS_PNB_CIoT            uint8 `bitfield:"2"`                     // Indicator for EPS PNB CIoT
    Spare_1                 uint8 `bitfield:"1"`                     // Spare field 1
    Spare_2                 uint8 `bitfield:"1"`   // Byte 2 End     // Spare field 2
}
// A function that converts any Nas5GSUpdateType object into a byte stream.
func (ie Nas5GSUpdateType) Encode(buffer *bytes.Buffer){
   newBuffer:= structex.NewBuffer(ie)
   err:= structex.Encode(newBuffer, ie)
   if err!= nil{
    fmt.Println(err)
   }
   buffer.Write(newBuffer.Bytes())
   // Call the Result function to print the result
   Result(buffer)
}
 // Result prints the encoded byte stream.
func Result(buffer *bytes.Buffer){
    var builder strings.Builder
    fmt.Printf("Bytestrom=")
    for index, value := range buffer.Bytes() {
        builder.WriteString(fmt.Sprintf("0x%02X", value))
        if index < len(buffer.Bytes())-1 {
            builder.WriteString(", ")
        }
    }
    fmt.Println(builder.String())
}
func main() {
	// Create a sample information element
    var outputBytes bytes.Buffer
        informationElement := Nas5GSUpdateType {
            IEI: 1,
            Length: 2,
            SMS_requested: 1,
            NG_RAN_RCU: 1,
            GS5_PNB_CIoT: 0,
            EPS_PNB_CIoT: 0,    
            Spare_1:0,
            Spare_2: 0,        
        }  
        // Encode the information element using the Encode function
        informationElement.Encode(&outputBytes)  
}

